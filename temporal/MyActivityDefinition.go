package temporal_workflow

import (	
	"context"

	"go.temporal.io/sdk/activity"
)

type MyActivityParam struct {
	ActivityParamX string
	ActivityParamY int
}

type MyAcrtivityResultOject struct {
	ResultFieldX string
	ResultFieldY int
}

func MySimpleActivityDefinition(ctx context.Context) error {
	return nil
}

// Maintains shared state across Activities
type MyActivityObject struct {
	Message *string
	Number *int
}

func (a *MyActivityObject) MyActivityDefinition(ctx context.Context, param *MyActivityParam) (*MyAcrtivityResultOject, error) {
	// Use Activities for calling external APIs
	logger := activity.GetLogger(ctx)
	logger.Info("The message is:", param.ActivityParamX)
	logger.Info("The number is:", param.ActivityParamY)

	result := &MyAcrtivityResultOject{
		ResultFieldX: "This is the result field X",
		ResultFieldY: 100,
	}

	// Return the result back to the Workflow Execution.
	// The resuts persist with the Event History of the Workflow Execution.
	return result, nil
}

func (a *MyActivityObject) PrintInfo(ctx context.Context, param MyActivityParam) error {
	logger := activity.GetLogger(ctx)
	logger.Info("The message is: ", param.ActivityParamX)
	logger.Info("The number is: ", param.ActivityParamY)
	return nil
}

func (a *MyActivityObject) GetInfo(ctx context.Context, param MyActivityParam) (*MyAcrtivityResultOject, error) {
	return &MyActivityObject{
		ResultFieldX: *a.Message,
		ResultFieldY: *a.Number,
	}, nil
}

