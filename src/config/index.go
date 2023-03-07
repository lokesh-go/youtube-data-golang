package config

import utils "github.com/lokesh-go/youtube-data-golang/src/utils"

// Initialize ...
func Initialize() (configModel *Config, err error) {
	// Gets the file path
	path := "config/config.json"

	// Reads json file
	err = utils.ReadJSONFile(path, &configModel)
	if err != nil {
		return nil, err
	}

	// Returns
	return configModel, nil
}
