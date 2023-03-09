package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"go.mongodb.org/mongo-driver/bson"
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

// JSONMarshal ...
func JSONMarshal(instance interface{}) (bytes []byte, err error) {
	bytes, err = json.Marshal(instance)
	if err != nil {
		return nil, err
	}

	// Returns
	return bytes, nil
}

// BSONMarshal ...
func BSONMarshal(instance interface{}) (bytes []byte, err error) {
	bytes, err = bson.Marshal(instance)
	if err != nil {
		return nil, err
	}

	// Returns
	return bytes, nil
}

// BSONUnmarshal ...
func BSONUnmarshal(bytes []byte, instance interface{}) (err error) {
	err = bson.Unmarshal(bytes, instance)
	if err != nil {
		return err
	}

	// Returns
	return nil
}

// GetEnv ...
func GetEnv(key string) (value string) {
	return os.Getenv(key)
}

// Contains ...
func Contains(keys []string, key string) (found bool) {
	for _, v := range keys {
		if v == key {
			return true
		}
	}
	return false
}
