package lib

import (
	"time"
)
import "github.com/golang-jwt/jwt/v5"

func NewAccessToken(userID string) (string, error) {

	/*	Стандарт RFC-7519
		iss — (issuer)      издатель токена (идентификатор текущего приложения)
		sub — (subject)     "тема", назначение токена
		aud — (audience)    аудитория, получатели токена
		exp — (expire time) срок действия токена
		nbf — (not before)  срок, до которого токен не действителен
		iat — (issued at)   время создания токена
		jti — (JWT id)      идентификатор токена
		Ни одно из значений не является обязательным.
	*/

	// TODO: exp задавать переменной окружения.
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"iss":     "IAA",
			"sub":     "auth",
			"exp":     time.Now().Add(time.Hour * 1).Unix(),
			"iat":     time.Now().Unix(),
			"user_id": userID,
		},
	)
	// TODO: jwt_secret задавать переменной окружения.
	accessToken, err := token.SignedString([]byte("jwt_secret"))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
