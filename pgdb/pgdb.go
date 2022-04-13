package pgdb

import (
	"github.com/INFURA/mobydick/utils"
	"fmt"
	"database/sql" 
	_ "github.com/lib/pq"
	"github.com/INFURA/mobydick/rpc"
	"time"
	"math/big"
)

func Connect() *sql.DB {
	dbConnectionString := utils.PostgresEndpoint
	// Connect to database
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
			fmt.Println(err)
	}

	return db
}

func InsertLastBlockNumber(lastBlockNumber int64, db *sql.DB) {
	sqlStatement :=`INSERT INTO "blockNumbers" ("blockNumber") VALUES  ($1)`
	_, err := db.Exec(sqlStatement, lastBlockNumber)
	if err != nil {
		fmt.Println("An error occured while executing query: ", err)
	}
}

func InsertTransfer(r rpc.GetLogsRespModel, tm time.Time, db *sql.DB, token utils.Token, tokenPrice float64){
	sqlStatement :=`INSERT INTO transfers ("timestamp", token, amount, "from", "to", "blockNumber", "txHash") VALUES  ($1,$2,$3,$4,$5,$6,$7)`
	data := utils.HexToBigInt(r.Data)
	var dataFormatted *big.Int
	if(token.Decimals == 6) {
		dataFormatted = new(big.Int).Quo(data, big.NewInt(utils.SixDigitsThreshold))
	} else {
		dataFormatted = new(big.Int).Quo(data, big.NewInt(utils.EighteenDigitsThreshold))
	}
	
	// This calculation is actually not precise, as get the token price at a slightly different time of when the actual tx happened
	_, err := db.Exec(sqlStatement, tm.Format("2006-01-02 15:04:05") , token.Name, int(float64(dataFormatted.Int64()) * tokenPrice), r.Topics[1], r.Topics[2], r.BlockNumber, r.TransactionHash)
	if err != nil {
		fmt.Println("An error occured while executing query: ", err)
	} else {
		fmt.Printf("Inserted transfer %s %s %d %s %s %s %s", tm.Format("2006-01-02 15:04:05") , r.Address,  int(float64(dataFormatted.Int64()) * tokenPrice), r.Topics[1], r.Topics[2], r.BlockNumber, r.TransactionHash)
	}
}

func GetLastBlockNumber(db *sql.DB)int64{
	rows, err := db.Query(`SELECT "blockNumber"	FROM "blockNumbers" ORDER BY "blockNumber" DESC LIMIT 1`)
	if err != nil {
							panic(err)
	}
	defer rows.Close()
	var blockNumber int 
	for rows.Next() {
		err := rows.Scan(&blockNumber)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return int64(blockNumber)
}