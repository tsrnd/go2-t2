package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/shared/utils"
)

// JWTObject interface
type JWTObject interface {
	GetCustomClaims() map[string]interface{}
	GetIdentifier() uint64
}

// GenerateToken Generate token
func GenerateToken(object JWTObject) (accessToken string, err error) {
	if object == nil {
		err = utils.ErrorsNew("Object is nil")
		return
	}
	emptyID := uint64(0)
	if object.GetIdentifier() == emptyID {
		err = utils.ErrorsNew("Object is empty")
		return
	}
	exp := infrastructure.GetConfigInt64("jwt.claim.exp")
	issuer := infrastructure.GetConfigString("jwt.claim.issuer")
	customClaims := object.GetCustomClaims()
	standardClaims := jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Duration(exp) * time.Second).Unix(),
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
	}
	pay := payload{
		StandardClaims: standardClaims,
		Context:        customClaims,
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = pay
	key := infrastructure.GetConfigByte("jwt.key")
	accessToken, err = token.SignedString(key)
	return
}
