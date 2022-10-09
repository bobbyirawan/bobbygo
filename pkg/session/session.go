package session

import (
	"bobbygo/pkg/crypto"
	"bobbygo/pkg/errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type (
	Session struct {
		UserID   string `json:"userID" validate:"required"`
		Username string `json:"username" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Role     int    `json:"role" validate:"required"`
		Expired  int64  `json:"exp" validate:"required"`
	}
)

func (ss *Session) IsSessionExpired() error {
	if time.Now().After(time.Unix(ss.Expired, 0)) {
		return errors.ErrExpiredSession
	}

	return nil
}

func (ss *Session) ExtendSession(cr crypto.Crypto, duration int64) (string, error) {
	ss.Expired = time.Now().Add(time.Duration(duration) * time.Second).Unix()

	return ss.Encrypt(cr)
}

func (ss *Session) Encrypt(cr crypto.Crypto) (string, error) {
	enc, _ := cr.EncryptHS256(ss)

	return string(enc), nil
}

func (ss *Session) Valid() error {
	return nil
}

func NewSession(cr crypto.Crypto, session string) (*Session, error) {
	var (
		ss       = &Session{}
		dec, err = cr.DecryptHS256(ss, session)
	)

	if err != nil {
		return nil, errors.ErrSession.WithUnderlyingErrors(err)
	}

	if err := validator.New().Struct(dec); err != nil {
		return nil, errors.ErrSession.WithUnderlyingErrors(err)
	}

	if err := ss.IsSessionExpired(); err != nil {
		return nil, err
	}

	return ss, nil
}
