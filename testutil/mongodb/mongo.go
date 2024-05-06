// Package mongodb provides a shared MongoDB container for testing purposes.
// It includes the ability to create a new database off the shared container.
// This is particularly useful for running isolated tests against a fresh MongoDB instance.
package mongodb

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync/atomic"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoEnvironment is a MongoDB test environment.
// It starts a single MongoDB container and provides a client to it.
type MongoEnvironment struct {
	idCounter atomic.Int64
	client    *mongo.Client
	URI       string // server address including scheme and port
}

// NewMongoEnvironment creates a new MongoDB test environment.
func NewMongoEnvironment() *MongoEnvironment {
	return &MongoEnvironment{}
}

// Start starts the MongoDB container and returns a function to stop it.
// Make sure to call the returned cancel function when done with the container,
// regardless of whether an error occurred.
func (m *MongoEnvironment) Start() (context.CancelFunc, error) {
	const (
		startupTimeout = 3 * time.Minute
		mappedPort     = "27017/tcp"
	)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	mc, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image:        "mongo:latest",
				ExposedPorts: []string{mappedPort},
				Env: map[string]string{
					"MONGO_INITDB_DATABASE": "test",
				},
				WaitingFor: wait.ForLog("Waiting for connections").WithStartupTimeout(startupTimeout),
			},
			Started: true,
		},
	)
	if err != nil {
		return cancel, fmt.Errorf("failed to start MongoDB container: %w", err)
	}

	mongoHost, _ := mc.Host(ctx)
	mongoPort, _ := mc.MappedPort(ctx, mappedPort)

	uri := "mongodb://" + net.JoinHostPort(mongoHost, mongoPort.Port())

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return cancel, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	m.client = client
	m.URI = uri

	return func() {
		defer cancel()

		if dErr := client.Disconnect(ctx); dErr != nil {
			log.Printf("Failed to disconnect from MongoDB: %v", dErr)
		}

		if tErr := mc.Terminate(ctx); tErr != nil {
			log.Printf("Failed to terminate MongoDB container: %v", tErr)
		}
	}, nil
}

// NewInstance creates a new Mongo database.
// Make sure to call cancel function when done with the database.
func (m *MongoEnvironment) NewInstance() (*mongo.Database, context.CancelFunc) {
	const connectionTimeout = 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	db := m.client.Database(fmt.Sprintf("testdb-%d", m.idCounter.Add(1)))

	closer := func() {
		if err := db.Drop(ctx); err != nil {
			log.Fatal("Failed to drop database: ", err)
		}

		cancel()
	}

	return db, closer
}
