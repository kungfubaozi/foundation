package authenticate

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
	"time"
	"zskparker.com/foundation/base/authenticate/pb"
)

var TOKEN_KEY = []byte("e48df34a-0f32-11e9-ab14-d663bd873d93")

type UserClaims struct {
	Token *SimpleAuthorize

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

func encodeAccessToken(authorize *fs_base_authenticate.Authorize) (string, error) {
	sa := &SimpleAuthorize{
		UserId:   authorize.UserId,
		ClientId: authorize.UserId,
		UUID:     authorize.AccessTokenUUID,
		Access:   true,
		Relation: authorize.Relation,
	}
	return encodeToken(time.Minute*10, sa)
}

func encodeRefreshToken(authorize *fs_base_authenticate.Authorize) (string, error) {
	sa := &SimpleAuthorize{
		UserId:   authorize.UserId,
		ClientId: authorize.UserId,
		UUID:     authorize.RefreshTokenUUID,
		Access:   false,
		Relation: authorize.Relation,
	}
	return encodeToken(time.Hour*24*7, sa)
}

//加密token
func encodeToken(et time.Duration, authorize *SimpleAuthorize) (string, error) {
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
