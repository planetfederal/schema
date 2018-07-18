// Copyright 2018, Boundless, https://boundlessgeo.com - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential

package schema

import "encoding/json"

type aliasCommand Command
type jsonCommand struct {
	Data map[string]interface{} `protobuf:"message,3,opt,name=data,proto3" json:"data,omitempty"`
	*aliasCommand
}

// MarshalJSON returns the JSON encoding of a Command
func (command *Command) MarshalJSON() ([]byte, error) {
	// get the data field as json
	var dataAsJSON map[string]interface{}
	if unmarshalErr := json.Unmarshal(command.Data, &dataAsJSON); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return json.Marshal(&jsonCommand{
		Data:         dataAsJSON,
		aliasCommand: (*aliasCommand)(command),
	})
}

// UnmarshalJSON parses JSON-encoded data into a Command
func (command *Command) UnmarshalJSON(bytes []byte) error {
	// create a jsonCommand object
	aux := &jsonCommand{
		aliasCommand: (*aliasCommand)(command),
	}

	// read the JSON data into a jsonCommand
	if err := json.Unmarshal(bytes, &aux); err != nil {
		return err
	}

	// convert the Data json object to bytes
	data, parseErr := json.Marshal(aux.Data)
	if parseErr != nil {
		return parseErr
	}

	command.Data = data

	// no errors
	return nil
}
