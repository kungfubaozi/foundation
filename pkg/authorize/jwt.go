package authorize

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
	"time"
	"zskparker.com/foundation/base/authenticate/pb"
)

var TOKEN_KEY = []byte("1DMx*o34&42j/B?+6R#M3qr40]$3W")

type UserClaims struct {
	Token *fs_base_authenticate.Authorize

	jwt.StandardClaims
}

func DecodeToken(token string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return TOKEN_KEY, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*UserClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, err
}

//加密token
func EncodeToken(authorize *fs_base_authenticate.Authorize) (string, error) {
	//3天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()

	c := jwt.StandardClaims{
		Issuer:    "foundation.authenticate", // 签发者
		ExpiresAt: expireTime,
		Id:        uuid.New(),
	}

	claims := UserClaims{
		Token:          authorize,
		StandardClaims: c,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(TOKEN_KEY)

}
