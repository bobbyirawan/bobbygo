package crypto

import (
	"github.com/golang-jwt/jwt"
)

type (
	Crypto interface {
		EncryptHS256(claims jwt.Claims) ([]byte, error)
		DecryptHS256(claims jwt.Claims, text string) (jwt.Claims, error)
	}

	impl struct {
		secret []byte
	}
)

func (i *impl) EncryptHS256(claims jwt.Claims) ([]byte, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	ciphertext, err := token.SignedString(i.secret)
	if err != nil {
		return nil, err
	}

	return []byte(ciphertext), nil
}

func (i *impl) DecryptHS256(claims jwt.Claims, tokenString string) (jwt.Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return i.secret, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}

func New(secret string) (Crypto, error) {
	return &impl{[]byte(secret)}, nil
}
