package main

import (
	"context"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"

	"github.com/moonorange/gomicroservice/protogo/gen"
	"github.com/moonorange/gomicroservice/protogo/gen/genconnect"
)

const (
	defaultPort = "8082"
	defaultHost = "localhost"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	host := os.Getenv("COMMAND_SERVICE_HOST")
	if host == "" {
		host = defaultHost
	}

	mux := http.NewServeMux()
	path, handler := genconnect.NewTaskServiceHandler(&taskServer{})
	mux.Handle(path, handler)
	logrus.Println("... Listening on", host+":"+port)

	eg := errgroup.Group{}
	// Start the gRPC server
	eg.Go(func() error { return http.ListenAndServe(":"+port, h2c.NewHandler(mux, &http2.Server{})) })
	logrus.Printf("Command service is running on host %s", host+":"+port)

	err := eg.Wait()
	if err != nil {
		logrus.Fatal("failed to serve: ", err)
	}
}

// taskServer implements the TaskService API.
type taskServer struct {
	genconnect.UnimplementedTaskServiceHandler
}

// Just return a task for simplicity
func (t *taskServer) CreateTask(ctx context.Context, req *connect.Request[gen.CreateTaskRequest]) (*connect.Response[gen.CreateTaskResponse], error) {
	task := &gen.Task{
		Id:   1,
		Text: req.Msg.Text,
		Tags: req.Msg.Tags,
	}
	return connect.NewResponse(&gen.CreateTaskResponse{Task: task}), nil
}
