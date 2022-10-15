package service

import (
	dto "XETH/DTO"
	core "XETH/dbDriver"
	"XETH/model"
)

// CreateTokenInfoServiceWithDTO 推荐
func CreateTokenInfoServiceWithDTO(createTokenInfo dto.CreateTokenInfoDTO) bool {
	db := core.GetDB()
	tokenInfo := model.TokenInfo{
		Id:          createTokenInfo.Id,
		Address:     createTokenInfo.Address,
		Name:        createTokenInfo.Name,
		Symbol:      createTokenInfo.Symbol,
		TotalSupply: createTokenInfo.TotalSupply,
		Decimal:     createTokenInfo.Decimal,
	}
	tx := db.Create(&tokenInfo)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetTokenInfosService() []dto.CreateTokenInfoDTO {
	var schema []*model.TokenInfo
	core.GetDB().Find(&schema)
	return dto.TokenInfoDTOS(schema)
}

func GetTokenInfosWithConstraintsService(constraints []core.QueryConstraint) []dto.CreateTokenInfoDTO {
	var schema []*model.TokenInfo
	core.QueryWithDb(core.GetDB(), &schema, constraints)
	return dto.TokenInfoDTOS(schema)
}

func DeleteTokenInfoByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.TokenInfo{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.TokenInfo{}, id).Error
	return err, true
}

func GetTokenInfoByIdService(id int64) dto.CreateTokenInfoDTO {
	var schema *model.TokenInfo
	core.GetDB().First(&schema, id)
	return dto.TokenInfoDTOFromGorm(schema)
}
