package utils

import "encoding/json"

func Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}

// Deserialize unmarshals the input data into the output interface
func Deserialize(data []byte, output any) error {
	return json.Unmarshal(data, output)
}