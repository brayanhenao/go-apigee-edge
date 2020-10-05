package apigee

import (
	"path"
)

const targetServersPath = "targetservers"

// TargetServersService is an interface for interfacing with the Apigee Edge Admin API
// dealing with target servers.
type TargetServersService interface {
	Create(TargetServer, string) (*TargetServer, *Response, error)
	Delete(string, string) (*Response, error)
	Get(string, string) (*TargetServer, *Response, error)
	Update(TargetServer, string) (*TargetServer, *Response, error)
}

type TargetServersServiceOp struct {
	client *ApigeeClient
}

var _ TargetServersService = &TargetServersServiceOp{}

type TargetServer struct {
	Enabled bool     `json:"isEnabled"`
	Host    string   `json:"host,omitempty"`
	Name    string   `json:"name,omitempty"`
	Port    int      `json:"port,omitempty"`
	SSLInfo *SSLInfo `json:"sSLInfo,omitempty"`
}

// For some reason Apigee returns SOME bools as strings and others a bools.
type SSLInfo struct {
	Ciphers                []string `json:"ciphers,omitempty"`
	ClientAuthEnabled      string   `json:"clientAuthEnabled,omitempty"`
	IgnoreValidationErrors bool     `json:"ignoreValidationErrors"`
	KeyAlias               string   `json:"keyAlias,omitempty"`
	KeyStore               string   `json:"keyStore,omitempty"`
	Protocols              []string `json:"protocols,omitempty"`
	SSLEnabled             string   `json:"enabled,omitempty"`
	TrustStore             string   `json:"trustStore,omitempty"`
}

func (s *TargetServersServiceOp) Get(name string, env string) (*TargetServer, *Response, error) {

	path := path.Join(environmentsPath, env, targetServersPath, name)

	req, e := s.client.NewRequest("GET", path, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedTargetServer := TargetServer{}
	resp, e := s.client.Do(req, &returnedTargetServer)
	if e != nil {
		return nil, resp, e
	}
	return &returnedTargetServer, resp, e

}

func (s *TargetServersServiceOp) Create(targetServer TargetServer, env string) (*TargetServer, *Response, error) {

	return postOrPutTargetServer(targetServer, env, "POST", s)

}

func (s *TargetServersServiceOp) Update(targetServer TargetServer, env string) (*TargetServer, *Response, error) {

	return postOrPutTargetServer(targetServer, env, "PUT", s)

}

func (s *TargetServersServiceOp) Delete(name string, env string) (*Response, error) {

	path := path.Join(environmentsPath, env, targetServersPath, name)

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

func postOrPutTargetServer(targetServer TargetServer, env string, opType string, s *TargetServersServiceOp) (*TargetServer, *Response, error) {

	uripath := ""

	if opType == "PUT" {
		uripath = path.Join(environmentsPath, env, targetServersPath, targetServer.Name)
	} else {
		uripath = path.Join(environmentsPath, env, targetServersPath)
	}

	req, e := s.client.NewRequest(opType, uripath, targetServer)
	if e != nil {
		return nil, nil, e
	}

	returnedTargetServer := TargetServer{}

	resp, e := s.client.Do(req, &returnedTargetServer)
	if e != nil {
		return nil, resp, e
	}

	return &returnedTargetServer, resp, e

}
