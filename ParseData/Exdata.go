/*
 * @Author: ysjyx7
 * @Date: 2022-10-14 17:04:19
 * @LastEditors: ysjyx7
 * @LastEditTime: 2022-10-14 21:06:46
 * @Description: Parse the data in etherscan
 */
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
)

//type Header struct {
//	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
//	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
//	Coinbase    common.Address `json:"miner"            gencodec:"required"`
//	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
//	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
//	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
//	Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
//	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
//	Number      *big.Int       `json:"number"           gencodec:"required"`
//	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
//	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
//	Time        uint64         `json:"timestamp"        gencodec:"required"`
//	Extra       []byte         `json:"extraData"        gencodec:"required"`
//	MixDigest   common.Hash    `json:"mixHash"`
//	Nonce       BlockNonce     `json:"nonce"`

var (
	putf = fmt.Printf
	puts = fmt.Println

	upNum       = 0
	endNum      = 139999999
	dbPath      = "/media/liangyf/Expansion1/xethdata/geth/chaindata"
	ancientPath = dbPath + "/ancient"
)

func main() {
	ancientDb, err := rawdb.NewLevelDBDatabaseWithFreezer(dbPath, 16, 1, ancientPath, "", false)
	if err != nil {
		panic(err)
	}

	// ReadHeadHeaderHash retrieves the hash of the current canonical head header.
	//currHeader := rawdb.ReadHeadHeaderHash(ancientDb)
	//fmt.Printf("currHeader: %x\n", currHeader)

	// ReadHeaderNumber returns the header number assigned to a hash.
	//currHiehgt := rawdb.ReadHeaderNumber(ancientDb, currHeader)
	//fmt.Printf("currHiehgt: %d\n", currHiehgt)

	//puts("----------------------------------------------------------------")
	//cnt:=0
	//nextcnt:=1000000
	for i := upNum; i <= endNum; i++ {
		// ReadCanonicalHash retrieves the hash assigned to a canonical block number.
		blkHash := rawdb.ReadCanonicalHash(ancientDb, uint64(i))

		// hash := rawdb.ReadAllHashes(ancientDb, uint64(i))

		//fmt.Printf("etherscan url: https://etherscan.io/block/%v\n", i)

		if blkHash == (common.Hash{}) {
			fmt.Printf("i: %v\n", i)
		} else {
			fmt.Printf("blkHash: %x\n", blkHash)
		}

		// ReadBody retrieves the block body corresponding to the hash.
		blkHeader := rawdb.ReadHeader(ancientDb, blkHash, uint64(i))
		fmt.Printf("blkHeader Coinbase: 0x%x\n", blkHeader.Coinbase)
		fmt.Printf("blkHeader Time: %d\n", blkHeader.Time)
		fmt.Printf("blkHeader GasUsed: %d\n", blkHeader.GasUsed)
		fmt.Printf("blkHeader GasLimit: %d\n", blkHeader.GasLimit)

		// ReadBody retrieves the block body corresponding to the hash.
		blkBody := rawdb.ReadBody(ancientDb, blkHash, uint64(i))
		fmt.Printf("blkBody: %v\n", blkBody)
		fmt.Printf("blkBody Uncles size: %x\n", len(blkBody.Uncles))
		for _, uncle := range blkBody.Uncles {
			fmt.Printf("uncle Hash: 0x%x\n", uncle.Hash())
		}

		//		fmt.Printf("blkBody Tx size: %x\n", len(blkBody.Transactions))
		// for _, tx := range blkBody.Transactions {
		// 	putf("tx Hash: 0x%x\n", tx.Hash())
		// 	putf("tx from addr: 0x%x\n", getFromAddr(tx))
		// 	putf("tx To: 0x%x\n", tx.To())
		// }

		// ReadBlock retrieves an entire block corresponding to the hash
		block := rawdb.ReadBlock(ancientDb, blkHash, uint64(i))
		fmt.Printf("block hash: 0x%x\n", block.Hash())

		puts("----------------------------------------------------------------")
	}

}

func getFromAddr(tx *types.Transaction) common.Address {
	var signer types.Signer = types.FrontierSigner{}

	from, err := types.Sender(signer, tx)
	if err != nil {
		panic(err)
	}

	return from
}
