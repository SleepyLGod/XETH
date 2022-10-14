package DTO

type CreateContractInfoDTO struct {
	CreatedBlock           int64  `json:"createdBlock"`
	CreatedTimeStamp       int64  `json:"createdTimeStamp"`
	Address                int64  `json:"address"`
	CreatedTransactionHash string `json:"createdTransactionHash"`
	Creator                string `json:"creator"`
	CreatorIsContract      int8   `json:"creatorIsContract"`
	CreateValue            int64  `json:"createValue"`
	CreationCode           string `json:"creationCode"`
	ContractCode           string `json:"contractCode"`
}
