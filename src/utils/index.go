package utils

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSONFile ...
func ReadJSONFile(path string, instance interface{}) (err error) {
	// Forms bytes
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Unmarshals
	err = JSONUnmarshal(bytes, instance)
	if err != nil {
		return err
	}

	// Returns
	return nil
}

// JSONUnmarshal ...
func JSONUnmarshal(bytes []byte, instance interface{}) (err error) {
	err = json.Unmarshal(bytes, instance)
	if err != nil {
		return err
	}

	// Returns
	return nil
}
