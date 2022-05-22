package handler

import (
	"context"
	"log"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/raspiantoro/go-zeebe/approval/internal/commons"
	"github.com/raspiantoro/go-zeebe/approval/internal/service"
)

type Option struct {
	commons.Option
	Service service.Service
}

func FailJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())
	ctx := context.Background()
	_, err := client.NewFailJobCommand().
		JobKey(job.GetKey()).Retries(job.Retries - 1).
		Send(ctx)
	if err != nil {
		panic(err)
	}
}
