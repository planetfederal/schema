// Copyright 2018, Boundless, https://boundlessgeo.com - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential

package schema

import "encoding/json"

// MarshalJSON returns the JSON encoding of a Command
func (command *Command) MarshalJSON() ([]byte, error) {
	// objMap := make(map[string]interface{})
	// objMap["id"] = command.Id
	// objMap["action"] = command.Action

	// var dataAsJSON interface{}
	// if unmarshalErr := json.Unmarshal(command.Data, &dataAsJSON); unmarshalErr != nil {
	// 	return nil, unmarshalErr
	// }
	// objMap["data"] = dataAsJSON

	// objMap["context"] = command.Context
	// objMap["metadata"] = command.Metadata
	// objMap["created"] = command.Created
	// return json.Marshal(objMap)

	var dataAsJSON map[string]interface{}
	if unmarshalErr := json.Unmarshal(command.Data, &dataAsJSON); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	type Alias Command
	return json.Marshal(&struct {
		Data map[string]interface{} `protobuf:"message,3,opt,name=data,proto3" json:"data,omitempty"`
		*Alias
	}{
		Data:  dataAsJSON,
		Alias: (*Alias)(command),
	})
}

// UnmarshalJSON parses JSON-encoded data into a Command
func (command *Command) UnmarshalJSON(bytes []byte) error {
	// deserialize into a map
	// var objMap map[string]*json.RawMessage
	// err := json.Unmarshal(bytes, &objMap)
	// if err != nil {
	// 	return err
	// }

	// // get the command's id
	// if err = json.Unmarshal(*objMap["id"], &command.Id); err != nil {
	// 	return err
	// }

	// // get the command's action
	// if err = json.Unmarshal(*objMap["action"], &command.Action); err != nil {
	// 	return err
	// }

	// // get the command's data
	// if err = json.Unmarshal(*objMap["data"], &command.Data); err != nil {
	// 	return err
	// }

	// // get the command's context
	// if err = json.Unmarshal(*objMap["context"], &command.Context); err != nil {
	// 	return err
	// }

	// // get the command's metadata
	// if err = json.Unmarshal(*objMap["metadata"], &command.Metadata); err != nil {
	// 	return err
	// }

	// // get the command's created
	// if err = json.Unmarshal(*objMap["created"], &command.Created); err != nil {
	// 	return err
	// }

	// // no errors
	// return nil

	type Alias Command
	aux := &struct {
		Data map[string]interface{} `protobuf:"message,3,opt,name=data,proto3" json:"data,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(command),
	}

	if err := json.Unmarshal(bytes, &aux); err != nil {
		return err
	}

	dataJSON, parseErr := json.Marshal(aux.Data)
	if parseErr != nil {
		return parseErr
	}
	command.Data = dataJSON

	return nil
}
