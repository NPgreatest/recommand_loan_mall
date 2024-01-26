package mall

import "main.go/service"

type MallGroup struct {
	MallIndexApi
	MallGoodsInfoApi
	MallGoodsCategoryApi
	MallUserApi
	MallUserAddressApi
	MallShopCartApi
	MallOrderApi
	MallQueryApi
}

var mallCarouselService = service.ServiceGroupApp.MallServiceGroup.MallCarouselService
var mallGoodsInfoService = service.ServiceGroupApp.MallServiceGroup.MallGoodsInfoService
var mallIndexConfigService = service.ServiceGroupApp.MallServiceGroup.MallIndexInfoService
var mallGoodsCategoryService = service.ServiceGroupApp.MallServiceGroup.MallGoodsCategoryService
var mallUserService = service.ServiceGroupApp.MallServiceGroup.MallUserService
var mallUserAddressService = service.ServiceGroupApp.MallServiceGroup.MallUserAddressService
var mallShopCartService = service.ServiceGroupApp.MallServiceGroup.MallShopCartService
var mallOrderService = service.ServiceGroupApp.MallServiceGroup.MallOrderService
var mallQueryService = service.ServiceGroupApp.MallServiceGroup.MallQueryService
