package sql

import (
	dto "XETH/DTO"
	"XETH/service"
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

func InsertBlockInfoFromFile(blockInfoFilePath string) error {
	file, err := os.OpenFile(blockInfoFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	_, err = buf.ReadString('\n')
	if err != nil {
		return nil
	}
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		splits := strings.Split(line, ",")
		blockNum, err := strconv.ParseInt(splits[0], 10, 64)
		timeStamp, err := strconv.ParseInt(splits[1], 10, 64)
		txnCount, err := strconv.ParseInt(splits[4], 10, 64)
		itnlTxnCntSimple, err := strconv.ParseInt(splits[5], 10, 64)
		itnlTxnCntAdvanced, err := strconv.ParseInt(splits[6], 10, 64)
		itnlTxnCnt := itnlTxnCntSimple + itnlTxnCntAdvanced
		blockReward, err := strconv.ParseInt(splits[5], 10, 64)
		size, err := strconv.ParseInt(splits[2], 10, 64)
		gasUsed, err := strconv.ParseInt(splits[12], 10, 64)
		gasLimit, err := strconv.ParseInt(splits[11], 10, 64)
		baseFeePerGas, err := strconv.ParseInt(splits[17], 10, 64)
		burntFees, err := strconv.ParseInt(splits[18], 10, 64)

		block := dto.CreateBlockDTO{
			BlockNum:                blockNum,
			Timestamp:               timeStamp,
			TransactionCount:        int32(txnCount),
			InternalTransctionCount: int32(itnlTxnCnt),
			MinerAddress:            splits[9],
			BlockReward:             blockReward,
			UnclesReward:            splits[6],
			Difficulty:              splits[3],
			TotalDifficulty:         splits[8],
			Size:                    int32(size),
			GasUsed:                 int32(gasUsed),
			GasLimit:                int32(gasLimit),
			BaseFeePerGas:           baseFeePerGas,
			BurntFees:               burntFees,
			ExtraData:               "",
			Hash:                    "",
			ParentHash:              "",
			Sha3Uncles:              "",
			StateRoot:               "",
			Nounce:                  "",
		}
		if service.CreateBlockServiceWithDTO(block) == false {
			return errors.New("failed to create block")
		}
	}
	return nil
}
