package myapp

import (
	"context"
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

func MySimpleWorkflowDefinition(ctx context.Context) error {
	return nil
}

// MyWorkflowDefinition is My custom Workflow Definition.
func MyWorkflowDefinition(ctx workflow.Context, param MyWorkflowParam) (*MyWorkflowResultObject, error) {

	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	}

	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	activityParam := MyActivityParam{
		ActivityParamX: param.WorkflowParamX,
		ActivityParamY: param.WorkflowParamY,
	}

	var a *MyActivityObject
	var activityResult *MyActivityResultObject

	err := workflow.ExecuteActivity(ctx, a.MyActivityDefinition, activityParam).Get(ctx, &activityResult)
	if err!= nil {
        return nil, err
    }

	var infoResult *MyActivityResultObject
	err = workflow.ExecuteActivity(ctx, a.GetInfo).Get(ctx, &infoResult)
	if err!= nil {
        return nil, err
    }

	infoParam := MyActivityParam{
		ActivityParamX: infoResult.ResultFieldX,
        ActivityParamY: infoResult.ResultFieldY,
	}
	err = workflow.ExecuteActivity(ctx, a.PrintInfo, infoParam).Get(ctx, nil)
	if err!= nil {
        return nil, err
    }

	return &MyWorkflowResultObject{
		activityResult.ResultFieldX,
		activityResult.ResultFieldY,
	}, nil
}


