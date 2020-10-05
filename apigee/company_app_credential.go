package apigee

import (
	"path"
)

const keysPath = "keys"

// CompanyAppCredentialsService is an interface for interfacing with the Apigee Edge Admin API
// dealing with companyApp credentials/keys.
type CompanyAppCredentialsService interface {
	Create(string, string, Credential) (*Credential, *Response, error)
	Delete(string, string, string) (*Response, error)
	Get(string, string, string) (*Credential, *Response, error)
	RemoveApiProduct(string, string, string, string) (*Response, error)
	Update(string, string, string, Credential) (*Credential, *Response, error)
}

type CompanyAppCredentialsServiceOp struct {
	client *ApigeeClient
}

var _ CompanyAppCredentialsService = &CompanyAppCredentialsServiceOp{}

// Create a company app's consumer key and secret
func (s *CompanyAppCredentialsServiceOp) Create(companyName string, appName string, companyAppCredential Credential) (*Credential, *Response, error) {

	uripath := path.Join(companiesPath, companyName, appPath, appName, keysPath, "create")

	req, e := s.client.NewRequest("POST", uripath, companyAppCredential)
	if e != nil {
		return nil, nil, e
	}

	returnedCompanyAppCredentials := Credential{}

	resp, e := s.client.Do(req, &returnedCompanyAppCredentials)
	if e != nil {
		return nil, resp, e
	}

	return &returnedCompanyAppCredentials, resp, e

}

// Update existing company app's consumer key with new API products or attributes
func (s *CompanyAppCredentialsServiceOp) Update(companyName string, appName string, consumerKey string, companyAppCredential Credential) (*Credential, *Response, error) {

	uripath := path.Join(companiesPath, companyName, appPath, appName, keysPath, consumerKey)

	req, e := s.client.NewRequest("POST", uripath, companyAppCredential)
	if e != nil {
		return nil, nil, e
	}

	returnedCompanyAppCredentials := Credential{}

	resp, e := s.client.Do(req, &returnedCompanyAppCredentials)
	if e != nil {
		return nil, resp, e
	}

	return &returnedCompanyAppCredentials, resp, e

}

// Get information about a company app's consumer key
func (s *CompanyAppCredentialsServiceOp) Get(companyName string, appName string, consumerKey string) (*Credential, *Response, error) {

	uripath := path.Join(companiesPath, companyName, appPath, appName, keysPath, consumerKey)

	req, e := s.client.NewRequest("GET", uripath, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedCompanyAppCredential := Credential{}
	resp, e := s.client.Do(req, &returnedCompanyAppCredential)
	if e != nil {
		return nil, resp, e
	}
	return &returnedCompanyAppCredential, resp, e

}

// Delete a company app's consumer key
func (s *CompanyAppCredentialsServiceOp) Delete(companyName string, appName string, consumerKey string) (*Response, error) {

	uripath := path.Join(companiesPath, companyName, appPath, appName, keysPath, consumerKey)

	req, e := s.client.NewRequest("DELETE", uripath, nil)
	if e != nil {
		return nil, e
	}

	resp, e := s.client.Do(req, nil)
	if e != nil {
		return resp, e
	}

	return resp, e

}

// Remove an API product from a company app's consumer key
func (s *CompanyAppCredentialsServiceOp) RemoveApiProduct(companyName string, appName string, consumerKey string, apiProductName string) (*Response, error) {

	uripath := path.Join(companiesPath, companyName, appPath, appName, keysPath, consumerKey, productsPath, apiProductName)

	req, e := s.client.NewRequest("DELETE", uripath, nil)
	if e != nil {
		return nil, e
	}

	resp, e := s.client.Do(req, nil)
	if e != nil {
		return resp, e
	}

	return resp, e

}
