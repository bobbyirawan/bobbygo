package middlewares

import (
	"github.com/bobbygo/pkg/context"
	"github.com/bobbygo/pkg/crypto"
	"github.com/bobbygo/pkg/logger"
	"github.com/bobbygo/pkg/session"
	"github.com/labstack/echo/v4"
)

const (
	AuthorizationHeader = "authorization"
)

type (
	SessionMiddleware interface {
		AuthenticateSession(next echo.HandlerFunc) echo.HandlerFunc
	}
	impl struct {
		secret string
		crypto crypto.Crypto
		logger logger.Logger
	}
)

func (i *impl) AuthenticateSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			sctx  = context.NewEmptyBobbyGoContext(ctx)
			token = ctx.Request().Header.Get(AuthorizationHeader)
		)

		NewSession, err := session.NewSession(i.crypto, token)
		if err != nil {
			return err
		}

		sctx.Session = NewSession

		return next(sctx)
	}
}

func NewSessionMiddleware(secret string, crypto crypto.Crypto, logger logger.Logger) (SessionMiddleware, error) {
	return &impl{secret, crypto, logger}, nil
}
