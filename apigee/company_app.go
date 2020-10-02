package apigee

import (
	"path"
)

const appPath = "apps"

// CompanyAppsService is an interface for interfacing with the Apigee Edge Admin API
// dealing with companyApps.
type CompanyAppsService interface {
	Get(string, string) (*CompanyApp, *Response, error)
	Create(string, CompanyApp) (*CompanyApp, *Response, error)
	Delete(string, string) (*Response, error)
	Update(string, CompanyApp) (*CompanyApp, *Response, error)
}

type CompanyAppsServiceOp struct {
	client *ApigeeClient
}

var _ CompanyAppsService = &CompanyAppsServiceOp{}

type CompanyApp struct {
	Name        string       `json:"name,omitempty"`
	ApiProducts []string     `json:"apiProducts,omitempty"`
	Attributes  []Attribute  `json:"attributes,omitempty"`
	Scopes      []string     `json:"scopes,omitempty"`
	CallbackUrl string       `json:"callbackUrl,omitempty"`
	Credentials []Credential `json:"credentials,omitempty"`
	AppId       string       `json:"appId,omitempty"`
	CompanyName string       `json:"companyName,omitempty"`
	AppFamily   string       `json:"appFamily,omitempty"`
	Status      string       `json:"status,omitempty"`
}

func (s *CompanyAppsServiceOp) Get(companyName string, name string) (*CompanyApp, *Response, error) {

	path := path.Join(companiesPath, companyName, appPath, name)

	req, e := s.client.NewRequest("GET", path, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedCompanyApp := CompanyApp{}
	resp, e := s.client.Do(req, &returnedCompanyApp)
	if e != nil {
		return nil, resp, e
	}
	return &returnedCompanyApp, resp, e

}

func (s *CompanyAppsServiceOp) Create(companyName string, companyApp CompanyApp) (*CompanyApp, *Response, error) {

	return postOrPutCompanyApp(companyName, companyApp, "POST", s)

}

func (s *CompanyAppsServiceOp) Update(companyName string, companyApp CompanyApp) (*CompanyApp, *Response, error) {

	return postOrPutCompanyApp(companyName, companyApp, "PUT", s)

}

func (s *CompanyAppsServiceOp) Delete(companyName string, name string) (*Response, error) {

	path := path.Join(companiesPath, companyName, appPath, name)

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

func postOrPutCompanyApp(companyName string, companyApp CompanyApp, opType string, s *CompanyAppsServiceOp) (*CompanyApp, *Response, error) {

	uripath := ""

	if opType == "PUT" {
		uripath = path.Join(companiesPath, companyName, appPath, companyApp.Name)
	} else {
		uripath = path.Join(companiesPath, companyName, appPath)
	}

	req, e := s.client.NewRequest(opType, uripath, companyApp)
	if e != nil {
		return nil, nil, e
	}

	returnedCompanyApp := CompanyApp{}

	resp, e := s.client.Do(req, &returnedCompanyApp)
	if e != nil {
		return nil, resp, e
	}

	return &returnedCompanyApp, resp, e

}
