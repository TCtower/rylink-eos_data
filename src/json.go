package main

import (
	"fmt"
	"encoding/json"
)


type Block struct{
	Timestamp string `json:"timestamp"`
	//time
	Producer string `json:"producer"`
	Confirmed int `json:"confirmed"`
	Previous string `json:"previous"`
	Transaction_mroot string `json:"transaction_mroot"`
	Action_mroot string `json:"action_mroot"`
	Schedule_version int `json:"schedule_version"`
	New_producers string `json:"new_producers"`
	Header_extensions []string `json:"header_extensions"`
	Producer_signature string `json:"producer_signature"`
	Transactions []Transaction `json:"transactions"`
	Block_extensions []string `json:"block_extensions"`
	Id string `json:"id"`
	Block_num int `json:"block_num"`
	Ref_block_prefix int `json:"ref_block_prefix"`
}

type Transaction struct{
	Status string `json:"status"`
	Cpu_usage_us int `json:"cpu_usage_us"`
	Net_usage_words int `json:"net_usage_words"`
	Trx Trx_detail `json:"trx"`
}

type Trx_detail struct{
	Id string `json:"id"`
	Signatures []string `json:"signatures"`
	Compression string `json:"compression"`
	//could be enum
	Packed_context_free_data string `json:"packed_context_free_data"`
	Context_free_data []string `json:"context_free_data"`
	Packed_trx string `json:"packed_trx"`
	Transaction Transaction_detail `json:"transaction"`
}

type Transaction_detail struct{
	Expiration string `json:"expiration"`
	//time
	Ref_block_num int `json:"ref_block_num"`
	Ref_block_prefix int `json:"ref_block_prefix"`
	Max_net_usage_words int `json:"max_net_usage_words"`
	Max_cpu_usage_ms int `json:"max_cpu_usage_ms"`
	Delay_sec int `json:"delay_sec"`
	Context_free_actions []string `json:"context_free_actions"`
	Actions []Action `json:"actions"`
	Transaction_extensions []string `json:"transaction_extensions"`
}


type Action struct{
	Account string `json:"account"`
	Name string `json:"name"`
	//could be enum
	Authorization []Authorization_detail `json:"authorization"`
	Data Data_interface `json:"data"`
	Hex_data string `json:"hex_data"`
}

type Data_interface interface{}


type Authorization_detail struct{
	Actor string `json:"actor"`
	Permission string `json:"permission"`
}

