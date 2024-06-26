package main

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"connectrpc.com/connect"
	"github.com/moonorange/gomicroservice/protogo/gen"
	"github.com/moonorange/gomicroservice/protogo/gen/genconnect"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
)

const (
	defaultPort = "8081"
	defaultHost = "localhost"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	host := os.Getenv("QUERY_SERVICE_HOST")
	if host == "" {
		host = defaultHost
	}

	mux := http.NewServeMux()
	path, handler := genconnect.NewTaskServiceHandler(&taskServer{})
	mux.Handle(path, handler)
	logrus.Println("... Listening on", host)

	eg := errgroup.Group{}
	// Start the gRPC server
	eg.Go(func() error { return http.ListenAndServe(":"+port, h2c.NewHandler(mux, &http2.Server{})) })
	logrus.Printf("Query service is running on host %s", host)

	err := eg.Wait()
	if err != nil {
		logrus.Fatal("failed to serve: ", err)
	}
}

type taskServer struct {
	genconnect.UnimplementedTaskServiceHandler
}

// Just return a task for simplicity
func (t *taskServer) GetTask(ctx context.Context, req *connect.Request[gen.GetTaskRequest]) (*connect.Response[gen.GetTaskResponse], error) {
	id, _ := strconv.ParseInt(req.Msg.TaskId, 10, 32)
	task := gen.Task{
		Id:   int32(id),
		Text: "This is a task",
		Tags: []string{"tag1", "tag2"},
	}
	return connect.NewResponse(&gen.GetTaskResponse{Task: &task}), nil
}

// Just return a list of tasks for simplicity
func (t *taskServer) ListTasksByTag(ctx context.Context, req *connect.Request[gen.ListTasksByTagRequest]) (*connect.Response[gen.ListTasksByTagResponse], error) {
	tasks := []*gen.Task{
		{
			Id:   1,
			Text: "This is a task",
			Tags: []string{req.Msg.TagName},
		},
		{
			Id:   2,
			Text: "This is a task",
			Tags: []string{req.Msg.TagName},
		},
	}
	return connect.NewResponse(&gen.ListTasksByTagResponse{Tasks: tasks}), nil
}
