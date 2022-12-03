package app

import abcitypes "github.com/tendermint/tendermint/abci/types"

func (app *Application) Query(req abcitypes.RequestQuery) abcitypes.ResponseQuery {
	value, err := app.db.Get(string(req.Data))
	if err != nil {
		return abcitypes.ResponseQuery{
			Code:  1,
			Log:   "Does not exist",
			Value: value.([]byte),
		}
	}
	return abcitypes.ResponseQuery{
		Code:  0,
		Log:   "Success",
		Value: value.([]byte),
	}
}
