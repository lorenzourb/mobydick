package main

import (
	"strconv"
	"github.com/INFURA/mobydick/utils"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	 "time"	 
	"github.com/INFURA/mobydick/rpc"
	"github.com/INFURA/mobydick/pgdb"
	"database/sql" 
	"math/big"
)

func HandleRequest(ctx context.Context) {
	db := pgdb.Connect()

	lastBlockNumber := rpc.GetBlockNumber()
	lastBlockNumberInt :=  utils.HexToInt(lastBlockNumber)
	lastBlockDB := pgdb.GetLastBlockNumber(db)

	fmt.Printf(fmt.Sprintf("Block difference: %d\n", lastBlockNumberInt - lastBlockDB))

	if(int64(lastBlockNumberInt) - lastBlockDB == 0) {
		fmt.Println("Nothing to update as no block diff")
		return
	} 

	allTokenPrices := utils.FetchAllTokenPrices()
	
	for _,t := range utils.Tokens {
		logs := rpc.GetLogs(fmt.Sprintf("0x%x", lastBlockDB), lastBlockNumber, t.Address)
		price,_ := strconv.ParseFloat(utils.FilterTokenPrice(allTokenPrices, t.ID), 64)
		insertTxs(logs, db, t, price)
	}

	pgdb.InsertLastBlockNumber(lastBlockNumberInt, db)
}

func insertTxs(logs rpc.GetLogsResponse, db *sql.DB, token utils.Token, tokenPrice float64) {	
	for _,r := range logs.GetLogsResp {
		data := utils.HexToBigInt(r.Data)
		if(data.Cmp(big.NewInt(utils.TxAmountThreshold)) > 0){
			getBlockByNumberResponse := rpc.GetBlockByNumber(r.BlockNumber)
			i :=  utils.HexToInt(getBlockByNumberResponse.GetBlockByNumberResp.Timestamp)
			tm := time.Unix(i, 0)
			pgdb.InsertTransfer(r, tm, db, token, tokenPrice)		
		}
	}
}

func main() {
	lambda.Start(HandleRequest)
}
