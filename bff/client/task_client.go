package client

import (
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/moonorange/gomicroservice/protogo/gen/genconnect"
	"github.com/sirupsen/logrus"
)

var (
	queryClient   genconnect.TaskServiceClient
	commandClient genconnect.TaskServiceClient
)

func NewQueryServiceClient() genconnect.TaskServiceClient {
	queryHost := os.Getenv("QUERY_SERVICE_HOST")
	if queryHost == "" {
		logrus.Fatal("empty QUERY_SERVICE_HOST")
	}
	// Set up a connection to the server.
	// Create a gRPC client using the connect.WithGRPC() option
	if queryClient != nil {
		return queryClient
	}
	queryClient = genconnect.NewTaskServiceClient(
		http.DefaultClient,
		queryHost,
		connect.WithGRPC(),
	)

	return queryClient
}

func NewCommandServiceClient() genconnect.TaskServiceClient {
	commandHost := os.Getenv("COMMAND_SERVICE_HOST")
	if commandHost == "" {
		logrus.Fatal("empty COMMAND_SERVICE_HOST")
	}
	if commandClient != nil {
		return commandClient
	}
	// Set up a connection to the server.
	// Create a gRPC client using the connect.WithGRPC() option
	commandClient = genconnect.NewTaskServiceClient(
		http.DefaultClient,
		commandHost,
		connect.WithGRPC(),
	)

	return commandClient
}
