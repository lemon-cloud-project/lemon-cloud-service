package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lemon-cloud-project/lemon-cloud-service/define"
	"github.com/lemon-cloud-project/lemon-cloud-service/model"
	"github.com/lemon-cloud-project/lemon-cloud-service/utils"
	"sync"
	"time"
)

type AuthService struct {
}

var authServiceInstance *AuthService
var authServiceInitOnce sync.Once

func Auth() *AuthService {
	authServiceInitOnce.Do(func() {
		authServiceInstance = &AuthService{}
	})
	return authServiceInstance
}

func (i *AuthService) GenerateJwtTokenStr(userKey string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, i.generateJwtPayload(userKey))
	//tokenString, _ := token.SignedString(sysinfo.GetHmacKeyBytes())
	tokenString, _ := token.SignedString("sysinfo.GetHmacKeyBytes()")
	return tokenString
}

func (i *AuthService) generateJwtPayload(userKey string) model.TokenPayload {
	expireDur, _ := time.ParseDuration(fmt.Sprintf("%dm", 1))
	return model.TokenPayload{
		Id:        utils.String().Uuid(true),
		Issuer:    userKey,
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(expireDur).Unix(),
		Audience:  userKey,
		Subject:   define.AppInfo().GetName(),
	}
}

func (i *AuthService) CheckToken(jwtTokenStr string) bool {
	//jwtToken, err := jwt.Parse(jwtTokenStr, func(token *jwt.Token) (i interface{}, e error) {
	//	return "sysinfo.GetHmacKeyBytes()", nil
	//})
	//if jwtToken == nil {
	//	return false
	//}
	//userKey := jwtToken.Claims.(jwt.MapClaims)["iss"]
	//user := i.userDao.FirstByExample(&entity.User{UserKey: userKey.(string)})
	//// user not found or have error
	//if user.UserKey == "" || err != nil {
	//	return false
	//}
	//if _, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
	//	return true
	//} else {
	//	return false
	//}
	return false
}
