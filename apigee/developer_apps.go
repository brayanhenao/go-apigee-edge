package apigee

import (
	"errors"
	"net/url"
	"path"
)

// DeveloperAppsService is an interface for interfacing with the Apigee Edge Admin API
// dealing with apps that belong to a particular developer.
type DeveloperAppsService interface {
	Approve(string, string) (*Response, error)
	Create(string, DeveloperApp) (*DeveloperApp, *Response, error)
	Delete(string, string) (*DeveloperApp, *Response, error)
	Get(string, string) (*DeveloperApp, *Response, error)
	List(string) ([]string, *Response, error)
	Revoke(string, string) (*Response, error)
	Update(string, DeveloperApp) (*DeveloperApp, *Response, error)
}

type DeveloperAppsServiceOp struct {
	client *ApigeeClient
}

var _ DeveloperAppsService = &DeveloperAppsServiceOp{}

// DeveloperApp holds information about a registered DeveloperApp.
type DeveloperApp struct {
	ApiProducts      []string     `json:"apiProducts,omitempty"`
	Attributes       []Attribute  `json:"attributes,omitempty"`
	CallbackUrl      string       `json:"callbackUrl,omitempty"`
	Credentials      []Credential `json:"credentials,omitempty"`
	DeveloperId      string       `json:"developerId,omitempty"`
	Id               string       `json:"appId,omitempty"`
	InitialKeyExpiry string       `json:"keyExpiresIn,omitempty"`
	Name             string       `json:"name,omitempty"`
	Scopes           []string     `json:"scopes,omitempty"`
	Status           string       `json:"status,omitempty"`
}

func (s *DeveloperAppsServiceOp) Create(developerEmail string, app DeveloperApp) (*DeveloperApp, *Response, error) {
	if app.Name == "" {
		return nil, nil, errors.New("cannot create a developerapp with no name")
	}
	appsPath := path.Join(developersPath, developerEmail, appPath)
	req, e := s.client.NewRequest("POST", appsPath, app)
	if e != nil {
		return nil, nil, e
	}
	returnedApp := DeveloperApp{}
	resp, e := s.client.Do(req, &returnedApp)
	if e != nil {
		return nil, resp, e
	}
	return &returnedApp, resp, e
}

func (s *DeveloperAppsServiceOp) Delete(developerEmail string, appName string) (*DeveloperApp, *Response, error) {
	path := path.Join(developersPath, developerEmail, appPath, appName)
	req, e := s.client.NewRequest("DELETE", path, nil)
	if e != nil {
		return nil, nil, e
	}
	deletedApp := DeveloperApp{}
	resp, e := s.client.Do(req, &deletedApp)
	if e != nil {
		return nil, resp, e
	}
	return &deletedApp, resp, e
}

func updateAppStatus(s DeveloperAppsServiceOp, developerEmail string, appName string, desiredStatus string) (*Response, error) {

	appPath := path.Join(developersPath, developerEmail, appPath, appName)

	// append the necessary query param
	origURL, e := url.Parse(appPath)
	if e != nil {
		return nil, e
	}
	q := origURL.Query()
	q.Add("action", desiredStatus)
	origURL.RawQuery = q.Encode()
	appPath = origURL.String()

	req, e := s.client.NewRequest("POST", appPath, nil)
	if e != nil {
		return nil, e
	}
	resp, e := s.client.Do(req, nil)
	if e != nil {
		return resp, e
	}
	return resp, e
}

func (s *DeveloperAppsServiceOp) Revoke(developerEmail string, appName string) (*Response, error) {
	return updateAppStatus(*s, developerEmail, appName, "revoke")
}

func (s *DeveloperAppsServiceOp) Approve(developerEmail string, appName string) (*Response, error) {
	return updateAppStatus(*s, developerEmail, appName, "approve")
}

func (s *DeveloperAppsServiceOp) List(developerEmail string) ([]string, *Response, error) {
	appsPath := path.Join(developersPath, developerEmail, appPath)
	req, e := s.client.NewRequest("GET", appsPath, nil)
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

func (s *DeveloperAppsServiceOp) Get(developerEmail string, appName string) (*DeveloperApp, *Response, error) {
	appPath := path.Join(developersPath, developerEmail, appPath, appName)
	req, e := s.client.NewRequest("GET", appPath, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedApp := DeveloperApp{}
	resp, e := s.client.Do(req, &returnedApp)
	if e != nil {
		return nil, resp, e
	}
	return &returnedApp, resp, e
}

func (s *DeveloperAppsServiceOp) Update(developerEmail string, app DeveloperApp) (*DeveloperApp, *Response, error) {
	if app.Name == "" {
		return nil, nil, errors.New("missing the Name of the App to update")
	}
	appPath := path.Join(developersPath, developerEmail, appPath, app.Name)

	req, e := s.client.NewRequest("POST", appPath, app)
	if e != nil {
		return nil, nil, e
	}
	returnedApp := DeveloperApp{}
	resp, e := s.client.Do(req, &returnedApp)
	if e != nil {
		return nil, resp, e
	}
	return &returnedApp, resp, e
}
