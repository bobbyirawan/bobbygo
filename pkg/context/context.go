package context

import (
	"bobbygo/pkg/errors"
	"bobbygo/pkg/session"

	"github.com/labstack/echo/v4"
)

type (
	BobbyGoContext struct {
		echo.Context
		Session *session.Session
	}

	Success struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta,omitempty"`
	}

	Failed struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}
)

func (sc *BobbyGoContext) Success(data interface{}) error {
	return sc.JSON(200, Success{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func (sc *BobbyGoContext) SuccessWithMeta(data, meta interface{}) error {
	return sc.JSON(200, Success{
		Code:    200,
		Message: "success",
		Data:    data,
		Meta:    meta,
	})
}

func (sc *BobbyGoContext) Fail(err error) error {
	var (
		ed = errors.ExtractError(err)
	)

	return sc.JSON(ed.HttpCode, Failed{
		Code:    ed.Code,
		Message: "failed",
		Error:   ed.Message,
	})
}

func NewEmptyBobbyGoContext(parent echo.Context) *BobbyGoContext {
	return &BobbyGoContext{parent, nil}
}

func NewBobbyGoContext(parent echo.Context) (*BobbyGoContext, error) {
	pctx, ok := parent.(*BobbyGoContext)
	if !ok {
		return nil, errors.ErrSession
	}
	if pctx.Session == nil {
		return nil, errors.ErrSession
	}
	return pctx, nil
}
