package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	GoodsCategoryVerify           = Rules{"CategoryRank": {NotEmpty()}, "CategoryName": {NotEmpty()}}
	AdminUserRegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	MallUserRegisterVerify        = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	AdminUserChangePasswordVerify = Rules{"Password": {NotEmpty()}}
	GoodsAddParamVerify           = Rules{"GoodsName": {Le("128")}, "GoodsIntro": {Le("200")}, "GoodsCategoryId": {Ge("1")}, "GoodsCoverImg": {NotEmpty()}, "OriginalPrice": {Ge("1"), Le("1000000")},
		"sellingPrice": {Ge("1"), Le("1000000")}, "stockNum": {Ge("1"), Le("100000")}, "Tag": {Le("16")}, "goodsDetailContent": {NotEmpty()}}
	CarouselAddParamVerify       = Rules{"CarouselUrl": {NotEmpty()}, "RedirectUrl": {NotEmpty()}, "CarouselRank": {NotEmpty(), Ge("0"), Le("200")}}
	IndexConfigAddParamVerify    = Rules{"ConfigName": {NotEmpty()}, "ConfigType": {Ge("1"), Le("5")}, "GoodsId": {NotEmpty()}, "ConfigRank": {Ge("1"), Le("200")}}
	IndexConfigUpdateParamVerify = Rules{"ConfigId": {NotEmpty()}, "ConfigName": {NotEmpty()}, "ConfigType": {Ge("1"), Le("5")}, "GoodsId": {NotEmpty()}, "ConfigRank": {Ge("1"), Le("200")}}
	SaveOrderParamVerify         = Rules{"CartItemIds": {NotEmpty()}, "AddressId": {NotEmpty()}}
)

type Claim struct {
	Name string
	jwt.StandardClaims
}

var jwtKey = []byte("undergrad-thesis-prject-NP_123")

func CreateToken(name string, expireDuration time.Duration) (string, error) {
	claims := &Claim{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "NP_123",
			Subject:   "User_Token",
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func VerifyToken(token string) (string, bool) {
	if token == "" {
		return "", false
	}

	tok, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println("ParseWithClaims error %v", err)
		return "", false
	}

	if claims, ok := tok.Claims.(*Claim); ok && tok.Valid {
		return claims.Name, true
	} else {
		fmt.Println("%v", err)
		return "", false
	}

}
