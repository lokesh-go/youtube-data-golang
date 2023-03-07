package mongodb

import (
	"context"
	"crypto/tls"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type clients struct {
	database *mongo.Database
}

// New initializes and returns new mongodb client connection
func New(config *Config) (Methods, error) {
	dbClient, err := config.connect()
	if err != nil {
		return nil, err
	}

	// Returns
	return dbClient, nil
}

func (c *Config) connect() (dbClient *clients, err error) {
	// Assigns hosts
	mongoConnOptions := &options.ClientOptions{
		Hosts: c.Hosts,
	}

	// Checks TLS
	if c.TLSEnabled {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		mongoConnOptions.TLSConfig = tlsConfig
	}

	// Checks auth
	if c.AuthEnabled {
		creds := &options.Credential{
			Username:   c.User,
			Password:   c.Password,
			AuthSource: c.AuthSource,
		}
		mongoConnOptions.Auth = creds
	}

	// Checks connection config
	if c.Connection != nil {
		// Sets replica set name
		if c.Connection.ReplicaSetName != nil {
			mongoConnOptions.SetReplicaSet(*c.Connection.ReplicaSetName)
		}

		// Sets min pool size
		if c.Connection.MinPoolSize != nil {
			mongoConnOptions.SetMinPoolSize(*c.Connection.MinPoolSize)
		}

		// Sets max pool size
		if c.Connection.MaxPoolSize != nil {
			mongoConnOptions.SetMaxPoolSize(*c.Connection.MaxPoolSize)
		}

		// Sets max connecting
		// maximum number of connections a connection pool may establish simultaneously. (default is 2)
		if c.Connection.MaxConnecting != nil {
			mongoConnOptions.SetMaxConnecting(*c.Connection.MaxConnecting)
		}

		// Sets connection timeout
		// Amount of time that a single operation run on this client can execute before returning an error
		if c.Connection.Timeout != nil {
			mongoConnOptions.SetConnectTimeout(time.Duration(*c.Connection.Timeout) * time.Millisecond)
		}

		// Sets socket timeout
		// How long the driver will wait for a socket read or write to return before returning a network error
		if c.Connection.SocketTimeout != nil {
			mongoConnOptions.SetSocketTimeout(time.Duration(*c.Connection.SocketTimeout) * time.Millisecond)
		}

		// Sets max idle time
		// The maximum amount of time that a connection will remain idle in a connection pool before it is removed from the pool and closed
		if c.Connection.MaxConnIdleTime != nil {
			mongoConnOptions.SetMaxConnIdleTime(time.Duration(*c.Connection.MaxConnIdleTime) * time.Millisecond)
		}

		// Sets read concern majority
		if c.Connection.ReadConcernWithMajority {
			majorityLevel := readconcern.Majority().GetLevel()
			rc := readconcern.New(readconcern.Level(majorityLevel))
			mongoConnOptions.SetReadConcern(rc)
		}

		// Sets read secondary preferred
		if c.Connection.ReadSecondaryPreferred {
			rp := readpref.SecondaryPreferred()
			mongoConnOptions.SetReadPreference(rp)
		}
	}

	// New client
	client, err := mongo.NewClient(mongoConnOptions)
	if err != nil {
		return nil, err
	}

	// Assigns timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect mongo client
	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.New("client connection failed ->" + err.Error())
	}

	// Checks client ping
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.New("client ping failed ->" + err.Error())
	}

	// Assigns client
	dbClient = &clients{}
	dbClient.database = client.Database(c.Database)

	// Returns
	return dbClient, nil
}
