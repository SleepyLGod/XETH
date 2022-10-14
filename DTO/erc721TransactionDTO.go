package DTO

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
