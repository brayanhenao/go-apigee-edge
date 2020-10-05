package apigee

import (
	"path"
)

const companiesPath = "companies"

// CompanyService is an interface for interfacing with the Apigee Edge Admin API
// dealing with companies.
type CompaniesService interface {
	Create(Company) (*Company, *Response, error)
	Delete(string) (*Response, error)
	Get(string) (*Company, *Response, error)
	Update(Company) (*Company, *Response, error)
}

type CompaniesServiceOp struct {
	client *ApigeeClient
}

var _ CompaniesService = &CompaniesServiceOp{}

type Company struct {
	Apps        []string    `json:"apps,omitempty"`
	Attributes  []Attribute `json:"attributes,omitempty"`
	DisplayName string      `json:"displayName,omitempty"`
	Name        string      `json:"name,omitempty"`
	Status      string      `json:"status,omitempty"`
}

func (s *CompaniesServiceOp) Get(name string) (*Company, *Response, error) {

	path := path.Join(companiesPath, name)

	req, e := s.client.NewRequest("GET", path, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedCompany := Company{}
	resp, e := s.client.Do(req, &returnedCompany)
	if e != nil {
		return nil, resp, e
	}
	return &returnedCompany, resp, e

}

func (s *CompaniesServiceOp) Create(company Company) (*Company, *Response, error) {

	return postOrPutCompany(company, "POST", s)

}

func (s *CompaniesServiceOp) Update(company Company) (*Company, *Response, error) {

	return postOrPutCompany(company, "PUT", s)

}

func (s *CompaniesServiceOp) Delete(name string) (*Response, error) {

	path := path.Join(companiesPath, name)

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

func postOrPutCompany(company Company, opType string, s *CompaniesServiceOp) (*Company, *Response, error) {

	uripath := ""

	if opType == "PUT" {
		uripath = path.Join(companiesPath, company.Name)
	} else {
		uripath = path.Join(companiesPath)
	}

	req, e := s.client.NewRequest(opType, uripath, company)
	if e != nil {
		return nil, nil, e
	}

	returnedCompany := Company{}

	resp, e := s.client.Do(req, &returnedCompany)
	if e != nil {
		return nil, resp, e
	}

	return &returnedCompany, resp, e

}
