package apigee

import "path"

const entriesPath = "entries"
const kvmEntryPath = "keys"

// KeyValueMapEntriesService is an interface for interfacing with the Apigee Edge Admin API
// dealing with KeyValueMapEntry.
type KeyValueMapEntriesService interface {
	Create(string, string, KeyValueMapEntryKeys) (*KeyValueMapEntry, *Response, error)
	Delete(string, string, string) (*Response, error)
	Get(string, string, string) (*KeyValueMapEntryKeys, *Response, error)
	List(string, string) ([]string, *Response, error)
	Update(string, string, KeyValueMapEntryKeys) (*KeyValueMapEntry, *Response, error)
}

// KeyValueMapEntriesServiceOp holds creds
type KeyValueMapEntriesServiceOp struct {
	client *ApigeeClient
}

var _ KeyValueMapEntriesService = &KeyValueMapEntriesServiceOp{}

// KeyValueMapEntryKeys to update
type KeyValueMapEntryKeys struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// KeyValueMapEntry Holds the Key value map
type KeyValueMapEntry struct {
	Entry   []KeyValueMapEntryKeys `json:"entry,omitempty"`
	KVMName string                 `json:"kvmName,omitempty"`
}

// Get the key value map entry
func (s *KeyValueMapEntriesServiceOp) Get(env string, keyValueMapName string, keyValueMapEntry string) (*KeyValueMapEntryKeys, *Response, error) {

	path := path.Join(environmentsPath, env, kvmPath, keyValueMapName, entriesPath, keyValueMapEntry)

	req, e := s.client.NewRequest("GET", path, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedKeyValueMapEntry := KeyValueMapEntryKeys{}
	resp, e := s.client.Do(req, &returnedKeyValueMapEntry)
	if e != nil {
		return nil, resp, e
	}
	return &returnedKeyValueMapEntry, resp, e
}

// Create a new key value map entry
func (s *KeyValueMapEntriesServiceOp) Create(env string, keyValueMapName string, keyValueMapEntry KeyValueMapEntryKeys) (*KeyValueMapEntry, *Response, error) {

	return postOrPutKeyValueMapEntry(keyValueMapName, keyValueMapEntry, env, "POST", s)
}

// Update an existing key value map entry
func (s *KeyValueMapEntriesServiceOp) Update(env string, keyValueMapName string, keyValueMapEntry KeyValueMapEntryKeys) (*KeyValueMapEntry, *Response, error) {

	return postOrPutKeyValueMapEntry(keyValueMapName, keyValueMapEntry, env, "PUT", s)

}

// Delete an existing key value map entry
func (s *KeyValueMapEntriesServiceOp) Delete(env string, keyValueMapName string, keyValueMapEntry string) (*Response, error) {

	path := path.Join(environmentsPath, env, kvmPath, keyValueMapName, entriesPath, keyValueMapEntry)

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

func (s *KeyValueMapEntriesServiceOp) List(env string, keyValueMapName string) ([]string, *Response, error) {
	path := path.Join(environmentsPath, env, kvmPath, keyValueMapName, kvmEntryPath)

	req, e := s.client.NewRequest("GET", path, nil)
	if e != nil {
		return nil, nil, e
	}
	nameList := make([]string, 0)
	resp, e := s.client.Do(req, &nameList)
	if e != nil {
		return nil, resp, e
	}
	return nameList, resp, e
}

func postOrPutKeyValueMapEntry(keyValueMapName string, keyValueMapEntry KeyValueMapEntryKeys, env string, opType string, s *KeyValueMapEntriesServiceOp) (*KeyValueMapEntry, *Response, error) {

	uripath := ""

	if opType == "PUT" {
		uripath = path.Join(environmentsPath, env, kvmPath, keyValueMapName, entriesPath, keyValueMapEntry.Name)
	} else {
		uripath = path.Join(environmentsPath, env, kvmPath, keyValueMapName, entriesPath)
	}

	req, e := s.client.NewRequest(opType, uripath, keyValueMapEntry)
	if e != nil {
		return nil, nil, e
	}

	returnedKeyValueMapEntry := KeyValueMapEntry{}

	resp, e := s.client.Do(req, &returnedKeyValueMapEntry)
	if e != nil {
		return nil, resp, e
	}

	return &returnedKeyValueMapEntry, resp, e

}
