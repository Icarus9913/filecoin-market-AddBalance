package main

import (
	"context"
	"filecoin-market-AddBalance/util"
	"fmt"
	"net/http"
	"filecoin-market-AddBalance/model"

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/builtin/market"
	"github.com/filecoin-project/go-address"
	rpc "github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/filecoin-project/lotus/chain/types"
)

type Admin struct {
	Conf *model.Config
	Api *apistruct.FullNodeStruct
}

func main() {
	var admin Admin
	conf, err := util.InitConfig()
	if nil!=err{
		panic(err)
	}
	admin.Conf = conf

	admin.NewConnect()
	msg,err := admin.CombineMsg()
	if nil!=err{
		panic(err)
	}
	admin.Push(msg)
}

func (a *Admin)NewConnect()  {
	var api apistruct.FullNodeStruct

	header := http.Header{
		"Authorization": []string{"Bearer " + a.Conf.AuthToken},
	}

	_, err := rpc.NewMergeClient(
		context.Background(),
		"ws://"+a.Conf.Addr+"/rpc/v0",
		"Filecoin",
		[]interface{}{&api.Internal, &api.CommonStruct.Internal},
		header)
	if nil != err {
		panic(err)
	}
	//defer closer()
	a.Api = &api
}

func (a *Admin)CombineMsg() (*types.Message,error) {
	wallet, err := address.NewFromString(a.Conf.Wallet)
	if nil != err {
		return nil,err
	}
	var amt = abi.NewTokenAmount(int64(1000000000000000000 * a.Conf.FilValue))
	params, err := actors.SerializeParams(&wallet)
	if nil != err {
		return nil,err
	}
	msg := &types.Message{
		To:     market.Address,
		From:   wallet,
		Value:  amt,
		Method: market.Methods.AddBalance,
		Params: params,
	}
	return msg,nil
}

func (a *Admin)Push(msg *types.Message)  {
	message, err := a.Api.MpoolPushMessage(context.Background(), msg, nil)
	if nil!=err{
		panic(err)
	}
	fmt.Println(message.Message)
}