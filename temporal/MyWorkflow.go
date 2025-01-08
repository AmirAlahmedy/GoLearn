package temporal_workflow

import (
	"time"


	"go.temporal.io/sdk/workflow"
)

// MyWorkflowParam is the object passed to the Workflow.
type MyWorkflowParam struct {
    WorkflowParamX string
    WorkflowParamY int
}

type MyWorkflowResultObject struct {
	ResultX string
	ResultY int
}

// MyWorkflowDefinition is My custom Workflow Definition.
func MyWorkflowDefinition(ctx workflow.Context, param MyWorkflowParam) (*MyWorkflowResultObject, error) {

	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	}

	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	actvityParam := MyActivityParam(
		ActivityParamX: param.WorkflowParamX,
		ActivityParamY: param.WorkflowParamY, 
	)


	return &MyWorkflowResultObject{
		"",
		0,
	}, nil
}
