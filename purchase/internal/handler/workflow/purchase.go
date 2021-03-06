package workflow

import (
	"context"
	"errors"

	"github.com/labstack/gommon/log"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/raspiantoro/go-zeebe/purchase/internal/handler"
)

type PurchaseHandler struct {
	handler.Option
}

func NewPurchaseHandler(opt handler.Option) *PurchaseHandler {
	return &PurchaseHandler{opt}
}

func (p *PurchaseHandler) Prepare(client worker.JobClient, job entities.Job) {
	ctx := context.Background()

	jobKey := job.GetKey()
	processKey := job.GetProcessInstanceKey()

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		handler.FailJob(client, job)
		return
	}

	variables["process_key"] = processKey

	headers, err := job.GetCustomHeadersAsMap()
	if err != nil {
		handler.FailJob(client, job)
		return
	}

	status, ok := headers["status"]
	if !ok {
		err = errors.New("invalid headers")
		log.Error(err)
		return
	}

	purchaseID, err := p.Service.Purchase.Prepare(ctx, status, variables)
	if err != nil {
		handler.FailJob(client, job)
		return
	}

	variables["purhcase_id"] = purchaseID
	request, err := client.NewCompleteJobCommand().
		JobKey(jobKey).
		VariablesFromMap(variables)
	if err != nil {
		handler.FailJob(client, job)
		return
	}

	log.Print("Preparing purchase complete. job", jobKey, "of type", job.Type, " from process ", processKey, "\n")

	_, err = request.Send(ctx)
	if err != nil {
		handler.FailJob(client, job)
		return
	}
}

func (p *PurchaseHandler) UpdateStatus(client worker.JobClient, job entities.Job) {
	ctx := context.Background()

	jobKey := job.GetKey()
	processKey := job.GetProcessInstanceKey()

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		handler.FailJob(client, job)
		return
	}

	headers, err := job.GetCustomHeadersAsMap()
	if err != nil {
		handler.FailJob(client, job)
		return
	}

	status, ok := headers["status"]
	if !ok {
		err = errors.New("invalid headers")
		log.Error(err)
		return
	}

	err = p.Service.Purchase.UpdateStatus(ctx, status, variables)
	if err != nil {
		handler.FailJob(client, job)
		return
	}

	request, err := client.NewCompleteJobCommand().
		JobKey(jobKey).
		VariablesFromMap(variables)
	if err != nil {
		handler.FailJob(client, job)
		return
	}

	log.Print("updating purchase status complete. job", jobKey, "of type", job.Type, " from process ", processKey, "\n")

	_, err = request.Send(ctx)
	if err != nil {
		handler.FailJob(client, job)
		return
	}
}
