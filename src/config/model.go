package config

// Config ...
type Config struct {
	Youtube struct {
		Auth struct {
			ClientSecretFilePath string
			TokenFilePath        string
		}
	}
}
