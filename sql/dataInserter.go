package sql

import (
	dto "XETH/DTO"
	"XETH/service"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// InsertBlockFromFile 从blockInfoFilePath加载数据，写入数据库
func InsertBlockFromFile(blockInfoFilePath string, blockRewardFilePath string) error {

	infoFile, err := os.OpenFile(blockInfoFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer infoFile.Close()

	// BlockReward和其他数据不在同一文件中，需要从两个文件中读入
	rewardFile, err := os.OpenFile(blockRewardFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer rewardFile.Close()

	infoBuf := bufio.NewReader(infoFile)
	rewardBuf := bufio.NewReader(rewardFile)

	// 忽略表头
	_, err = infoBuf.ReadString('\n')
	if err != nil {
		return nil
	}
	_, err = rewardBuf.ReadString('\n')
	if err != nil {
		return nil
	}
	cnt := 0
	for {
		infoLine, err := infoBuf.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		rewardLine, err := rewardBuf.ReadString('\n')
		if err != nil {
			return err
		}
		cnt += 1
		fmt.Println("!!!!!!!!!!!!!!!!!", cnt)
		splits := strings.Split(infoLine, ",")
		blockNum, err := strconv.ParseInt(splits[0], 10, 64)
		timeStamp, err := strconv.ParseInt(splits[1], 10, 64)
		txnCount, err := strconv.ParseInt(splits[4], 10, 64)
		itnlTxnCntSimple, err := strconv.ParseInt(splits[5], 10, 64)
		itnlTxnCntAdvanced, err := strconv.ParseInt(splits[6], 10, 64)
		itnlTxnCnt := itnlTxnCntSimple + itnlTxnCntAdvanced
		size, err := strconv.ParseInt(splits[2], 10, 64)
		gasUsed, err := strconv.ParseInt(splits[12], 10, 64)
		gasLimit, err := strconv.ParseInt(splits[11], 10, 64)
		baseFeePerGas, err := strconv.ParseInt(splits[17], 10, 64)
		burntFees, err := strconv.ParseInt(splits[18], 10, 64)

		rewardSplits := strings.Split(rewardLine[:len(rewardLine)-2], ",")
		blockReward, err := strconv.ParseInt(rewardSplits[3], 10, 64)
		block := dto.CreateBlockDTO{
			BlockNum:                blockNum,
			Timestamp:               timeStamp,
			TransactionCount:        int32(txnCount),
			InternalTransctionCount: int32(itnlTxnCnt),
			MinerAddress:            splits[9],
			BlockReward:             blockReward,
			UnclesReward:            "", // TODO: 获取数据
			Difficulty:              splits[3],
			TotalDifficulty:         splits[8],
			Size:                    int32(size),
			GasUsed:                 int32(gasUsed),
			GasLimit:                int32(gasLimit),
			BaseFeePerGas:           baseFeePerGas,
			BurntFees:               burntFees,
			ExtraData:               "", // TODO: 获取数据
			Hash:                    "", // TODO: 获取数据
			ParentHash:              "", // TODO: 获取数据
			Sha3Uncles:              "", // TODO: 获取数据
			StateRoot:               "", // TODO: 获取数据
			Nounce:                  "", // TODO: 获取数据
		}
		if service.CreateBlockServiceWithDTO(block) == false {
			return errors.New("failed to create block")
		}
	}
	return nil
}
