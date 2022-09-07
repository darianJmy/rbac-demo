package httputils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId int
	jwt.StandardClaims
}

func GenerateToken(uid int, jwtkey []byte) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtkey)
}

func ParseToken(token string, jwtkey []byte) (*Claims, error) {
	var clams Claims
	t, err := jwt.ParseWithClaims(token, &clams, func(t *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})

	if err != nil || !t.Valid {
		return nil, err
	}
	return &clams, err
}
