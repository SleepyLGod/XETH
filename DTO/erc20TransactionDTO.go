package DTO

import "XETH/model"

type CreateERC20TransactionDTO struct {
	Id              uint32 `json:"id"`
	TimeStamp       int64  `json:"timeStamp"`
	BlockNum        int64  `json:"blockNum"`
	TransactionHash string `json:"transactionHash"`
	TokenAddress    string `json:"tokenAddress"`
	From            string `json:"from"`
	To              string `json:"to"`
	FromIsContract  int8   `json:"fromIsContract"`
	ToIsContract    int8   `json:"toIsContract"`
	Value           int64  `json:"value"`
}

func ERC20TransactionDTOFromGorm(g *model.ERC20Transaction) CreateERC20TransactionDTO {
	return CreateERC20TransactionDTO{
		Id:              g.Id,
		TimeStamp:       g.TimeStamp,
		BlockNum:        g.BlockNum,
		TransactionHash: g.TransactionHash,
		TokenAddress:    g.TokenAddress,
		From:            g.From,
		To:              g.To,
		FromIsContract:  g.FromIsContract,
		ToIsContract:    g.ToIsContract,
		Value:           g.Value,
	}
}

func ERC20TransactionDTOS(gs []*model.ERC20Transaction) []CreateERC20TransactionDTO {
	ret := make([]CreateERC20TransactionDTO, 0)
	size := len(gs)
	for i := 0; i < size; i++ {
		ret = append(ret, ERC20TransactionDTOFromGorm(gs[i]))
	}
	return ret
}
