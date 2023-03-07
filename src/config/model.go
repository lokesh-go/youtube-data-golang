package config

// Config ...
type Config struct {
	App    string
	Server struct {
		HTTP struct {
			Port string
		}
	}
	Datastores struct {
		Youtube struct {
			Hosts       []string
			User        string
			Password    string
			Auth        bool
			Authsource  string
			Database    string
			Collections struct {
				Youtube string
			}
			Connections struct {
				ReplicaSetName         *string
				MinPoolSize            *uint64
				MaxPoolSize            *uint64
				MaxConnecting          *uint64
				MaxConnIdleTime        *int
				Timeout                *int
				SocketTimeout          *int
				ReadConcernMajority    bool
				ReadSecondaryPreferred bool
			}
		}
	}
	Youtube struct {
		Auth struct {
			ClientSecretFilePath string
			TokenFilePath        string
		}
		Search struct {
			Part       string
			Order      string
			Type       string
			MaxResults int64
			Pagination struct {
				Enabled bool
			}
		}
	}
	Job struct {
		Enabled  bool
		Interval string
	}
}
