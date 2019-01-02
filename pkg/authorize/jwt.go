package authorize

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
	"time"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/pb"
)

var TOKEN_KEY = []byte("1DMx*o34&42j/B?+6R#M3qr40]$3W")

type UserClaims struct {
	Token *authenticate.SimpleAuthorize

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

func EncodeAccessToken(authorize *fs_base_authenticate.Authorize) (string, error) {
	sa := &authenticate.SimpleAuthorize{
		UserId:    authorize.UserId,
		ProjectId: authorize.ProjectId,
		ClientId:  authorize.UserId,
		TokenAb:   uuid.New(),
		Access:    true,
	}
	return encodeToken(time.Minute*10, sa)
}

func EncodeRefreshToken(authorize *fs_base_authenticate.Authorize) (string, error) {
	sa := &authenticate.SimpleAuthorize{
		UserId:    authorize.UserId,
		ProjectId: authorize.ProjectId,
		ClientId:  authorize.UserId,
		TokenAb:   uuid.New(),
		Access:    false,
	}
	return encodeToken(time.Hour*24*7, sa)
}

//加密token
func encodeToken(et time.Duration, authorize *authenticate.SimpleAuthorize) (string, error) {
	expireTime := time.Now().Add(et).Unix()

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
