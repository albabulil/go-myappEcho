package dbmongo

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/albabulil/go-myappEcho/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConn struct {
	IsConnected bool
	Done        bool
}

func InitDatabase(ctx context.Context, dbURI string) *mongo.Client {
	client, err := NewMongoClient(ctx, dbURI)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err.Error())
	}

	return client
}

func NewMongoClient(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	maxRetryAttempt := 3
	if maxAttempt, err := strconv.Atoi(config.DB_TIMEOUT); err == nil {
		maxRetryAttempt = maxAttempt
	}

	message := make(chan string, maxRetryAttempt)
	status := make(chan *DBConn, maxRetryAttempt)

	connected := false

	go checkConnection(ctx, client, status, message, maxRetryAttempt, 0)

	for {
		done := false
		select {
		case m := <-message:
			log.Println(m)
		case s := <-status:
			done = s.Done
			connected = s.IsConnected
		}
		if done {
			close(message)
			close(status)
			break
		}
	}

	if !connected {
		log.Fatal("Failed to connect")
	}

	return client, nil
}

func checkConnection(ctx context.Context, client *mongo.Client, status chan<- *DBConn, message chan<- string, maxRetry int, retryAttempt int) {

	message <- "Verifiying connection"

	err := client.Ping(ctx, nil)
	if err == nil {
		message <- "Database connected!"
		status <- &DBConn{
			Done:        true,
			IsConnected: true,
		}
	}

	if err != nil && retryAttempt > maxRetry {
		message <- "Retry attempt exceeded!"
		status <- &DBConn{
			IsConnected: false,
			Done:        true,
		}
	}

	if err != nil && retryAttempt <= maxRetry {
		message <- "Connection failed"
		time.Sleep(time.Millisecond * 3000)
		message <- "Retrying..."
		client.Connect(ctx)
		checkConnection(ctx, client, status, message, maxRetry, retryAttempt+1)
	}
}
