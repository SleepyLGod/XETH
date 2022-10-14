package DTO

import "XETH/model"

type CreateContractInfoDTO struct {
	Id                     uint32 `json:"id"`
	CreatedBlock           int64  `json:"createdBlock"`
	CreatedTimeStamp       int64  `json:"createdTimeStamp"`
	Address                string `json:"address"`
	CreatedTransactionHash string `json:"createdTransactionHash"`
	Creator                string `json:"creator"`
	CreatorIsContract      int8   `json:"creatorIsContract"`
	CreateValue            int64  `json:"createValue"`
	CreationCode           string `json:"creationCode"`
	ContractCode           string `json:"contractCode"`
}

func ContractInfoDTOFromGorm(g *model.ContractInfo) CreateContractInfoDTO {
	return CreateContractInfoDTO{
		Id:                     g.Id,
		CreatedBlock:           g.CreatedBlock,
		CreatedTimeStamp:       g.CreatedTimeStamp,
		Address:                g.Address,
		CreatedTransactionHash: g.CreatedTransactionHash,
		Creator:                g.Creator,
		CreatorIsContract:      g.CreatorIsContract,
		CreateValue:            g.CreateValue,
		CreationCode:           g.CreationCode,
		ContractCode:           g.ContractCode,
	}
}

func ContractInfoDTOS(gs []*model.ContractInfo) []CreateContractInfoDTO {
	ret := make([]CreateContractInfoDTO, 0)
	size := len(gs)
	for i := 0; i < size; i++ {
		ret = append(ret, ContractInfoDTOFromGorm(gs[i]))
	}
	return ret
}
