package client

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/protogo/gen/genconnect"
)

var (
	taskClient genconnect.TaskServiceClient
)

func NewTaskServiceClient(baseURL string) genconnect.TaskServiceClient {
	// Set up a connection to the server.
	// Create a gRPC client using the connect.WithGRPC() option
	if taskClient != nil {
		return taskClient
	}
	taskClient = genconnect.NewTaskServiceClient(
		http.DefaultClient,
		baseURL,
		connect.WithGRPC(),
	)

	return taskClient
}
