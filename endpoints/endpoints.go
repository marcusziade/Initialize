// Description: This package contains all the endpoints of the application.
package endpoints

import "net/http"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var httpClient = &http.Client{}

type Endpoints struct {
	httpClient HttpClient
}

func NewEndpoints(client HttpClient) *Endpoints {
	return &Endpoints{httpClient: client}
}
