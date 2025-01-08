package main

import (
	"log"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"temporal/myapp"
)

func main() {
	temporalClient, err := client.Dial(client.Options{})
	if err!= nil {
        log.Fatalln("Unable to create temporal client", err)
    }
	defer temporalClient.Close()

	myworker := worker.New(temporalClient, "my-custom-task-queue-name", worker.Options{})

	myworker.RegisterWorkflow(myapp.MyWorkflowDefinition)
	registerWFOptions := workflow.RegisterOptions{
		Name: "JustAnotherWorkflow",
	}

	myworker.RegisterWorkflowWithOptions(myapp.MyWorkflowDefinition, registerWFOptions)
	message := "This could be a connection string or endpoint details"
	number := 100
	
	activities := &myapp.MyActivityObject{
		Message: &message,
		Number: &number,
	}

	myworker.RegisterActivity(activities)

	registerAOptions := activity.RegisterOptions{
		Name: "JustAnotherActivity",
	}

	myworker.RegisterActivityWithOptions(myapp.MySimpleActivityDefinition, registerAOptions)
	err = myworker.Run(worker.InterruptCh())
	if err!= nil {
        log.Fatalln("Unable to start worker", err)
    }
}