package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.temporal.io/sdk/client"

	"temporal/myapp"
)

func main() {
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort, 
	})

	if err != nil {
		log.Fatalln("Unable to create Temporal Client", err)
	}
	defer temporalClient.Close()

	http.HandleFunc("/startWorkflow", func(w http.ResponseWriter, r *http.Request) {
		startWorkflowHandler(w, r, temporalClient)
	})

	err = http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatalln("Unable to run http server", err)
	}
}

func startWorkflowHandler(w http.ResponseWriter, r *http.Request, temporalClient client.Client) {
	workflowOptions := client.StartWorkflowOptions{
		ID: "my-workflow-id",
		TaskQueue: "my-custom-task-queue-name",
	}

	workflowParams := myapp.MyWorkflowParam{
		WorkflowParamX: "Hello from Workflow",
        WorkflowParamY: 100,
	}

	// Make the call to the temporal cluster to start the workflow execution.
	workflowExecution, err := temporalClient.ExecuteWorkflow(
		context.Background(),
		workflowOptions,
        myapp.MyWorkflowDefinition,
        workflowParams,
	)
	if err!= nil {
        log.Fatalln("Unable to execute the workflow", err)
    }
	log.Println("Started Workflow!")
	log.Println("WorkflowID:", workflowExecution.GetID())
	log.Println("RunID:", workflowExecution.GetRunID())
	var result myapp.MyWorkflowResultObject
	workflowExecution.Get(context.Background(), &result)
	if err != nil {
		log.Println("Unable to get workflow result", err)
	}

	b, err := json.Marshal(result)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println(string(b))
}