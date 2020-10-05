package apigee

import (
	"path"
)

const virtualhostsPath = "virtualhosts"

// VirtualHostService is an interface for interfacing with the Apigee Edge Admin API
// dealing with target servers.
type VirtualHostsService interface {
	Create(VirtualHost, string) (*VirtualHost, *Response, error)
	Delete(string, string) (*Response, error)
	Get(string, string) (*VirtualHost, *Response, error)
	List(string) ([]string, *Response, error)
	Update(VirtualHost, string) (*VirtualHost, *Response, error)
}

type VirtualHostsServiceOp struct {
	client *ApigeeClient
}

var _ VirtualHostsService = &VirtualHostsServiceOp{}

// https://docs.apigee.com/api-platform/fundamentals/virtual-host-property-reference
type VirtualHost struct {
	// Interfaces           []string   `json:"interfaces,omitempty"`
	// PropagateTLSInformation hash    `json:"propagateTLSInformation,omitempty"`
	BaseUrl       string   `json:"baseUrl,omitempty"`
	HostAliases   []string `json:"hostAliases,omitempty"`
	ListenOptions []string `json:"listenOptions,omitempty"`
	Name          string   `json:"name,omitempty"`
	Port          int      `json:"port,omitempty"`
	Properties    []string `json:"properties,omitempty"`
	RetryOptions  []string `json:"retryOptions,omitempty"`
	SSLInfo       []string `json:"sSLInfo,omitempty"`
}

func (s *VirtualHostsServiceOp) List(env string) ([]string, *Response, error) {
	path := path.Join(environmentsPath, env, virtualhostsPath)

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

func (s *VirtualHostsServiceOp) Get(name string, env string) (*VirtualHost, *Response, error) {

	path := path.Join(environmentsPath, env, virtualhostsPath, name)

	req, e := s.client.NewRequest("GET", path, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedVirtualHost := VirtualHost{}
	resp, e := s.client.Do(req, &returnedVirtualHost)
	if e != nil {
		return nil, resp, e
	}
	return &returnedVirtualHost, resp, e

}

func (s *VirtualHostsServiceOp) Create(VirtualHost VirtualHost, env string) (*VirtualHost, *Response, error) {

	return postOrPutVirtualHost(VirtualHost, env, "POST", s)

}

func (s *VirtualHostsServiceOp) Update(VirtualHost VirtualHost, env string) (*VirtualHost, *Response, error) {

	return postOrPutVirtualHost(VirtualHost, env, "PUT", s)

}

func (s *VirtualHostsServiceOp) Delete(name string, env string) (*Response, error) {

	path := path.Join(environmentsPath, env, virtualhostsPath, name)

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

func postOrPutVirtualHost(virtualHost VirtualHost, env string, opType string, s *VirtualHostsServiceOp) (*VirtualHost, *Response, error) {

	uriPath := ""

	if opType == "PUT" {
		uriPath = path.Join(environmentsPath, env, virtualhostsPath, virtualHost.Name)
	} else {
		uriPath = path.Join(environmentsPath, env, virtualhostsPath)
	}

	req, e := s.client.NewRequest(opType, uriPath, virtualHost)
	if e != nil {
		return nil, nil, e
	}

	returnedVirtualHost := VirtualHost{}

	resp, e := s.client.Do(req, &returnedVirtualHost)
	if e != nil {
		return nil, resp, e
	}

	return &returnedVirtualHost, resp, e

}
