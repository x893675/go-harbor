package goharbor

import (
	"context"
	"encoding/json"
	"github.com/x893675/go-harbor/schema"
	"net/http"
	"testing"
)

func TestListProjects(t *testing.T) {
	t.Parallel()
	jsonProjects := `
[
  {
    "project_id": 1,
    "owner_id": 1,
    "name": "library",
    "creation_time": "2020-07-16T07:07:48.524986Z",
    "update_time": "2020-07-16T07:07:48.524986Z",
    "deleted": false,
    "owner_name": "",
    "current_user_role_id": 1,
    "current_user_role_ids": [
      1
    ],
    "repo_count": 1,
    "chart_count": 0,
    "metadata": {
      "public": "true"
    },
    "cve_whitelist": {
      "id": 0,
      "project_id": 0,
      "items": null,
      "creation_time": "0001-01-01T00:00:00Z",
      "update_time": "0001-01-01T00:00:00Z"
    }
  }
]`
	var expectedProjects []schema.Project
	err := json.Unmarshal([]byte(jsonProjects), &expectedProjects)
	if err != nil {
		t.Fatal(err)
	}
	//var req http.Request
	//server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte{1, 0, 0, 0, 0, 0, 0, 19})
	//	w.Write([]byte("something happened!"))
	//	req = *r
	//}))
	//defer server.Close()
	fakeRT := &FakeRoundTripper{message: jsonProjects, status: http.StatusOK}
	client := newTestClient(fakeRT)
	projects, err := client.ListProjects(context.TODO(), schema.ProjectListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(projects) != len(expectedProjects) {
		t.Errorf("ListProjects: wrong resp length. Want %q. Got %q.", len(expectedProjects), len(projects))
	}
}
