package workflow

import (
	"context"
	"log"

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

	purchaseID, err := p.Service.Purchase.Prepare(ctx, variables)
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

	log.Println("Complete job", jobKey, "of type", job.Type, " from process ", processKey)
	log.Println("Processing purchase:", purchaseID)

	_, err = request.Send(ctx)
	if err != nil {
		handler.FailJob(client, job)
		return
	}

	log.Println("Successfully completed job")
}
