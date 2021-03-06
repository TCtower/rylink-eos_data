package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"

)
type requestBlcok 	struct {
	Block_num_or_id int `json:"block_num_or_id"`
}

func getBlockInfo(blockNum int) string{
	//get block info
	requestBody := &requestBlcok{
		blockNum,
	}
	//fmt.Print(requestBody.Block_num_or_id)
	requestStr, err := json.Marshal(requestBody)
	if err != nil{
		fmt.Print("failed to endcode json")
	}

	//fmt.Println(string(requestStr))

	body := bytes.NewBuffer([]byte(requestStr))
	res,err := http.Post("http://172.17.2.67:8888/v1/chain/get_block", "application/json;charset=utf-8", body)
	if err != nil {
		fmt.Print("Failed to post data")
		return "Fail"
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print("Failed to get response")
		return "Fail"
	}

	res.Body.Close()
	return string(result);
}

func main() {
	client := &http.Client{}

	//get info
	request, _ := http.NewRequest("GET", "http://172.17.2.67:8888/v1/chain/get_info", nil)
	//receive
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(bodystr), &dat); err == nil {
			//fmt.Println(dat)
			block_num := dat["head_block_num"].(float64)
			fmt.Println(int(block_num))
		}

		fmt.Println(bodystr)
	}

	//fieldCheck(getBlockInfo(189))

}
