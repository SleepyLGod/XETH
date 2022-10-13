package DTO

type Block struct {
	BlockNum                int64
	Timestamp               int64
	TransactionCount        int32
	InternalTransctionCount int32
	MinerAddress            string
	BlockReward             int64
	UnclesReward            string // TODO: 待检测string类型是否正确
	Difficulty              string
	TotalDifficulty         string
	Size                    int32
	GasUsed                 int32
	GasLimit                int32
	BaseFeePerGas           int64
	BurntFees               int64
	ExtraData               string
	Hash                    string
	ParentHash              string
	Sha3Uncles              string
	StateRoot               string
	Nounce                  string
}
