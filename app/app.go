package app

import (
	"github.com/kolynes/tendermint-app/store"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type Application struct {
	db    *store.Store
	batch *store.Batch
}

var _ abcitypes.Application = (*Application)(nil)

func NewApplication(db *store.Store) *Application {
	return &Application{
		db: db,
	}
}

func (Application) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (Application) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (Application) SetOption(abcitypes.RequestSetOption) abcitypes.ResponseSetOption {
	return abcitypes.ResponseSetOption{Code: 0}
}
