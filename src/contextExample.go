package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//In Go, the context package provides a way to carry around request-scoped values and cancelation signals across API
//boundaries. It is often used in conjunction with network operations, such as HTTP requests, to allow for request
//cancelation and timeout.
//
//The context package provides the Context type, which is an immutable value that carries request-scoped data. A Context
//value can be created using the context.WithValue and context.WithCancel functions, which allow you to attach values
//and cancelation signals to the context. The Context value is then passed as an argument to functions that need to
//perform the corresponding operation.
//
//Here is an example of how you might use the context package to cancel an HTTP request after a certain amount of time:

func main() {
	// Create a context with a cancelation signal
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Perform the HTTP request using the context
	req, _ := http.NewRequest("GET", "http://google.com", nil)
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Check if the context was canceled
		if ctx.Err() == context.Canceled {
			fmt.Println("Request canceled")
		} else {
			fmt.Println("Request failed:", err)
		}
		return
	}

	// Do something with the response
	fmt.Println("Response status:", res.Status)
}
