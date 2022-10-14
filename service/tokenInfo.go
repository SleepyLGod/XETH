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

func GetTokenInfosService() (tokenInfoList []*model.TokenInfo) {
	core.GetDB().Find(&tokenInfoList)
	return
}

func GetTokenInfosWithConstraintsService(constraints []core.QueryConstraint) (tokenInfoList []*model.TokenInfo) {
	core.QueryWithDb(core.GetDB(), &tokenInfoList, constraints)
	return
}

func DeleteTokenInfoByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.TokenInfo{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.TokenInfo{}, id).Error
	return err, true
}

func GetTokenInfoByIdService(id int64) (tokenInfo []*model.TokenInfo) {
	core.GetDB().First(&tokenInfo, id)
	return
}
