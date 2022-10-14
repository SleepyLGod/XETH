package DTO

type CreateInternalTransactionDTO struct {
	Id               uint32 `json:"id"`
	BlockNum         int64  `json:"blockNum"`
	TimeStamp        int64  `json:"timeStamp"`
	TransactionHash  string `json:"transactionHash"`
	TypeTraceAddress string `json:"typeTraceAddress"`
	From             string `json:"from"`
	To               string `json:"to"`
	FromIsContract   int8   `json:"fromIsContract"`
	ToIsContract     int8   `json:"toIsContract"`
	Value            int64  `json:"value"`
	CallingFunction  string `json:"callingFunction"`
	IsError          string `json:"isError"`
}
