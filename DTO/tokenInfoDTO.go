package DTO

import "XETH/model"

type CreateTokenInfoDTO struct {
	Id          uint32 `json:"id"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	TotalSupply string `json:"totalSupply"`
	Decimal     int32  `json:"decimal"`
}

func TokenInfoDTOFromGorm(g *model.TokenInfo) CreateTokenInfoDTO {
	return CreateTokenInfoDTO{
		Id:          g.Id,
		Address:     g.Address,
		Name:        g.Name,
		Symbol:      g.Symbol,
		TotalSupply: g.TotalSupply,
		Decimal:     g.Decimal,
	}
}

func TokenInfoDTOS(gs []*model.TokenInfo) []CreateTokenInfoDTO {
	ret := make([]CreateTokenInfoDTO, 0)
	size := len(gs)
	for i := 0; i < size; i++ {
		ret = append(ret, TokenInfoDTOFromGorm(gs[i]))
	}
	return ret
}
