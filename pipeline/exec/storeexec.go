package exec

import (
	"context"
	"fmt"

	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	"github.com/streamingfast/substreams/pipeline/execout"
	"github.com/streamingfast/substreams/reqctx"
	"github.com/streamingfast/substreams/store"
	"google.golang.org/protobuf/proto"
)

type StoreModuleExecutor struct {
	BaseExecutor
	outputStore store.DeltaAccessor
}

func NewStoreModuleExecutor(baseExecutor *BaseExecutor, outputStore store.DeltaAccessor) *StoreModuleExecutor {
	return &StoreModuleExecutor{BaseExecutor: *baseExecutor, outputStore: outputStore}
}

func (e *StoreModuleExecutor) Name() string { return e.moduleName }

func (e *StoreModuleExecutor) String() string { return e.Name() }

func (e *StoreModuleExecutor) Reset() { e.wasmModule.CurrentInstance = nil }

func (e *StoreModuleExecutor) applyCachedOutput(value []byte) error {
	deltas := &pbsubstreams.StoreDeltas{}
	err := proto.Unmarshal(value, deltas)
	if err != nil {
		return fmt.Errorf("unmarshalling output deltas: %w", err)
	}
	e.outputStore.SetDeltas(deltas.Deltas)
	return nil
}

func (e *StoreModuleExecutor) run(ctx context.Context, reader execout.ExecutionOutputGetter) (out []byte, moduleOutput pbsubstreams.ModuleOutputData, err error) {
	ctx, span := reqctx.WithSpan(ctx, "exec_store")
	defer span.EndWithErr(&err)

	if _, err := e.wasmCall(reader); err != nil {
		return nil, nil, fmt.Errorf("store wasm call: %w", err)
	}

	if e.holdsPartialStore() {
		return nil, nil, nil
	}

	return e.wrapDeltas()
}

func (e *StoreModuleExecutor) holdsPartialStore() bool {
	_, ok := e.outputStore.(*store.PartialKV)
	return ok
}

func (e *StoreModuleExecutor) wrapDeltas() (out []byte, moduleOutput pbsubstreams.ModuleOutputData, err error) {
	deltas := &pbsubstreams.StoreDeltas{
		Deltas: e.outputStore.GetDeltas(),
	}

	data, err := proto.Marshal(deltas)
	if err != nil {
		return nil, nil, fmt.Errorf("caching: marshalling delta: %w", err)
	}

	moduleOutput = &pbsubstreams.ModuleOutput_StoreDeltas{
		StoreDeltas: deltas,
	}
	return data, moduleOutput, nil
}
