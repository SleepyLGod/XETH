package DTO

type CreateBlockDTO struct {
	BlockNum                int64  `json:"blockNum"`
	Timestamp               int64  `json:"timestamp"`
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
