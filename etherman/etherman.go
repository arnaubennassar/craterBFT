package etherman

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/state"
	"github.com/0xPolygonHermez/zkevm-node/test/operations"
	"github.com/arnaubennassar/craterBFT/etherman/cratercontract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v4"
)

type ethereumClient interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.LogFilterer
	ethereum.TransactionReader
	ethereum.TransactionSender

	bind.DeployBackend
}

// Etherman is a package to abstract the Ethereum interactions
type Etherman struct {
	ethClient ethereumClient
	contract  *cratercontract.Cratercontract
	auth      *bind.TransactOpts
}

func New(clientURL string, auth *bind.TransactOpts, craterAddr common.Address) (*Etherman, error) {
	ethClient, err := ethclient.Dial(clientURL)
	if err != nil {
		return nil, err
	}
	contract, err := cratercontract.NewCratercontract(craterAddr, ethClient)
	if err != nil {
		return nil, err
	}

	return &Etherman{
		ethClient: ethClient,
		contract:  contract,
		auth:      auth,
	}, nil
}

func (e *Etherman) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return e.ethClient.CallContract(ctx, call, blockNumber)
}

// CheckTxWasMined check if a tx was already mined
func (e *Etherman) CheckTxWasMined(ctx context.Context, txHash common.Hash) (bool, *types.Receipt, error) {
	receipt, err := e.ethClient.TransactionReceipt(ctx, txHash)
	if errors.Is(err, ethereum.NotFound) {
		return false, nil, nil
	} else if err != nil {
		return false, nil, err
	}

	return true, receipt, nil
}

// CurrentNonce returns the current nonce for the provided account
func (e *Etherman) CurrentNonce(ctx context.Context, account common.Address) (uint64, error) {
	return e.ethClient.NonceAt(ctx, account, nil)
}

// GetTx function get ethereum tx
func (e *Etherman) GetTx(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	return e.ethClient.TransactionByHash(ctx, txHash)
}

// GetTxReceipt function gets ethereum tx receipt
func (e *Etherman) GetTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return e.ethClient.TransactionReceipt(ctx, txHash)
}

// WaitTxToBeMined waits for an L1 tx to be mined. It will return error if the tx is reverted or timeout is exceeded
func (e *Etherman) WaitTxToBeMined(ctx context.Context, tx *types.Transaction, timeout time.Duration) error {
	return operations.WaitTxToBeMined(ctx, e.ethClient, tx, timeout)
}

// SendTx sends a tx to L1
func (e *Etherman) SendTx(ctx context.Context, tx *types.Transaction) error {
	return e.ethClient.SendTransaction(ctx, tx)
}

// SuggestedGasPrice returns the suggest nonce for the network at the moment
func (e *Etherman) SuggestedGasPrice(ctx context.Context) (*big.Int, error) {
	return e.ethClient.SuggestGasPrice(ctx)
}

// EstimateGas returns the estimated gas for the tx
func (e *Etherman) EstimateGas(ctx context.Context, from common.Address, to *common.Address, value *big.Int, data []byte) (uint64, error) {
	return e.ethClient.EstimateGas(ctx, ethereum.CallMsg{
		From:  from,
		To:    to,
		Value: value,
		Data:  data,
	})
}

// SignTx tries to sign a transaction accordingly to the provided sender
func (e *Etherman) SignTx(ctx context.Context, sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
	return e.auth.Signer(e.auth.From, tx)
}

// GetRevertMessage tries to get a revert message of a transaction
func (e *Etherman) GetRevertMessage(ctx context.Context, tx *types.Transaction) (string, error) {
	if tx == nil {
		return "", nil
	}

	receipt, err := e.GetTxReceipt(ctx, tx.Hash())
	if err != nil {
		return "", err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		revertMessage, err := operations.RevertReason(ctx, e.ethClient, tx, receipt.BlockNumber)
		if err != nil {
			return "", err
		}
		return revertMessage, nil
	}
	return "", nil
}

func (e *Etherman) GetLastBlock(ctx context.Context, dbTx pgx.Tx) (*state.Block, error) {
	block, err := e.ethClient.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &state.Block{
		BlockNumber: block.NumberU64(),
		BlockHash:   block.Hash(),
		ParentHash:  block.ParentHash(),
		ReceivedAt:  time.Unix(int64(block.Time()), 0),
	}, nil
}
