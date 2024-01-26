package mall

import (
	"errors"
	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	"time"
)

type MallUserAddressService struct {
}

// GetMyAddress 获取收货地址
func (m *MallUserAddressService) GetMyAddress(userID string) (err error, userAddress []mall.MallUserAddress) {
	global.GVA_DB.Where("user_id=? and is_deleted=0", userID).Find(&userAddress)
	return
}

// SaveUserAddress 保存用户地址
func (m *MallUserAddressService) SaveUserAddress(userID int, req mallReq.AddAddressParam) (err error) {
	var defaultAddress mall.MallUserAddress
	copier.Copy(&defaultAddress, &req)
	defaultAddress.CreateTime = common.JSONTime{Time: time.Now()}
	defaultAddress.UpdateTime = common.JSONTime{Time: time.Now()}
	defaultAddress.UserId = userID
	if req.DefaultFlag == 1 {
		global.GVA_DB.Where("user_id=? and default_flag =1 and is_deleted = 0", userID).First(&defaultAddress)
		if defaultAddress != (mall.MallUserAddress{}) {
			defaultAddress.UpdateTime = common.JSONTime{time.Now()}
			err = global.GVA_DB.Save(&defaultAddress).Error
			if err != nil {
				return
			}
		}
	} else {
		err = global.GVA_DB.Create(&defaultAddress).Error
		if err != nil {
			return
		}
	}
	return
}

// UpdateUserAddress 更新用户地址
func (m *MallUserAddressService) UpdateUserAddress(userID int, req mallReq.UpdateAddressParam) (err error) {
	var userAddress mall.MallUserAddress
	if err = global.GVA_DB.Where("address_id =? and user_id =?", req.AddressId, userID).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址")
	}
	if userID != userAddress.UserId {
		return errors.New("禁止该操作！")
	}
	if req.DefaultFlag == 1 {
		var defaultUserAddress mall.MallUserAddress
		global.GVA_DB.Where("user_id=? and default_flag =1 and is_deleted = 0", userID).First(&defaultUserAddress)
		if defaultUserAddress != (mall.MallUserAddress{}) {
			defaultUserAddress.DefaultFlag = 0
			defaultUserAddress.UpdateTime = common.JSONTime{time.Now()}
			err = global.GVA_DB.Save(&defaultUserAddress).Error
			if err != nil {
				return
			}
		}
	}
	err = copier.Copy(&userAddress, &req)
	if err != nil {
		return
	}
	userAddress.UpdateTime = common.JSONTime{time.Now()}
	userAddress.UserId = userID
	err = global.GVA_DB.Save(&userAddress).Error
	return
}

func (m *MallUserAddressService) GetMallUserAddressById(userID int, id int) (err error, userAddress mall.MallUserAddress) {
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址"), userAddress
	}
	if userID != userAddress.UserId {
		return errors.New("禁止该操作！"), userAddress
	}
	return
}

func (m *MallUserAddressService) GetMallUserDefaultAddress(userID int) (err error, userAddress mall.MallUserAddress) {
	if err = global.GVA_DB.Where("user_id =? and default_flag =1 and is_deleted = 0 ", userID).First(&userAddress).Error; err != nil {
		return errors.New("不存在默认地址失败"), userAddress
	}
	return
}

func (m *MallUserAddressService) DeleteUserAddress(userID int, id int) (err error) {
	var userAddress mall.MallUserAddress
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址")
	}
	if userID != userAddress.UserId {
		return errors.New("禁止该操作！")
	}
	err = global.GVA_DB.Delete(&userAddress).Error
	return

}
