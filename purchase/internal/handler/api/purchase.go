package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/raspiantoro/go-zeebe/purchase/internal/handler"
)

type PurchaseHandler struct {
	handler.Option
}

type SubmitPurchaseRequest struct {
	Item  string `json:"item"`
	Price uint64 `json:"price"`
}

func NewPurchaseHandler(opt handler.Option) *PurchaseHandler {
	return &PurchaseHandler{opt}
}

func (p *PurchaseHandler) Submit(c echo.Context) (err error) {
	req := new(SubmitPurchaseRequest)

	err = c.Bind(req)
	if err != nil {
		return
	}

	variables := map[string]interface{}{
		"item":  req.Item,
		"price": req.Price,
	}

	cmd, err := p.Command.Purchase.VariablesFromMap(variables)
	if err != nil {
		log.Error(err)
		return
	}

	result, err := cmd.Send(c.Request().Context())
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println(result.String())

	return
}
