package pgdb

import (
	"github.com/INFURA/mobydick/utils"
	"fmt"
	"database/sql" 
	_ "github.com/lib/pq"
	"github.com/INFURA/mobydick/rpc"
	"time"
)

func Connect() *sql.DB {
	dbConnectionString := utils.PostgresEndpoint
	// Connect to database
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func InsertLastBlockNumber(lastBlockNumber int64, db *sql.DB) {
	sqlStatement :=`INSERT INTO "block_number" ("block_number") VALUES  ($1)`
	_, err := db.Exec(sqlStatement, lastBlockNumber)
	if err != nil {
		fmt.Println("An error occured while executing query: ", err)
		panic(err)
	}
}

func StoreWhale(address string, db *sql.DB) {
	selectStatement :=`SELECT COUNT(*) FROM whales WHERE address = '`+ address +`';`
	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Println("An error occured while executing query: ", err)
		panic(err) 
	}
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	if (count > 0) { 
		return 
	} 

	var typ string 
	if(rpc.IsContractAddress(address)){
		typ = "contract"
	} else {
		typ = "address"
	}
	insertStatement := `INSERT INTO "whales" ("address", "name", "type") VALUES  ($1, $2, $3)`
	_, err = db.Exec(insertStatement, address, "", typ)
	if err != nil {
		fmt.Println("An error occured while executing query: ", err)
		panic(err)
	}
}

func InsertTransfer(r rpc.GetLogsRespModel, tm time.Time, db *sql.DB, token utils.Token, amount int64){
	sqlStatement :=`INSERT INTO transfer ("timestamp", token, amount, "from", "to", "block_number", "tx_hash") VALUES  ($1,$2,$3,$4,$5,$6,$7)`
	// This calculation is actually not precise, as get the token price at a slightly different time of when the actual tx happened
	_, err := db.Exec(sqlStatement, tm.Format("2006-01-02 15:04:05") , token.Name, amount, utils.Unpad(r.Topics[1]), utils.Unpad(r.Topics[2]), r.BlockNumber, r.TransactionHash)
	if err != nil {
		fmt.Println("An error occured while executing query: ", err)
		panic(err)
	} else {
		fmt.Printf("Inserted transfer %s %s %d %s %s %s %s\n", tm.Format("2006-01-02 15:04:05") , r.Address, amount, utils.Unpad(r.Topics[1]), utils.Unpad(r.Topics[2]), r.BlockNumber, r.TransactionHash)
	}
}

func GetLastBlockNumber(db *sql.DB)int64{
	rows, err := db.Query(`SELECT "block_number"	FROM "block_number" ORDER BY "block_number" DESC LIMIT 1`)
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