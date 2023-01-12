package main

import (
	"net/http"
	"sync"
)

//The bulkhead pattern is a design pattern that is used to isolate failures in a microservice architecture. It works by
//partitioning the system into several "bulkheads", each of which is responsible for a specific subset of the overall
//system. This allows the system to continue functioning even if one bulkhead fails, as the other bulkheads can still
//handle requests.

//To implement the bulkhead pattern in Golang using REST, you would need to create a server that is able to handle
//requests and distribute them to the appropriate bulkheads. You can do this by using a load balancer or by implementing
//a custom request routing mechanism.

//The bulkhead pattern is different from the circuit breaker pattern in that the circuit breaker pattern is used to
//prevent a service from being overwhelmed by too many requests, while the bulkhead pattern is used to isolate failures
//within the system. The circuit breaker pattern works by opening a "circuit" when the service is overloaded, and only
//allowing a limited number of requests to pass through until the service has recovered. The bulkhead pattern, on the
//other hand, does not block requests, but instead routes them to a different bulkhead if one fails.

// Bulkhead is a type that represents a single bulkhead
type Bulkhead struct {
	// mu is a mutex used to synchronize access to the bulkhead
	mu sync.Mutex

	// requests is a slice that holds the requests currently being processed by the bulkhead
	requests []*http.Request

	// capacity is the maximum number of requests that the bulkhead can handle at a time
	capacity int
}

// NewBulkhead creates a new bulkhead with the given capacity
func NewBulkhead(capacity int) *Bulkhead {
	return &Bulkhead{
		capacity: capacity,
	}
}

// HandleRequest handles a single request by adding it to the bulkhead's request slice and processing it
func (b *Bulkhead) HandleRequest(r *http.Request) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Add the request to the bulkhead's request slice
	b.requests = append(b.requests, r)

	// Process the request
	// (actual request processing code goes here)
}

// LoadBalancer is a type that distributes requests to different bulkheads
type LoadBalancer struct {
	// bulkheads is a slice of all the bulkheads managed by the load balancer
	bulkheads []*Bulkhead
}

// NewLoadBalancer creates a new load balancer with the given bulkheads
func NewLoadBalancer(bulkheads []*Bulkhead) *LoadBalancer {
	return &LoadBalancer{
		bulkheads: bulkheads,
	}
}

// ServeHTTP implements the http.Handler interface, allowing the load balancer to act as an HTTP server
func (l *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Find the bulkhead with the lowest number of requests
	var leastLoaded *Bulkhead
	for _, b := range l.bulkheads {
		b.mu.Lock()
		if leastLoaded == nil || len(b.requests) < len(leastLoaded.requests) {
			leastLoaded = b
		}
		b.mu.Unlock()
	}

	// Pass the request to the least loaded bulkhead
	leastLoaded.HandleRequest(r)
}

func main() {
	// Create a load balancer with three bulkheads
	bulkheads := []*Bulkhead{
		NewBulkhead(5),
		NewBulkhead(5),
		NewBulkhead(5),
	}
	lb := NewLoadBalancer(bulkheads)

	// Start the HTTP server
	http.ListenAndServe(":8080", lb)
}
