package command

import "github.com/camunda/zeebe/clients/go/v8/pkg/commands"

type Command struct {
	Approval commands.PublishMessageCommandStep2
}
