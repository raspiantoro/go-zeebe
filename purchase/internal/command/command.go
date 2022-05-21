package command

import "github.com/camunda/zeebe/clients/go/v8/pkg/commands"

type Command struct {
	Purchase commands.CreateInstanceCommandStep3
	Approval commands.PublishMessageCommandStep2
}
