package rpc

import (
	"github.com/INFURA/mobydick/utils"
	"net/http"
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type GetLogsResponse struct {
	ID       int `json:"id"`
	Jsonrpc      string `json:"jsonrpc"`
	GetLogsResp  []GetLogsRespModel `json:"result"`
}
type GetLogsRespModel struct{
	Address  string `json:"address"`
	BlockHash  string `json:"blockHash"`
	BlockNumber  string `json:"blockNumber"`
	Topics  []string `json:"topics"`
	TransactionHash  string `json:"transactionHash"`
	Data  string `json:"data"`
	Gen  interface{}
}

type GetBlockByNumberRespModel struct{
	Timestamp  string `json:"timestamp"`
	Gen  interface{}
}

type GetBlockByNumberResponse struct {
	ID       int `json:"id"`
	Jsonrpc      string `json:"jsonrpc"`
	GetBlockByNumberResp GetBlockByNumberRespModel `json:"result"`
}

type GetBlockNumberResponse struct {
	ID       int `json:"id"`
	Jsonrpc      string `json:"jsonrpc"`
	Result string `json:"result"`
}

type GetLogsParam struct {
	Address string `json:"address"`
	FromBlock string `json:"fromBlock"`
	ToBlock string `json:"toBlock"`
	Topics []string `json:"topics"`
}

type GetLogsRequest struct{
	ID int `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	GetLogsParams []GetLogsParam `json:"params"`
}

type GetBlockNumberRequest struct{
	ID int `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	GetBlockNumberParams []string `json:"params"`
}

func GetBlockByNumber(blockNumber string) GetBlockByNumberResponse{
	params := []interface{}{blockNumber, true}
		jsonReq := map[string]interface{}{"jsonrpc": "2.0", "method": "eth_getBlockByNumber", "params": params, "id": 0}
		data, _ := json.Marshal(jsonReq)
		requestBody1 := bytes.NewBuffer(data)
		resp1, err := http.Post(utils.MainnetEndpoint, "application/json", requestBody1)
		if err != nil {
			fmt.Println(err)
		}
		defer resp1.Body.Close()
		body1, _ := ioutil.ReadAll(resp1.Body) 
		var getBlockByNumberResponse GetBlockByNumberResponse
		if err := json.Unmarshal(body1, &getBlockByNumberResponse); err != nil {
				fmt.Println("All good marshalling")
		}
		return getBlockByNumberResponse
}

func GetLogs(from string, to string, address string) GetLogsResponse{
	getLogsParams := GetLogsParam{
		Address: address,
		FromBlock: from,
		ToBlock: to,
		// ERC20 Transfer method signature
		Topics: []string{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"},
	}	
	getLogsRequest := GetLogsRequest{
		ID: 1234,
	 	Jsonrpc: "2.0",
		Method: "eth_getLogs",
		GetLogsParams: []GetLogsParam{getLogsParams},
	}
	postBody, _ := json.Marshal(getLogsRequest)
	requestBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(utils.MainnetEndpoint, "application/json", requestBody)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) 

	var getLogsResponse GetLogsResponse
	if err := json.Unmarshal(body, &getLogsResponse); err != nil {   // Parse []byte to go struct pointer
			fmt.Println("All good marshalling")
	}

	return getLogsResponse
}

func GetBlockNumber() string{
	jsonReq := map[string]interface{}{"jsonrpc": "2.0", "method": "eth_blockNumber", "id": 0}
	postBody, _ := json.Marshal(jsonReq)
	requestBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(utils.MainnetEndpoint, "application/json", requestBody)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) 

	var getBlockNumberResponse GetBlockNumberResponse
	if err := json.Unmarshal(body, &getBlockNumberResponse); err != nil {  
			fmt.Println("All good marshalling")
	}

	return getBlockNumberResponse.Result
}