package dal

import (
	config "github.com/lokesh-go/youtube-data-golang/src/config"
	mongodb "github.com/lokesh-go/youtube-data-golang/src/pkg/datastore/mongodb"
)

type dal struct {
	config     *config.Config
	dbServices mongodb.Methods
}

// Initialize ...
func Initialize(config *config.Config) (Methods, error) {
	// Connects mongo DB
	mongoServices, err := connectMongoDB(config)
	if err != nil {
		return nil, err
	}

	// Returns
	return &dal{config, mongoServices}, nil
}

func connectMongoDB(config *config.Config) (dbServices mongodb.Methods, err error) {
	// Forms mongo connection config
	connectionConfig := &mongodb.Connection{
		ReplicaSetName:          config.Datastores.Youtube.Connections.ReplicaSetName,
		MinPoolSize:             config.Datastores.Youtube.Connections.MinPoolSize,
		MaxPoolSize:             config.Datastores.Youtube.Connections.MaxPoolSize,
		MaxConnecting:           config.Datastores.Youtube.Connections.MaxConnecting,
		MaxConnIdleTime:         config.Datastores.Youtube.Connections.MaxConnIdleTime,
		Timeout:                 config.Datastores.Youtube.Connections.Timeout,
		SocketTimeout:           config.Datastores.Youtube.Connections.SocketTimeout,
		ReadConcernWithMajority: config.Datastores.Youtube.Connections.ReadConcernMajority,
		ReadSecondaryPreferred:  config.Datastores.Youtube.Connections.ReadSecondaryPreferred,
	}

	// Forms mongo config
	mongoDbConfig := &mongodb.Config{
		Hosts:       config.Datastores.Youtube.Hosts,
		User:        config.Datastores.Youtube.User,
		Password:    config.Datastores.Youtube.Password,
		AuthEnabled: config.Datastores.Youtube.Auth,
		TLSEnabled:  false,
		AuthSource:  config.Datastores.Youtube.Authsource,
		Database:    config.Datastores.Youtube.Database,
		Connection:  connectionConfig,
	}

	// Creates connection
	dbServices, err = mongodb.New(mongoDbConfig)
	if err != nil {
		return nil, err
	}

	// Returns
	return dbServices, nil
}
