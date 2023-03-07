package mongodb

// Config contains all properties required for creating a connection
type Config struct {
	Hosts       []string    // Database server hosts
	AuthEnabled bool        // Enables auth to required user & password to establish connection
	User        string      // Db Username for authentication
	Password    string      // Db password for authentication
	AuthSource  string      // The name of database to use for authentication
	TLSEnabled  bool        // TLS to encrypt all of mongodb's network traffic
	Database    string      // Db name
	Connection  *Connection // More client options
}

type Connection struct {
	ReplicaSetName          *string
	MinPoolSize             *uint64
	MaxPoolSize             *uint64
	MaxConnecting           *uint64
	MaxConnIdleTime         *int
	SocketTimeout           *int
	Timeout                 *int
	ReadConcernWithMajority bool
	ReadSecondaryPreferred  bool
}
