package main

import (
	"errors"
	"fmt"
	"sync"
)

// The CircuitBreaker pattern is used to protect a service from being overwhelmed by requests, by temporarily
// suspending requests to the service that is experiencing problems

// CircuitBreaker monitors the status of the service and decides when open/close a circuit, based on consecuitve
// error rate.
type CircuitBreaker struct {
	threshold int
	failures  int
	lock      sync.Mutex
	state     bool
}

func (cb *CircuitBreaker) AllowRequest() bool {
	cb.lock.Lock()
	defer cb.lock.Unlock()
	if cb.state {
		// Circuit is open, do not allow requests
		return false
	}
	// Circuit is closed, allow requests
	return true
}

func (cb *CircuitBreaker) RecordFailure() {
	cb.lock.Lock()
	defer cb.lock.Unlock()
	cb.failures++
	if cb.failures >= cb.threshold {
		// Open the circuit
		cb.state = true
		cb.failures = 0
	}
}

func (cb *CircuitBreaker) RecordSuccess() {
	cb.lock.Lock()
	defer cb.lock.Unlock()
	cb.failures = 0
	if cb.state {
		// Close the circuit
		cb.state = false
	}
}

func callService() error {
	return errors.New("Zero cannot be used")
}
func main() {
	cb := &CircuitBreaker{
		threshold: 3,
		state:     false,
	}

	for i := 0; i < 10; i++ {
		if cb.AllowRequest() {
			// Call the service
			err := callService()
			if err != nil {
				// Record the failure
				cb.RecordFailure()
			} else {
				// Record the success
				cb.RecordSuccess()
			}
		} else {
			fmt.Println("Circuit is open, skipping request")
		}
	}
}

//In this example, the circuit breaker is initialized with a threshold of 3 failures before the circuit is opened. The
//main loop makes 10 requests to the service, checking the circuit breaker's state before each request. If the circuit
//is open, the request is skipped. If the circuit is closed, the request is made, and the circuit breaker is updated
//based on the result of the request.
//
//This is just a basic example of how you could use a circuit breaker in Go. You may want to customize the behavior of
//your circuit breaker based on your specific needs. For example, you might want to add a timer to close the circuit
//after a certain period of time, or add additional criteria for opening or closing the circuit.
