package DTO

import "XETH/model"

type CreateBlockTransactionDTO struct {
	Id                   uint32 `json:"id"`
	BlockNum             int64  `json:"blockNum"`
	TimeStamp            int64  `json:"timeStamp"`
	TransactionHash      string `json:"transactionHash"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	ToCreate             string `json:"toCreate"`
	FromIsContract       int8   `json:"fromIsContract"`
	ToIsContract         int8   `json:"toIsContract"`
	Value                int64  `json:"value"`
	GasLimit             int32  `json:"gasLimit"`
	GasPrice             int32  `json:"gasPrice"`
	GasUsed              int32  `json:"gasUsed"`
	CallingFunction      string `json:"callingFunction"`
	IsError              string `json:"isError"`
	Eip2718type          string `json:"eip2718Type"`
	BaseFeePerGas        int64  `json:"baseFeePerGas"`
	MaxFeePerGas         int64  `json:"maxFeePerGas"`
	MaxPriorityFeePerGas int64  `json:"maxPriorityFeePerGas"`
}

func BlockTransactionDTOFromGorm(g *model.BlockTransaction) CreateBlockTransactionDTO {
	return CreateBlockTransactionDTO{
		Id:                   g.Id,
		BlockNum:             g.BlockNum,
		TimeStamp:            g.TimeStamp,
		TransactionHash:      g.TransactionHash,
		From:                 g.From,
		To:                   g.To,
		ToCreate:             g.ToCreate,
		FromIsContract:       g.FromIsContract,
		ToIsContract:         g.ToIsContract,
		Value:                g.Value,
		GasLimit:             g.GasLimit,
		GasPrice:             g.GasPrice,
		GasUsed:              g.GasUsed,
		CallingFunction:      g.CallingFunction,
		IsError:              g.IsError,
		Eip2718type:          g.Eip2718type,
		BaseFeePerGas:        g.BaseFeePerGas,
		MaxFeePerGas:         g.MaxFeePerGas,
		MaxPriorityFeePerGas: g.MaxPriorityFeePerGas,
	}
}

func BlockTransactionDTOS(gs []*model.BlockTransaction) []CreateBlockTransactionDTO {
	ret := make([]CreateBlockTransactionDTO, 0)
	size := len(gs)
	for i := 0; i < size; i++ {
		ret = append(ret, BlockTransactionDTOFromGorm(gs[i]))
	}
	return ret
}
