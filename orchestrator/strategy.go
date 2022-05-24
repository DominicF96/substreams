package orchestrator

import (
	"context"
	"fmt"

	"github.com/streamingfast/substreams/block"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	"github.com/streamingfast/substreams/state"
	"go.uber.org/zap"
)

type Strategy interface {
	GetNextRequest() (*pbsubstreams.Request, error)
}

type LinearStrategy struct {
	requests []*pbsubstreams.Request
}

func NewLinearStrategy(ctx context.Context, request *pbsubstreams.Request, builders []*state.Builder, upToBlockNum uint64) (*LinearStrategy, error) {
	res := &LinearStrategy{}

	for _, builder := range builders {
		zlog.Debug("squashables", zap.String("builder", builder.Name))
		zlog.Debug("up to block num", zap.Uint64("up_to_block_num", upToBlockNum))
		if upToBlockNum == builder.ModuleStartBlock {
			continue // nothing to synchronize
		}

		endBlock := upToBlockNum
		info, err := builder.Info(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting builder info: %w", err)
		}

		lastExclusiveEndBlock := info.LastKVSavedBlock
		zlog.Debug("got info", zap.Object("builder", builder), zap.Uint64("up_to_block", upToBlockNum), zap.Uint64("end_block", lastExclusiveEndBlock))
		if upToBlockNum <= lastExclusiveEndBlock {
			zlog.Debug("no request created", zap.Uint64("up_to_block", upToBlockNum), zap.Uint64("last_exclusive_end_block", lastExclusiveEndBlock))
			continue // not sure if we should pop here
		}

		reqStartBlock := lastExclusiveEndBlock
		if reqStartBlock == 0 {
			reqStartBlock = builder.ModuleStartBlock
		}

		requestBlockRange := &block.Range{
			StartBlock:        reqStartBlock,
			ExclusiveEndBlock: endBlock,
		}

		ranges := requestBlockRange.Split(200)
		for _, r := range ranges {
			req := createRequest(r.StartBlock, r.ExclusiveEndBlock, builder.Name, request.ForkSteps, request.IrreversibilityCondition, request.Modules)
			res.requests = append(res.requests, req)
			zlog.Info("request created", zap.String("module_name", builder.Name), zap.Object("block_range", r))
		}
	}

	return res, nil
}

func (s *LinearStrategy) GetNextRequest() (*pbsubstreams.Request, error) {
	if len(s.requests) == 0 {
		return nil, fmt.Errorf("no requests to fetch")
	}

	var request *pbsubstreams.Request
	request, s.requests = s.requests[len(s.requests)-1], s.requests[:len(s.requests)-1]

	return request, nil
}

func createRequest(
	startBlock, stopBlock uint64,
	outputModuleName string,
	forkSteps []pbsubstreams.ForkStep,
	irreversibilityCondition string,
	modules *pbsubstreams.Modules,
) *pbsubstreams.Request {
	return &pbsubstreams.Request{
		StartBlockNum:            int64(startBlock),
		StopBlockNum:             stopBlock,
		ForkSteps:                forkSteps,
		IrreversibilityCondition: irreversibilityCondition,
		Modules:                  modules,
		OutputModules:            []string{outputModuleName},
	}
}