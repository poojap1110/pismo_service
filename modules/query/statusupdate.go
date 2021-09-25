package query

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"errors"
)

func UpdateStatus(status, response, description string, ledger *model.TransactionLedger, fndtrnsfer *model.FundTransfer, transactions *model.Transaction, statuscode string) (err2 error) {

	ledger.Status.SetValid(status)
	ledger.Description.SetValid(description)
	ledger.ResponseMessage.SetValid(response)
	ledger.StatusCode.SetValid(statuscode)
	_, err2 = ledger.Update()
	if err2 != nil {
		err2 = errors.New(constant.InternalServerError)
		return
	}

	if fndtrnsfer != nil {
		fndtrnsfer.IntradayMatched.SetValid(fndtrnsfer.IntradayMatched.Int64)
		fndtrnsfer.EodMatched.SetValid(fndtrnsfer.EodMatched.Int64)
		fndtrnsfer.Status.SetValid(status)
		fndtrnsfer.Description.SetValid(description)
		_, err2 = fndtrnsfer.Update()
		if err2 != nil {
			err2 = errors.New(constant.InternalServerError)
			return
		}
	}
	transactions.IntradayMatched.SetValid(transactions.IntradayMatched.Int64)
	transactions.EodMatched.SetValid(transactions.EodMatched.Int64)
	transactions.Status.SetValid(status)
	transactions.Description.SetValid(description)
	transactions.StatusCode.SetValid(statuscode)
	_, err2 = transactions.Update()
	if err2 != nil {
		err2 = errors.New(constant.InternalServerError)
		return
	}
	return
}
