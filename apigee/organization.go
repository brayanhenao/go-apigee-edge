package apigee

import (
	"path"
)

const organizationsPath = "organizations"

// OrganizationsService is an interface for interfacing with the Apigee Edge Admin API
// querying Edge environments.
type OrganizationService interface {
	Get(string) (*Organization, *Response, error)
}

type OrganizationServiceOp struct {
	client *ApigeeClient
}

var _ OrganizationService = &OrganizationServiceOp{}

// {
//   "createdAt" : 1371096055089,
//   "createdBy" : "noreply_admin@apigee.com",
//   "lastModifiedAt" : 1456865874610,
//   "lastModifiedBy" : "noreply_cpsadmin@apigee.com",
//   "displayName" : "cheeso",
//   "environments" : [ "test", "prod" ],
//   "name" : "cheeso",
//   "properties" : {
//     "property" : [ {
//       "name" : "features.isCpsEnabled",
//       "value" : "true"
//     } ]
//   },
//   "type" : "trial"
// }

type Organization struct {
	CreatedAt      Timestamp   `json:"createdAt,omitempty"`
	CreatedBy      string      `json:"createdBy,omitempty"`
	DisplayName    string      `json:"displayName,omitempty"`
	Environments   []string    `json:"environments,omitempty"`
	LastModifiedAt Timestamp   `json:"lastModifiedAt,omitempty"`
	LastModifiedBy string      `json:"lastModifiedBy,omitempty"`
	Name           string      `json:"name,omitempty"`
	Properties     []Attribute `json:"properties,omitempty"`
	Type           string      `json:"type,omitempty"`
}

// Get retrieves the information about an Organization, information including
// the properties, and the created and last modified details, the list of Environments,
// etc.
func (s *OrganizationServiceOp) Get(org string) (*Organization, *Response, error) {
	orgPath := ""
	if org != "" {
		orgPath = path.Join(organizationsPath, org)
	}
	req, e := s.client.NewRequest("GET", orgPath, nil)
	if e != nil {
		return nil, nil, e
	}
	returnedOrg := Organization{}
	resp, e := s.client.Do(req, &returnedOrg)
	if e != nil {
		return nil, resp, e
	}
	return &returnedOrg, resp, e
}
