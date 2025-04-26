package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"
	"time"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	log.Printf("Received log request: %s", input.Name) // Log entry

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	startTime := time.Now() // Start timer
	err := l.Models.LogEntry.Insert(ctx, logEntry)
	duration := time.Since(startTime) // Calculate duration

	if err != nil {
		log.Printf("Failed to insert log entry %s. Duration: %v. Error: %v", input.Name, duration, err) // Log failure + duration
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	log.Printf("Successfully inserted log entry %s. Duration: %v", input.Name, duration) // Log success + duration
	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}

func (app *Config) grpcListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	srv := grpc.NewServer()

	logs.RegisterLogServiceServer(srv, &LogServer{Models: app.Models})

	log.Printf("gRPC Server started on port %s", gRpcPort)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