func ite_map(m map[string]interface{}){
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func toStruct(str string){
	block_msg :=Block{}
	err := json.Unmarshal([]byte(str), &block_msg)
	if err != nil {
		fmt.Println("error is %v\n", err)
	} else {
		fmt.Printf("%v\n", block_msg)
	}
}

func test() {
	str := `{
  "timestamp": "2018-07-02T05:58:52.000",
  "producer": "eosbeijingbp",
  "confirmed": 0,
  "previous": "0038dd01209820968a3b389b55136b6251138c7e347ee0fd0a4bfc2f23f7f1fc",
  "transaction_mroot": "2425f795728976d353b2b7f451749fd5747245859a02eb12bf79689070957b3a",
  "action_mroot": "e9d949ee2898cf6ec27ccb97f2ca0e0a9dac4a5dd2070b90f776a23e794ea169",
  "schedule_version": 93,
  "new_producers": null,
  "header_extensions": [],
  "producer_signature": "SIG_K1_KamF3p7ds9cTD8tHBmpnGMxmjGbu4xk5osndfSXNxFsrVvtY912jycjw2S3B5zZXQHaHL7SHynVhRdsRQTkksRoUJwNCYm",
  "transactions": [
    {
      "status": "executed",
      "cpu_usage_us": 11276,
      "net_usage_words": 42,
      "trx": {
        "id": "0ef3c9c883e1bdc73c51d08276cef90e3063f4315951c0a3970515f9fc3ef336",
        "signatures": [
          "SIG_K1_JzkeBXiv59TRHwB1QmHm8HLQ1Gk6o1TDCYAaLDAC2ini3n8XUXJJksWYmtLrLx3bG4vVK3ufNaUKBRJU6ozkK7Aijgsh53"
        ],
        "compression": "none",
        "packed_context_free_data": "",
        "context_free_data": [],
        "packed_trx": "57bf395bb5db2d03d0a600000000030000000000ea305500409e9a2264b89a0190a7a60819a5ee0a00000000a8ed32326690a7a60819a5ee0a3044a0a6fde90ebd01000000010002b7a56b23289e29dbf566f0d28fc0390d9d093aa93c05a77085c0e39919e910330100000001000000010002b7a56b23289e29dbf566f0d28fc0390d9d093aa93c05a77085c0e39919e91033010000000000000000ea305500b0cafe4873bd3e0190a7a60819a5ee0a00000000a8ed32321490a7a60819a5ee0a3044a0a6fde90ebd001000000000000000ea305500003f2a1ba6a24a0190a7a60819a5ee0a00000000a8ed32323190a7a60819a5ee0a3044a0a6fde90ebd640000000000000004454f5300000000640000000000000004454f53000000000100",
        "transaction": {
          "expiration": "2018-07-02T05:59:51",
          "ref_block_num": 56245,
          "ref_block_prefix": 2798650157,
          "max_net_usage_words": 0,
          "max_cpu_usage_ms": 0,
          "delay_sec": 0,
          "context_free_actions": [],
          "actions": [
            {
              "account": "eosio",
              "name": "newaccount",
              "authorization": [
                {
                  "actor": "1freeaccount",
                  "permission": "active"
                }
              ],
              "data": {
                "creator": "1freeaccount",
                "name": "robinzhao123",
                "owner": {
                  "threshold": 1,
                  "keys": [
                    {
                      "key": "EOS6HNMAWkwwLNXmqMkbxxYw3FFZQjwcqniwiBNxr6Jn7EEXcjL9Z",
                      "weight": 1
                    }
                  ],
                  "accounts": [],
                  "waits": []
                },
                "active": {
                  "threshold": 1,
                  "keys": [
                    {
                      "key": "EOS6HNMAWkwwLNXmqMkbxxYw3FFZQjwcqniwiBNxr6Jn7EEXcjL9Z",
                      "weight": 1
                    }
                  ],
                  "accounts": [],
                  "waits": []
                }
              },
              "hex_data": "90a7a60819a5ee0a3044a0a6fde90ebd01000000010002b7a56b23289e29dbf566f0d28fc0390d9d093aa93c05a77085c0e39919e910330100000001000000010002b7a56b23289e29dbf566f0d28fc0390d9d093aa93c05a77085c0e39919e9103301000000"
            },
            {
              "account": "eosio",
              "name": "buyrambytes",
              "authorization": [
                {
                  "actor": "1freeaccount",
                  "permission": "active"
                }
              ],
              "data": {
                "payer": "1freeaccount",
                "receiver": "robinzhao123",
                "bytes": 4096
              },
              "hex_data": "90a7a60819a5ee0a3044a0a6fde90ebd00100000"
            },
            {
              "account": "eosio",
              "name": "delegatebw",
              "authorization": [
                {
                  "actor": "1freeaccount",
                  "permission": "active"
                }
              ],
              "data": {
                "from": "1freeaccount",
                "receiver": "robinzhao123",
                "stake_net_quantity": "0.0100 EOS",
                "stake_cpu_quantity": "0.0100 EOS",
                "transfer": 1
              },
              "hex_data": "90a7a60819a5ee0a3044a0a6fde90ebd640000000000000004454f5300000000640000000000000004454f530000000001"
            }
          ],
          "transaction_extensions": []
        }
      }
    }
  ],
  "block_extensions": [],
  "id": "0038dd029ec7bdfc177fb4470c4c4084ee94f225c3b335e2e3cb37f89b990461",
  "block_num": 3726594,
  "ref_block_prefix": 1203011351
}`
	toStruct(str)
}