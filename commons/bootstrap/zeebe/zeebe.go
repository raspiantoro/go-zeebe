package zeebe

import (
	"sync"

	"github.com/camunda/zeebe/clients/go/v8/pkg/commands"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
)

type Config struct {
	Gateway  string
	Insecure bool
}

var (
	doOnce sync.Once
	client zbc.Client
)

func InitClient(cfg Config) (err error) {
	doOnce.Do(func() {
		client, err = zbc.NewClient(&zbc.ClientConfig{
			GatewayAddress:         cfg.Gateway,
			UsePlaintextConnection: cfg.Insecure,
		})
		if err != nil {
			return
		}
	})

	return
}

type CommandOption func(c *commandCtx)

type commandCtx struct {
	version   int32
	variables map[string]interface{}
}

func WithVersion(version int32) CommandOption {
	return func(c *commandCtx) {
		c.version = version
	}
}

func WithVariables(variables map[string]interface{}) CommandOption {
	return func(c *commandCtx) {
		c.variables = variables
	}
}

func NewCommand(proccessID string, opts ...CommandOption) (command commands.CreateInstanceCommandStep3, err error) {
	var (
		cmd interface{}
		ctx commandCtx
	)

	cmd = client.NewCreateInstanceCommand().
		BPMNProcessId(proccessID)

	for _, opt := range opts {
		opt(&ctx)
	}

	if ctx.version == 0 {
		cmd = cmd.(commands.CreateInstanceCommandStep2).LatestVersion()
	} else {
		cmd = cmd.(commands.CreateInstanceCommandStep2).Version(ctx.version)
	}

	if len(ctx.variables) > 0 {
		cmd.(commands.CreateInstanceCommandStep3).VariablesFromMap(ctx.variables)
	}

	command = cmd.(commands.CreateInstanceCommandStep3)

	return
}

func NewJobWorker(jobType string, handler worker.JobHandler) (jobWorker worker.JobWorker) {
	jobWorker = client.NewJobWorker().JobType(jobType).Handler(handler).Open()
	return
}

func NewMessageCommand(messageName string) (command commands.PublishMessageCommandStep2) {
	command = client.NewPublishMessageCommand().
		MessageName(messageName)

	return
}
