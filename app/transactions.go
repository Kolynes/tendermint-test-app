package app

import (
	"bytes"
	"errors"

	"github.com/kolynes/tendermint-app/store"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func decodeTranasaction(transaction []byte) (string, any, error) {
	parts := bytes.Split(transaction, []byte("="))
	if len(parts) != 2 {
		return "", "", errors.New("invalid transaction")
	}
	return string(parts[0]), parts[1], nil
}

func isValid(transaction []byte) uint32 {
	_, _, err := decodeTranasaction(transaction)
	if err != nil {
		return 1
	}
	return 0
}

func (app *Application) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	return abcitypes.ResponseCheckTx{Code: isValid(req.Tx)}
}

func (app *Application) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	app.batch = store.NewBatch(app.db)
	return abcitypes.ResponseBeginBlock{}
}

func (app *Application) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	key, value, err := decodeTranasaction(req.Tx)
	if err != nil {
		return abcitypes.ResponseDeliverTx{Code: 1}
	}
	app.batch.Put(key, value)
	return abcitypes.ResponseDeliverTx{Code: 0}
}

func (Application) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (app *Application) Commit() abcitypes.ResponseCommit {
	app.batch.Commit()
	return abcitypes.ResponseCommit{Data: []byte{}}
}
