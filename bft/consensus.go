package bft

import (
	"context"

	"github.com/cometbft/cometbft/abci/types"
)

func (app *CraterNode) PrepareProposal(ctx context.Context, proposal *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {

	// latestBlock, err := app.eth.BlockNumber(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// latestHash, err := app.eth.BlockByNumber(ctx, big.NewInt(int64(latestBlock)))
	// if err != nil {
	// 	return nil, err
	// }

	// state := engine.ForkChoiceState{
	// 	HeadHash:           latestHash.Hash(),
	// 	SafeBlockHash:      latestHash.Hash(),
	// 	FinalizedBlockHash: common.Hash{},
	// }

	// // The engine complains when the withdrawals are empty
	// withdrawals := []*engine.Withdrawal{
	// 	{
	// 		Index:     "0x0",
	// 		Validator: "0x0",
	// 		Address:   common.Address{}.Hex(),
	// 		Amount:    "0x0",
	// 	},
	// }

	// attrs := engine.PayloadAttributes{
	// 	Timestamp:             hexutil.Uint64(time.Now().UnixMilli()),
	// 	PrevRandao:            common.Hash{}, // do we need to generate a randao for the EVM?
	// 	SuggestedFeeRecipient: app.addr,
	// 	Withdrawals:           withdrawals,
	// }

	// choice, err := app.engine.ForkchoiceUpdatedV2(&state, &attrs)
	// if err != nil {
	// 	return nil, err
	// }

	// payloadId := choice.PayloadId
	// status := choice.PayloadStatus

	// if status.Status != "VALID" {
	// 	app.logger.Error("validation err: %v, critical err: %v", status.ValidationError, status.CriticalError)
	// 	return nil, errors.New(status.ValidationError)
	// }

	// payload, err := app.engine.GetPayloadV2(payloadId)
	// if err != nil {
	// 	return nil, err
	// }

	// // this is where we could filter/reorder transactions, or mark them for filtering so consensus could be checked

	// bytes, err := json.Marshal(payload.ExecutionPayload)
	// if err != nil {
	// 	return nil, err
	// }
	// txs := make([][]byte, 1)
	// txs[0] = bytes

	// return &types.ResponsePrepareProposal{
	// 	Txs: txs,
	// }, nil
	return nil, nil
}

func (app *CraterNode) ProcessProposal(_ context.Context, proposal *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	// TODO: accept or reject the proposal
	return &types.ResponseProcessProposal{
		Status: types.ResponseProcessProposal_ACCEPT,
	}, nil
}

func (app *CraterNode) FinalizeBlock(_ context.Context, block *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	// // the first and only "tx" is the serialized ExecutionPayload (with the actual txs)
	// var tx0 engine.ExecutionPayload
	// err := json.Unmarshal(block.Txs[0], &tx0)
	// if err != nil {
	// 	log.Warn("transaction decode error: %v")
	// 	result := &types.ExecTxResult{
	// 		Code:      EXECUTION_ERROR,
	// 		Data:      nil,
	// 		Log:       err.Error(),
	// 		Info:      "failed to decode transaction, cannot proceed",
	// 		GasWanted: 0,
	// 		GasUsed:   0,
	// 		Events:    nil,
	// 		Codespace: EXECUTION_CODESPACE,
	// 	}
	// 	return &types.ResponseFinalizeBlock{TxResults: []*types.ExecTxResult{result}, AppHash: app.state.Hash()}, nil
	// }

	// block.GetTime()

	// payload, err := app.retryUntilNewPayload(tx0)
	// if err != nil {
	// 	// this should not happen
	// 	log.Warn("new payload error: %v")
	// 	result := &types.ExecTxResult{
	// 		Code:      EXECUTION_ERROR,
	// 		Data:      nil,
	// 		Log:       err.Error(),
	// 		Info:      "new payload failed, cannot proceed",
	// 		GasWanted: 0,
	// 		GasUsed:   0,
	// 		Events:    nil,
	// 		Codespace: EXECUTION_CODESPACE,
	// 	}
	// 	return &types.ResponseFinalizeBlock{TxResults: []*types.ExecTxResult{result}, AppHash: app.state.Hash()}, nil
	// }

	// state := engine.ForkChoiceState{
	// 	HeadHash:           common.HexToHash(payload.LatestValidHash),
	// 	SafeBlockHash:      common.HexToHash(payload.LatestValidHash),
	// 	FinalizedBlockHash: common.HexToHash(tx0.ParentHash), // latestHash from the Proposal stage
	// }

	// // The engine complains when the withdrawals are empty
	// withdrawals := []*engine.Withdrawal{
	// 	{
	// 		Index:     "0x0",
	// 		Validator: "0x0",
	// 		Address:   common.Address{}.Hex(),
	// 		Amount:    "0x0",
	// 	},
	// }

	// attrs := engine.PayloadAttributes{
	// 	Timestamp:             hexutil.Uint64(time.Now().UnixMilli()),
	// 	PrevRandao:            common.Hash{},
	// 	SuggestedFeeRecipient: app.addr,
	// 	Withdrawals:           withdrawals,
	// }

	// var choice *engine.ForkchoiceUpdatedResponse
	// choice, err = app.retryUntilForkchoiceUpdated(&state, &attrs)
	// if err != nil {
	// 	infoLog := "fork choice failed, cannot proceed"
	// 	log.Error(infoLog, "error", err.Error())
	// 	if choice != nil && choice.PayloadStatus.Status != "VALID" {
	// 		infoLog = fmt.Sprintf("%s: %s", choice.PayloadStatus.ValidationError)
	// 	}
	// 	result := &types.ExecTxResult{
	// 		Code:      EXECUTION_ERROR,
	// 		Data:      nil,
	// 		Log:       err.Error(),
	// 		Info:      infoLog,
	// 		GasWanted: 0,
	// 		GasUsed:   0,
	// 		Events:    nil,
	// 		Codespace: EXECUTION_CODESPACE,
	// 	}
	// 	return &types.ResponseFinalizeBlock{TxResults: []*types.ExecTxResult{result}, AppHash: app.state.Hash()}, nil
	// }

	// result := &types.ExecTxResult{Code: EXECUTION_SUCCESS}

	// app.state.Height = block.Height // save state later in Commit
	// app.state.Size = block.Height   // only one transaction is executed, the ExecutionPayload, so the size is the height

	// return &types.ResponseFinalizeBlock{TxResults: []*types.ExecTxResult{result}, AppHash: app.state.Hash()}, nil
	return nil, nil
}

func (app *CraterNode) Commit(_ context.Context, _ *types.RequestCommit) (*types.ResponseCommit, error) {
	// if err := app.state.Save(); err != nil {
	// 	app.logger.Error("app failed to save state: %v", err)
	// 	return nil, err // what happens when commit fails?
	// }

	return &types.ResponseCommit{}, nil
}

func (app *CraterNode) ExtendVote(_ context.Context, _ *types.RequestExtendVote) (*types.ResponseExtendVote, error) {
	return &types.ResponseExtendVote{}, nil
}

func (app *CraterNode) VerifyVoteExtension(_ context.Context, _ *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error) {
	return &types.ResponseVerifyVoteExtension{}, nil
}

// func (app *CraterNode) retryUntilNewPayload(payload engine.ExecutionPayload) (response *engine.NewPayloadResponse, err error) {
// 	forever := backoff.NewExponentialBackOff()
// 	err = backoff.Retry(func() error {
// 		response, err = app.engine.NewPayloadV2(payload)
// 		if forever.NextBackOff() > RESET_EXPONENTIAL_BACKOFF {
// 			forever.Reset()
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}, forever)
// 	if err != nil {
// 		return nil, err // should not happen, retries forever
// 	}
// 	return response, nil
// }

// func (app *CraterNode) retryUntilForkchoiceUpdated(state *engine.ForkChoiceState, attrs *engine.PayloadAttributes) (response *engine.ForkchoiceUpdatedResponse, err error) {
// 	forever := backoff.NewExponentialBackOff()
// 	err = backoff.Retry(func() error {
// 		response, err = app.engine.ForkchoiceUpdatedV2(state, attrs)
// 		if forever.NextBackOff() > RESET_EXPONENTIAL_BACKOFF {
// 			forever.Reset()
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}, forever)
// 	if err != nil {
// 		return nil, err // should not happen, retries forever
// 	}
// 	return response, nil
// }
