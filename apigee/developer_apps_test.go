package apigee

import (
	"encoding/json"
	"testing"
)

const (
	appJson1 = `{
  "attributes": [ {
    "name" : "tag1",
    "value" : "created by golang" }],
  "name" : "will-be-replaced",
  "keyExpiresIn" : "3600000"
}`
)

func randomAppFromTemplate() (DeveloperApp, error) {
	got := DeveloperApp{}
	e := json.Unmarshal([]byte(appJson1), &got)

	if e != nil {
		return got, e
	}
	// assign values
	tag := testPrefix + randomString(7)
	got.Name = tag + "app"
	return got, e
}

func TestDeveloperAppCreateDelete(t *testing.T) {
	client := NewClientForTesting(t)
	dev, e := randomDeveloperFromTemplate()
	createdDeveloper, resp, e := client.Developers.Create(dev)
	if e != nil {
		t.Errorf("while creating Edge developer, error:\n%#v\n", e)
		return
	}
	t.Logf("Create: got=%+v", createdDeveloper)
	t.Logf("resp: got=%+v", resp)

	teardown := func(t *testing.T) {
		t.Logf("remove developer")
		deletedDeveloper, _, e := client.Developers.Delete(createdDeveloper.Email)
		if e != nil {
			t.Errorf("while deleting Edge developer, error:\n%#v\n", e)
			return
		}
		t.Logf("Delete: got=%v", deletedDeveloper)
	}

	defer teardown(t)
	wait(1)

	devapp, e := randomAppFromTemplate()

	createdApp, resp, e := client.DeveloperApps.Create(createdDeveloper.Email, devapp)
	if e != nil {
		t.Errorf("while creating developer app, error:\n%#v\n", e)
		return
	}
	t.Logf("CreateApp: got=%v", createdApp)

	wait(1)

	resp, e = client.DeveloperApps.Revoke(createdDeveloper.Email, createdApp.Name)
	if e != nil {
		t.Errorf("while revoking developer app, error:\n%#v\n", e)
		return
	}
	t.Logf("RevokeApp")
	wait(1)

	got, resp, e := client.DeveloperApps.Get(createdDeveloper.Email, createdApp.Name)
	if e != nil {
		t.Errorf("while getting developer app, error:\n%#v\n", e)
		return
	}
	if got.Name != createdApp.Name {
		t.Errorf("inconsistent name")
	}
	if got.Status != "revoked" {
		t.Errorf("inconsistent status")
	}
	t.Logf("GetApp")

	resp, e = client.DeveloperApps.Approve(createdDeveloper.Email, createdApp.Name)
	if e != nil {
		t.Errorf("while approving developer app, error:\n%#v\n", e)
		return
	}
	t.Logf("ApproveApp")

	wait(1)

	got, resp, e = client.DeveloperApps.Get(createdDeveloper.Email, createdApp.Name)
	if e != nil {
		t.Errorf("while getting developer app, error:\n%#v\n", e)
		return
	}
	if got.Name != createdApp.Name {
		t.Errorf("inconsistent name")
	}
	if got.Status != "approved" {
		t.Errorf("inconsistent status")
	}
	t.Logf("GetApp")

	deletedApp, resp, e := client.DeveloperApps.Delete(createdDeveloper.Email, createdApp.Name)
	if e != nil {
		t.Errorf("while creating developer app, error:\n%#v\n", e)
		return
	}
	t.Logf("DeleteApp: got=%v", deletedApp)

	wait(1)

}
