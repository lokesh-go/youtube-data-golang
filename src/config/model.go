package config

// Config ...
type Config struct {
	App     string
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
}
