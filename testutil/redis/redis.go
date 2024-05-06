package redis

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Instance represents a Redis instance.
type Instance struct {
	Client  *redis.Client
	Address string // host:port
	Host    string // host only
	Port    string // port only
}

// NewInstance creates a new Redis instance.
func NewInstance() (*Instance, func(), error) {
	const (
		startupTimeout = 3 * time.Minute
		mappedPort     = "6379/tcp"
	)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	rc, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image:        "redis:latest",
				ExposedPorts: []string{mappedPort},
				WaitingFor:   wait.ForLog("Ready to accept connections").WithStartupTimeout(startupTimeout),
			},
			Started: true,
		},
	)
	if err != nil {
		cancel()

		return nil, cancel, fmt.Errorf("failed to start Redis container: %w", err)
	}

	redisHost, _ := rc.Host(ctx)
	redisPort, _ := rc.MappedPort(ctx, mappedPort)

	address := net.JoinHostPort(redisHost, redisPort.Port())

	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		cancel()

		return nil, cancel, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &Instance{
		Client:  client,
		Address: address,
		Host:    redisHost,
		Port:    redisPort.Port(),
	}, cancel, nil
}
