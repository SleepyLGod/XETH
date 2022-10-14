package DTO

import "XETH/model"

type CreateInternalTransactionDTO struct {
	Id               uint32 `json:"id"`
	BlockNum         int64  `json:"blockNum"`
	TimeStamp        int64  `json:"timeStamp"`
	TransactionHash  string `json:"transactionHash"`
	TypeTraceAddress string `json:"typeTraceAddress"`
	From             string `json:"from"`
	To               string `json:"to"`
	FromIsContract   int8   `json:"fromIsContract"`
	ToIsContract     int8   `json:"toIsContract"`
	Value            int64  `json:"value"`
	CallingFunction  string `json:"callingFunction"`
	IsError          string `json:"isError"`
}

func InternalTransactionDTOFromGorm(g *model.InternalTransaction) CreateInternalTransactionDTO {
	return CreateInternalTransactionDTO{
		Id:               g.Id,
		BlockNum:         g.BlockNum,
		TimeStamp:        g.TimeStamp,
		TransactionHash:  g.TransactionHash,
		TypeTraceAddress: g.TypeTraceAddress,
		From:             g.From,
		To:               g.To,
		FromIsContract:   g.FromIsContract,
		ToIsContract:     g.ToIsContract,
		Value:            g.Value,
		CallingFunction:  g.CallingFunction,
		IsError:          g.IsError,
	}
}

func InternalTransactionDTOS(gs []*model.InternalTransaction) []CreateInternalTransactionDTO {
	ret := make([]CreateInternalTransactionDTO, 0)
	size := len(gs)
	for i := 0; i < size; i++ {
		ret = append(ret, InternalTransactionDTOFromGorm(gs[i]))
	}
	return ret
}
