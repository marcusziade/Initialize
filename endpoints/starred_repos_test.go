package endpoints_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	endpoints "github.com/marcusziade/github-api/endpoints"

	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	jsonResponse := `[{"id": 1, "name": "Repo1"}, {"id": 2, "name": "Repo2"}]`

	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(jsonResponse)),
	}, nil
}

func TestGetStarredRepos(t *testing.T) {
	e := endpoints.NewEndpoints(&MockClient{})

	username := "testUser"
	token := "testToken"
	pages := 1

	starredRepos, err := e.GetStarredRepos(username, token, pages)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(starredRepos))
	assert.Equal(t, "Repo1", starredRepos[0].Name)
	assert.Equal(t, "Repo2", starredRepos[1].Name)
}
