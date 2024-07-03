package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Rpc struct {
	Client *ethclient.Client
}

func (r *Rpc) GetNonce(address common.Address) (uint64, error) {
	nonce, err := r.Client.PendingNonceAt(context.Background(), address)

	if err != nil {
		return 0, err
	}

	return nonce, nil
}

func (r *Rpc) GetGasPrice() (*big.Int, error) {
	gasPrice, err := r.Client.SuggestGasPrice(context.Background())

	if err != nil {
		return nil, err
	}

	return gasPrice, nil
}

func (r *Rpc) GetChainId() (*big.Int, error) {
	chainID, err := r.Client.NetworkID(context.Background())

	if err != nil {
		return nil, err
	}

	return chainID, nil
}

func (r *Rpc) SendTx(signedTx *types.Transaction) (common.Hash, error) {

	err := r.Client.SendTransaction(context.Background(), signedTx)

	if err != nil {
		return common.Hash{}, err
	}

	return signedTx.Hash(), nil

}

func NewClient(url string) (*Rpc, error) {

	c := Rpc{}

	cl, err := ethclient.Dial(url)

	if err != nil {
		return nil, err
	}

	c.Client = cl

	return &c, nil
}
