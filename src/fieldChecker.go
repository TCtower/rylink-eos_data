package main

import(
	"fmt"
	"encoding/json"
	"reflect"
)
func fieldCheck(blockInfo string){
	for i := 1; i <= 1000; i++{
		tmp := getBlockInfo(i)
		var Map map[string]interface{}
		if err := json.Unmarshal([]byte(tmp), &Map); err != nil {
			panic(err)
		}

		block := ToStruct(tmp)
		cntA := reflect.ValueOf(block).NumField()
		cntB := reflect.ValueOf(Map).Len()

		if (cntB == 3){i--; continue;}

		fmt.Print(i)
		fmt.Print(" ", cntA)
		fmt.Print(" ", cntB)
		if (cntA != cntB){
			fmt.Print("Strange value");
			break;
		}

		fmt.Print(" ", len(block.Transactions))
		fmt.Print(" ", reflect.ValueOf(Map["transactions"]).Len())
		trans := reflect.ValueOf(Map["transactions"]);
		if trans.Len() > 0 {
			fmt.Print(trans.Index(0))
		}

		fmt.Println("")
	}

}