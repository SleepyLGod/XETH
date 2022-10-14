package DTO

type CreateBlockTransactionDTO struct {
	BlockNum             int64  `json:"blockNum"`
	Timestamp            int64  `json:"timestamp"`
	TransactionHash      string `json:"transactionHash"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	ToCreate             string `json:"toCreate"`
	FromIsContract       int8   `json:"fromIsContract"`
	ToIsContract         int8   `json:"toIsContract"`
	Value                int64  `json:"value"`
	GasLimit             int32  `json:"gasLimit"`
	GasPrice             int64  `json:"gasPrice"`
	GasUsed              int32  `json:"gasUsed"`
	CallingFunction      string `json:"callingFunction"`
	IsError              int8   `json:"isError"`
	Eip2718type          string `json:"eip2718Type"`
	BaseFeePerGas        int64  `json:"baseFeePerGas"`
	MaxFeePerGas         int64  `json:"maxFeePerGas"`
	MaxPriorityFeePerGas int64  `json:"maxPriorityFeePerGas"`
}
