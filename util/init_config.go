package util

import (
	"encoding/json"
	"filecoin-market-AddBalance/model"
	"os"
)

func InitConfig() (*model.Config,error) {
	var conf *model.Config
	file, err := os.Open("config.json")
	if nil!=err{
		return nil, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if nil!=err{
		return nil,err
	}
	return conf,nil
}