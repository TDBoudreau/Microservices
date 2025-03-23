package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// Register the RPC Server
	err = rpc.Register(new(RPCServer))
	go app.rpcListen()

	// listen for gRPC connections
	go app.grpcListen()

	// start web server
	log.Printf("Starting logger-service on port %s...\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func (app *Config) rpcListen() error {
	log.Println("Starting RPC server on port", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(rpcConn)
	}
}

func connectToMongo() (*mongo.Client, error) {
	usernamePath := os.Getenv("MONGO_USERNAME")
	passwordPath := os.Getenv("MONGO_PASSWORD")

	username, err := getSecret(usernamePath)
	if err != nil {
		log.Printf("error reading username secret from %s: %v", usernamePath, err)
		return nil, err
	}

	password, err := getSecret(passwordPath)
	if err != nil {
		log.Printf("error reading password secret from %s: %v", passwordPath, err)
		return nil, err
	}

	// create the connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting:", err)
		return nil, err
	}

	err = c.Ping(context.Background(), nil)
	if err != nil {
		log.Println("error pinging:", err)
		return nil, err
	}

	return c, nil
}

func getSecret(secretPath string) (string, error) {
	secret, err := os.ReadFile(secretPath)
	if err != nil {
		return "", err
	}
	// Trim whitespace/newlines
	return strings.TrimSpace(string(secret)), nil
}
