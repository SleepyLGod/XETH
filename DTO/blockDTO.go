package DTO

import "XETH/model"

type CreateBlockDTO struct {
	BlockNum                int64  `json:"blockNum"`
	TimeStamp               int64  `json:"timeStamp"`
	TransactionCount        int32  `json:"transactionCount"`
	InternalTransctionCount int32  `json:"internalTransctionCount"`
	MinerAddress            string `json:"minerAddress"`
	BlockReward             int64  `json:"blockReward"`
	UnclesReward            string `json:"unclesReward"` // TODO: 待检测string类型是否正确
	Difficulty              string `json:"difficulty"`
	TotalDifficulty         string `json:"totalDifficulty"`
	Size                    int32  `json:"size"`
	GasUsed                 int32  `json:"gasUsed"`
	GasLimit                int32  `json:"gasLimit"`
	BaseFeePerGas           int64  `json:"baseFeePerGas"`
	BurntFees               int64  `json:"burntFees"`
	ExtraData               string `json:"extraData"`
	Hash                    string `json:"hash"`
	ParentHash              string `json:"parentHash"`
	Sha3Uncles              string `json:"sha3Uncles"`
	StateRoot               string `json:"stateRoot"`
	Nounce                  string `json:"nounce"`
}

func BlockDTOFromGorm(g *model.Block) CreateBlockDTO {
	return CreateBlockDTO{
		BlockNum:                g.BlockNum,
		TimeStamp:               g.TimeStamp,
		TransactionCount:        g.TransactionCount,
		InternalTransctionCount: g.InternalTransactionCount,
		MinerAddress:            g.MinerAddress,
		BlockReward:             g.BlockReward,
		UnclesReward:            g.UnclesReward,
		Difficulty:              g.Difficulty,
		TotalDifficulty:         g.TotalDifficulty,
		Size:                    g.Size,
		GasUsed:                 g.GasUsed,
		GasLimit:                g.GasLimit,
		BaseFeePerGas:           g.BaseFeePerGas,
		BurntFees:               g.BurntFees,
		ExtraData:               g.ExtraData,
		Hash:                    g.Hash,
		ParentHash:              g.ParentHash,
		Sha3Uncles:              g.Sha3Uncles,
		StateRoot:               g.StateRoot,
		Nounce:                  g.Nounce,
	}
}

func BlockDTOS(gs []*model.Block) []CreateBlockDTO {
	ret := make([]CreateBlockDTO, 0)
	size := len(gs)
	for i := 0; i < size; i++ {
		ret = append(ret, BlockDTOFromGorm(gs[i]))
	}
	return ret
}
