package apigee

const sharedFlowPath = "sharedflows"

// SharedFlowsService is an interface for interfacing with the Apigee Admin API
// dealing with apiproxies.
type SharedFlowsService interface {
	Delete(string) (*DeletedItemInfo, *Response, error)
	DeleteRevision(string, Revision) (*DeployableRevision, *Response, error)
	Deploy(string, string, Revision, bool, int) (*RevisionDeployment, *Response, error)
	Export(string, Revision) (string, *Response, error)
	Get(string) (*DeployableAsset, *Response, error)
	GetDeployments(string) (*Deployment, *Response, error)
	Import(string, string) (*DeployableRevision, *Response, error)
	List() ([]string, *Response, error)
	Undeploy(string, string, Revision) (*RevisionDeployment, *Response, error)
}

type SharedFlowsServiceOp struct {
	client     *ApigeeClient
	deployable Deployable
}

var _ SharedFlowsService = &SharedFlowsServiceOp{}

func (s *SharedFlowsServiceOp) List() ([]string, *Response, error) {
	return s.deployable.List(s.client, sharedFlowPath)
}

func (s *SharedFlowsServiceOp) Get(proxyName string) (*DeployableAsset, *Response, error) {
	return s.deployable.Get(s.client, sharedFlowPath, proxyName)
}

func (s *SharedFlowsServiceOp) Import(proxyName string, source string) (*DeployableRevision, *Response, error) {
	return s.deployable.Import(s.client, sharedFlowPath, proxyName, source)
}

func (s *SharedFlowsServiceOp) Export(proxyName string, rev Revision) (string, *Response, error) {
	return s.deployable.Export(s.client, sharedFlowPath, proxyName, rev)
}

func (s *SharedFlowsServiceOp) DeleteRevision(proxyName string, rev Revision) (*DeployableRevision, *Response, error) {
	return s.deployable.DeleteRevision(s.client, sharedFlowPath, proxyName, rev)
}

func (s *SharedFlowsServiceOp) Undeploy(proxyName, env string, rev Revision) (*RevisionDeployment, *Response, error) {
	return s.deployable.Undeploy(s.client, sharedFlowPath, proxyName, env, rev)
}

func (s *SharedFlowsServiceOp) Deploy(proxyName, env string, rev Revision, override bool, delay int) (*RevisionDeployment, *Response, error) {
	return s.deployable.Deploy(s.client, sharedFlowPath, proxyName, "", env, rev, override, delay)
}

func (s *SharedFlowsServiceOp) Delete(proxyName string) (*DeletedItemInfo, *Response, error) {
	return s.deployable.Delete(s.client, sharedFlowPath, proxyName)
}

func (s *SharedFlowsServiceOp) GetDeployments(proxyName string) (*Deployment, *Response, error) {
	return s.deployable.GetDeployments(s.client, sharedFlowPath, proxyName)
}
