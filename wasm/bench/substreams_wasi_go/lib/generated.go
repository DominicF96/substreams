package lib

import (
	"fmt"

	"github.com/streamingfast/substreams/wasm/bench/substreams_wasi_go/pb"
	"github.com/streamingfast/substreams/wasm/wasi/substreams"
)

func init() {
	substreams.Register("mapBlock", ExecuteMapBlock)
}

type MapBlockInput struct {
	block      *pb.Block
	readStore1 substreams.StoreGet[string]
	readStore2 substreams.StoreGet[string]
}

func ExecuteMapBlock(input []byte) error {
	res := &pb.MapBlockInput{}
	err := res.UnmarshalVT(input)
	if err != nil {
		return fmt.Errorf("unmarshalling args: %w", err)
	}
	mapBlockInputs := &MapBlockInput{
		block:      res.Block,
		readStore1: substreams.NewGetStringStore(res.ReadStore),
		readStore2: substreams.NewGetStringStore(res.ReadStore2),
	}

	out, err := mapBlock(mapBlockInputs)
	if err != nil {
		return fmt.Errorf("mapping block: %w", err)
	}
	data, err := out.MarshalVT()
	if err != nil {
		return fmt.Errorf("marshalling output: %w", err)
	}

	_, err = substreams.WriteOutput(data)
	if err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}
