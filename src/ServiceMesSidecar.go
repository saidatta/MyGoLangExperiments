package main

import (
	"io"
	"net/http"
)

//To integrate the UserServiceSidecar and OrderServiceSidecar with the UserService and OrderService using a service mesh
//architecture, you will need to do the following:
//
//Deploy the UserService and OrderService in your service mesh. This will typically involve creating Kubernetes pods or
//other containerized deployments for the services, and configuring the service mesh to route traffic to them.
//Deploy the UserServiceSidecar and OrderServiceSidecar in your service mesh as separate deployments from the UserService
//and OrderService. The sidecars should be deployed alongside the corresponding service, so that they are in the same
//network namespace and can communicate with the service using localhost.
//Configure the service mesh to route traffic to the sidecars instead of the UserService and OrderService. This will
//typically involve creating virtual service and destination rule objects in the service mesh to specify the routing rules.
//Modify the UserService and OrderService code to listen for requests on a different port than the sidecars. This will
//allow the sidecars to forward requests to the service on a different port, while still being able to communicate with
//the service using localhost.
//Modify the UserServiceSidecar and OrderServiceSidecar code to forward requests to the UserService and OrderService on
//the correct port. This will typically involve setting the Host field of the HTTP request to localhost and the correct port for the service.
//Test that the sidecars and services are working correctly by sending requests to the sidecars and verifying that the
//correct responses are returned.

// UserServiceSidecar is a sidecar proxy for the UserService.
type UserServiceSidecar struct {
	client *http.Client
}

func NewUserServiceSidecar() *UserServiceSidecar {
	return &UserServiceSidecar{
		client: &http.Client{},
	}
}

func (s *UserServiceSidecar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Forward the request to the actual UserService.
	resp, err := s.client.Do(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response from the UserService to the client.
	for k, vals := range resp.Header {
		for _, val := range vals {
			w.Header().Add(k, val)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// OrderServiceSidecar is a sidecar proxy for the OrderService.
type OrderServiceSidecar struct {
	client *http.Client
}

func NewOrderServiceSidecar() *OrderServiceSidecar {
	return &OrderServiceSidecar{
		client: &http.Client{},
	}
}

func (s *OrderServiceSidecar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Forward the request to the actual OrderService.
	resp, err := s.client.Do(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response from the OrderService to the client.
	for k, vals := range resp.Header {
		for _, val := range vals {
			w.Header().Add(k, val)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
