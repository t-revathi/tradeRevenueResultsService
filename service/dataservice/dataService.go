package service

import (
	"context"
	"fmt"
	"time"
	"net/http"
	

)

type DataService struct {
}

func NewDataService() *DataService {
	return &DataService{}
}

func (d *DataService) ShowAllRevenue(ctx context.Context, w http.ResponseWriter, r *http.Request, req *DataCalculateRevenue) RevenueCollection{

	fmt.Printf("Received financial year:%s on %s", req.Config.FinancialYear, time.Now().String())
	transactionData := req.TransactionData

	/*render.JSON(w, r,
	DataCalculateRevenue{
		transactionData,
		req.Config,
	})*/
	income := ProcessTransactions(transactionData, req.Config)

	/*render.JSON(w, r,
		income.Items)*/
		return income

}


