package DTO

type CreateTokenInfoDTO struct {
	Id          uint32 `json:"id"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	TotalSupply string `json:"totalSupply"`
	Decimal     int32  `json:"decimal"`
}
