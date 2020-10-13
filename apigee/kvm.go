package apigee

import "path"

const kvmPath = "keyvaluemaps"

// KeyValueMapsService is an interface for interfacing with the Apigee Edge Admin API
// dealing with KeyValueMap.
type KeyValueMapsService interface {
	Create(string, KeyValueMap) (*KeyValueMap, *Response, error)
	Delete(string, string) (*Response, error)
	Get(string, string) (*KeyValueMap, *Response, error)
}

// KeyValueMapsServiceOp holds creds
type KeyValueMapsServiceOp struct {
	client *ApigeeClient
}

var _ KeyValueMapsService = &KeyValueMapsServiceOp{}

// EntryStruct Holds the Key value map entry
type EntryStruct struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// KeyValueMap Holds the Key value map
type KeyValueMap struct {
	Encrypted bool          `json:"encrypted,omitempty"`
	Entry     []EntryStruct `json:"entry,omitempty"`
	Name      string        `json:"name,omitempty"`
}

// Get the Keyvaluemap
func (s *KeyValueMapsServiceOp) Get(env string, name string) (*KeyValueMap, *Response, error) {

	path := path.Join(environmentsPath, env, kvmPath, name)

	req, e := s.client.NewRequest("GET", path, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedKeyValueMap := KeyValueMap{}
	resp, e := s.client.Do(req, &returnedKeyValueMap)
	if e != nil {
		return nil, resp, e
	}
	return &returnedKeyValueMap, resp, e

}

// Create a new key value map
func (s *KeyValueMapsServiceOp) Create(env string, keyValueMap KeyValueMap) (*KeyValueMap, *Response, error) {

	return postOrPutKeyValueMap(env, keyValueMap, "POST", s)
}

// Delete an existing key value map
func (s *KeyValueMapsServiceOp) Delete(env string, name string) (*Response, error) {

	path := path.Join(environmentsPath, env, kvmPath, name)

	req, e := s.client.NewRequest("DELETE", path, nil)
	if e != nil {
		return nil, e
	}

	resp, e := s.client.Do(req, nil)
	if e != nil {
		return resp, e
	}

	return resp, e

}

func postOrPutKeyValueMap(env string, keyValueMap KeyValueMap, opType string, s *KeyValueMapsServiceOp) (*KeyValueMap, *Response, error) {

	uripath := ""

	if opType == "PUT" {
		uripath = path.Join(environmentsPath, env, kvmPath, keyValueMap.Name)
	} else {
		uripath = path.Join(environmentsPath, env, kvmPath)
	}

	req, e := s.client.NewRequest(opType, uripath, keyValueMap)
	if e != nil {
		return nil, nil, e
	}

	returnedKeyValueMap := KeyValueMap{}

	resp, e := s.client.Do(req, &returnedKeyValueMap)
	if e != nil {
		return nil, resp, e
	}

	return &returnedKeyValueMap, resp, e

}
