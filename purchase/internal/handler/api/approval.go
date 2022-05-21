package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/raspiantoro/go-zeebe/purchase/internal/handler"
)

type ApprovalHandler struct {
	handler.Option
}

type ApprovalRequest struct {
	PurchaseID string `json:"purchase_id"`
	Action     string `json:"action"`
}

func NewApprovalHandler(opt handler.Option) *ApprovalHandler {
	return &ApprovalHandler{opt}
}

func (a *ApprovalHandler) Submit(c echo.Context) (err error) {
	req := new(ApprovalRequest)

	err = c.Bind(req)
	if err != nil {
		return
	}

	variables := map[string]interface{}{
		"approval_action": req.Action,
	}

	cmd, err := a.Command.Approval.CorrelationKey(req.PurchaseID).VariablesFromMap(variables)
	if err != nil {
		log.Error(err)
		return
	}

	cmd = cmd.MessageId(req.PurchaseID)

	result, err := cmd.Send(c.Request().Context())
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println(result.String())

	return
}
