package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/streamingfast/logging"
	"github.com/streamingfast/substreams/metrics"
	"github.com/streamingfast/substreams/wasm"
	_ "github.com/streamingfast/substreams/wasm/wasi"
	_ "github.com/streamingfast/substreams/wasm/wasmtime"
	_ "github.com/streamingfast/substreams/wasm/wazero"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func init() {
	logging.InstantiateLoggers()
}

func BenchmarkExecution(b *testing.B) {
	type runtime struct {
		name                string
		wasmCodeType        string
		code                []byte
		shouldReUseInstance bool
	}

	type testCase struct {
		tag        string
		entrypoint string
		arguments  []wasm.Argument
		// Right now there is differences between runtime, so we accept all those values
		acceptedByteCount []int
	}

	for _, testCase := range []*testCase{
		//{"bare", "map_noop", args(wasm.NewParamsInput("")), []int{0}},

		// Decode proto only decode and returns the block.number as the output (to ensure the block is not elided at compile time)
		//{"decode_proto_only", "map_decode_proto_only", args(blockInputFile(b, "testdata/ethereum_mainnet_block_16021772.binpb")), []int{0}},

		{"map_block", "map_block", args(blockInputFile(b, "testdata/ethereum_mainnet_block_16021772.binpb")), []int{53}},
	} {
		var reuseInstance = true
		//var freshInstanceEachRun = false

		//wasmCode := readCode(b, "substreams_wasm/substreams.wasm")
		//wasmCodep1 := readCode(b, "substreams_ts/index.wasm")
		wasmTinyGo := readCode(b, "substreams_tiny_go/main.wasm")
		//wasmCodep1 := readCode(b, "substreams_wasi/wasi_hello_world/hello.wasm")

		stats := metrics.NewReqStats(&metrics.Config{}, zap.NewNop())
		for _, config := range []*runtime{
			//{"wasmtime", wasmCode, reuseInstance},
			//{"wasmtime", wasmCode, freshInstanceEachRun},
			//
			//{"wazero", wasmCode, reuseInstance},
			//{"wazero", wasmCode, freshInstanceEachRun},

			{name: "wasi", wasmCodeType: "go/wasi", code: wasmTinyGo, shouldReUseInstance: reuseInstance},
			//{"wasip1", wasmCodep1, freshInstanceEachRun},
		} {
			instanceKey := "reused"
			if !config.shouldReUseInstance {
				instanceKey = "fresh"
			}

			b.Run(fmt.Sprintf("vm=%s,instance=%s,tag=%s", config.name, instanceKey, testCase.tag), func(b *testing.B) {
				ctx := context.Background()

				wasmRuntime := wasm.NewRegistryWithRuntime(config.name, nil)

				module, err := wasmRuntime.NewModule(ctx, config.code, config.wasmCodeType)
				require.NoError(b, err)

				cachedInstance, err := module.NewInstance(ctx)
				require.NoError(b, err)
				defer cachedInstance.Close(ctx)

				call := wasm.NewCall(nil, testCase.tag, testCase.entrypoint, stats, testCase.arguments)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					start := time.Now()
					instance := cachedInstance
					if !config.shouldReUseInstance {
						instance, err = module.NewInstance(ctx)
						require.NoError(b, err)
					}

					_, err := module.ExecuteNewCall(ctx, call, instance, testCase.arguments)
					if err != nil {
						require.NoError(b, err)
					}
					fmt.Println("call output", string(call.Output()))
					fmt.Println("duration", time.Since(start))
					//require.Contains(b, testCase.acceptedByteCount, len(call.Output()), "invalid byte count got %d expected one of %v", len(call.Output()), testCase.acceptedByteCount)
				}
			})
		}
	}
}

func readCode(t require.TestingT, filename string) []byte {
	content, err := os.ReadFile(filename)
	require.NoError(t, err)

	return content
}

func args(ins ...wasm.Argument) []wasm.Argument {
	return ins
}

func blockInputFile(t require.TestingT, filename string) wasm.Argument {
	content, err := os.ReadFile(filename)
	require.NoError(t, err)

	input := wasm.NewSourceInput("sf.ethereum.type.v2.Block", 0)
	input.SetValue(content)

	return input
}
