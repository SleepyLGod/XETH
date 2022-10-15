package DTO

import "XETH/model"

type CreateERC721TransactionDTO struct {
	Id              uint32 `json:"id"`
	TimeStamp       int64  `json:"timeStamp"`
	BlockNum        int64  `json:"blockNum"`
	TransactionHash string `json:"transactionHash"`
	TokenAddress    string `json:"tokenAddress"`
	From            string `json:"from"`
	To              string `json:"to"`
	FromIsContract  int8   `json:"fromIsContract"`
	ToIsContract    int8   `json:"toIsContract"`
	TokenId         uint32 `json:"tokenId"`
}

func ERC721TransactionDTOFromGorm(g *model.ERC721Transaction) CreateERC721TransactionDTO {
	return CreateERC721TransactionDTO{
		Id:              g.Id,
		TimeStamp:       g.TimeStamp,
		BlockNum:        g.BlockNum,
		TransactionHash: g.TransactionHash,
		TokenAddress:    g.TokenAddress,
		From:            g.From,
		To:              g.To,
		FromIsContract:  g.FromIsContract,
		ToIsContract:    g.ToIsContract,
		TokenId:         g.TokenId,
	}
}

func ERC721TransactionDTOS(gs []*model.ERC721Transaction) []CreateERC721TransactionDTO {
	ret := make([]CreateERC721TransactionDTO, 0)
	size := len(gs)
	for i := 0; i < size; i++ {
		ret = append(ret, ERC721TransactionDTOFromGorm(gs[i]))
	}
	return ret
}
