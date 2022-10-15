package service

import (
	dto "XETH/DTO"
	core "XETH/dbDriver"
	"XETH/model"
)

// CreateContractInfoServiceWithDTO 推荐
func CreateContractInfoServiceWithDTO(createContractInfo dto.CreateContractInfoDTO) bool {
	db := core.GetDB()
	ctt := model.ContractInfo{
		Id:                     createContractInfo.Id,
		CreatedBlock:           createContractInfo.CreatedBlock,
		CreatedTimeStamp:       createContractInfo.CreatedTimeStamp,
		Address:                createContractInfo.Address,
		CreatedTransactionHash: createContractInfo.CreatedTransactionHash,
		Creator:                createContractInfo.Creator,
		CreatorIsContract:      createContractInfo.CreatorIsContract,
		CreateValue:            createContractInfo.CreateValue,
		CreationCode:           createContractInfo.CreationCode,
		ContractCode:           createContractInfo.ContractCode,
	}
	tx := db.Create(&ctt)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetContractInfosService() []dto.CreateContractInfoDTO {
	var schema []*model.ContractInfo
	core.GetDB().Find(&schema)
	return dto.ContractInfoDTOS(schema)
}

func GetContractInfosWithConstraintsService(constraints []core.QueryConstraint) []dto.CreateContractInfoDTO {
	var schema []*model.ContractInfo
	core.QueryWithDb(core.GetDB(), &schema, constraints)
	return dto.ContractInfoDTOS(schema)
}

func DeleteContractInfoByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.ContractInfo{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.ContractInfo{}, id).Error
	return err, true
}

func GetContractInfoByIdService(id int64) dto.CreateContractInfoDTO {
	var schema *model.ContractInfo
	core.GetDB().First(&schema, id)
	return dto.ContractInfoDTOFromGorm(schema)
}
