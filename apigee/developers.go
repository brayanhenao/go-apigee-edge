package apigee

import (
	"errors"
	"net/url"
	"path"
)

const developersPath = "developers"

// DevelopersService is an interface for interfacing with the Apigee Edge Admin API
// dealing with developers.
type DevelopersService interface {
	Approve(string) (*Response, error)
	GetApps(string) ([]DeveloperApp, *Response, error)
	Create(Developer) (*Developer, *Response, error)
	Delete(string) (*Developer, *Response, error)
	Get(string) (*Developer, *Response, error)
	List() ([]string, *Response, error)
	Revoke(string) (*Response, error)
	Update(Developer) (*Developer, *Response, error)
}

type DevelopersServiceOp struct {
	client *ApigeeClient
}

var _ DevelopersService = &DevelopersServiceOp{}

// Developer contains information about a registered Developer within an Edge organization.
type Developer struct {
	Apps             []string    `json:"apps,omitempty"`
	Attributes       []Attribute `json:"attributes,omitempty"`
	Companies        []string    `json:"companies,omitempty"`
	Email            string      `json:"email,omitempty"`
	FirstName        string      `json:"firstName,omitempty"`
	Id               string      `json:"uuid,omitempty"`
	LastName         string      `json:"lastName,omitempty"`
	OrganizationName string      `json:"organizationName,omitempty"`
	Status           string      `json:"status,omitempty"` // active, inactive, ??
	UserName         string      `json:"userName,omitempty"`
}

func (s *DevelopersServiceOp) Update(dev Developer) (*Developer, *Response, error) {
	if dev.Email == "" && dev.Id == "" {
		return nil, nil, errors.New("must specify the Email or Id of the Developer to update")
	}
	// NB: it is legal to pass developer.Status, but that has no effect on the developer entity.
	// TODO (maybe): implement updating the status.
	var dpath string = ""
	if dev.Email != "" {
		dpath = path.Join(developersPath, dev.Email)
	} else {
		dpath = path.Join(developersPath, dev.Id)
	}

	req, e := s.client.NewRequest("POST", dpath, dev)
	if e != nil {
		return nil, nil, e
	}
	returnedDeveloper := Developer{}
	resp, e := s.client.Do(req, &returnedDeveloper)
	if e != nil {
		return nil, resp, e
	}
	return &returnedDeveloper, resp, e
}

func (s *DevelopersServiceOp) Create(dev Developer) (*Developer, *Response, error) {
	if dev.Id != "" {
		return nil, nil, errors.New("cannot create a developer with a specific Id")
	}
	req, e := s.client.NewRequest("POST", developersPath, dev)
	if e != nil {
		return nil, nil, e
	}
	returnedDeveloper := Developer{}
	resp, e := s.client.Do(req, &returnedDeveloper)
	if e != nil {
		return nil, resp, e
	}
	return &returnedDeveloper, resp, e
}

func (s *DevelopersServiceOp) Delete(devEmailOrId string) (*Developer, *Response, error) {
	path := path.Join(developersPath, devEmailOrId)
	req, e := s.client.NewRequest("DELETE", path, nil)
	if e != nil {
		return nil, nil, e
	}
	deletedDeveloper := Developer{}
	resp, e := s.client.Do(req, &deletedDeveloper)
	if e != nil {
		return nil, resp, e
	}
	return &deletedDeveloper, resp, e
}

func (s *DevelopersServiceOp) List() ([]string, *Response, error) {
	req, e := s.client.NewRequest("GET", developersPath, nil)
	if e != nil {
		return nil, nil, e
	}
	namelist := make([]string, 0)
	resp, e := s.client.Do(req, &namelist)
	if e != nil {
		return nil, resp, e
	}
	return namelist, resp, e
}

func (s *DevelopersServiceOp) Get(developerEmailOrId string) (*Developer, *Response, error) {
	devPath := path.Join(developersPath, developerEmailOrId)
	req, e := s.client.NewRequest("GET", devPath, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedDeveloper := Developer{}
	resp, e := s.client.Do(req, &returnedDeveloper)
	if e != nil {
		return nil, resp, e
	}
	return &returnedDeveloper, resp, e
}

func updateDeveloperStatus(s DevelopersServiceOp, developerEmailOrId string, desiredStatus string) (*Response, error) {

	devPath := path.Join(developersPath, developerEmailOrId)

	// append the necessary query param
	origURL, e := url.Parse(devPath)
	if e != nil {
		return nil, e
	}
	q := origURL.Query()
	q.Add("action", desiredStatus)
	origURL.RawQuery = q.Encode()
	devPath = origURL.String()

	req, e := s.client.NewRequest("POST", devPath, nil)
	if e != nil {
		return nil, e
	}
	resp, e := s.client.Do(req, nil)
	if e != nil {
		return resp, e
	}
	return resp, e
}

func (s *DevelopersServiceOp) Revoke(developerEmailOrId string) (*Response, error) {
	return updateDeveloperStatus(*s, developerEmailOrId, "inactive")
}

func (s *DevelopersServiceOp) Approve(developerEmailOrId string) (*Response, error) {
	return updateDeveloperStatus(*s, developerEmailOrId, "active")
}

func (s *DevelopersServiceOp) GetApps(developerEmailOrId string) ([]DeveloperApp, *Response, error) {
	appsPath := path.Join(developersPath, developerEmailOrId, "apps") + "?expand=true"
	req, e := s.client.NewRequest("GET", appsPath, nil)
	if e != nil {
		return nil, nil, e
	}
	apps := make([]DeveloperApp, 0)
	resp, e := s.client.Do(req, &apps)
	if e != nil {
		return nil, resp, e
	}
	return apps, resp, e
}
