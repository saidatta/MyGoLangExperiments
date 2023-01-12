package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ReplicaShard struct{}

type Operation struct {
	IsIndexOperation bool
}

func (r *ReplicaShard) ExecuteOperation(op Operation) error {

}

type PrimaryShard struct {
	index    string
	translog string
	// replicas is a slice of pointers to the replicas of this primary shard
	replicas []*ReplicaShard
}

// ExecuteOperation processes an incoming operation on the primary shard.
// It returns an error if the operation is invalid or if there was an error executing it.
func (p *PrimaryShard) ExecuteOperation(op Operation) error {
	// Validate the incoming operation
	if err := p.validateOperation(op); err != nil {
		return fmt.Errorf("invalid operation: %w", err)
	}

	// Execute the operation locally
	if err := p.executeOperationLocally(op); err != nil {
		return fmt.Errorf("error executing operation locally: %w", err)
	}

	// Forward the operation to each replica in the current in-sync copies set
	var wg sync.WaitGroup
	for _, r := range p.replicas {
		wg.Add(1)
		go func(r *ReplicaShard) {
			defer wg.Done()
			if err := r.ExecuteOperation(op); err != nil {
				log.Printf("error executing operation on replica %v: %v", r, err)
			}
		}(r)
	}
	wg.Wait()

	// Acknowledge the successful completion of the request to the client
	return nil
}

// validateOperation checks if the operation is structurally valid.
// It returns an error if the operation is invalid.
func (p *PrimaryShard) validateOperation(op Operation) error {
	// Check if the operation has a valid structure
	if op.ObjectField != nil && !op.IsNumberExpected {
		return fmt.Errorf("object field where a number is expected")
	}
	return nil
}

// executeOperationLocally inserts the operation into the translog and performs the operation on the local primary shard.
// It returns an error if there was an error executing the operation.
func (p *PrimaryShard) executeOperationLocally(op Operation) error {
	// Validate the content of fields
	if err := p.validateFields(op); err != nil {
		return fmt.Errorf("error validating fields: %w", err)
	}

	// Insert the operation into the translog
	if err := p.translog.Insert(op); err != nil {
		return fmt.Errorf("error inserting operation into translog: %w", err)
	}

	// Index or delete the relevant document
	if err := p.indexOrDeleteDocument(op); err != nil {
		return fmt.Errorf("error indexing or deleting document: %w", err)
	}
	return nil
}

// indexOrDeleteDocument indexes or deletes the relevant document based on the operation.
// It returns an error if there was an error executing the operation.
func (p *PrimaryShard) indexOrDeleteDocument(op Operation) error {
	if op.IsIndexOperation {
		// Index the document
		return p.indexDocument(op.Document)
	}
	// Delete the document
	return p.deleteDocument(op.DocumentID)
}

// validateFields checks if the fields of the operation are valid.
// It returns an error if any of the fields are invalid.
// validateFields checks if the fields of the operation are valid.
// It returns an error if any of the fields are invalid.
func (p *PrimaryShard) validateFields(op Operation) error {
	// Check if the keyword value is too long for indexing in Lucene
	if len(op.KeywordValue) > maxKeywordLength {
		return fmt.Errorf("keyword value is too long for indexing in Lucene")
	}
	return nil
}

// indexDocument indexes the given document in the primary shard.
// It returns an error if there was an error executing the operation.
func (p *PrimaryShard) indexDocument(doc Document) error {
	// Create a new document to be indexed
	indexDoc := lucene.NewDocument()

	// Add the fields of the document to the index document
	indexDoc.AddField("id", doc.ID, lucene.STORED)
	indexDoc.AddField("title", doc.Title, lucene.STORED)
	indexDoc.AddField("content", doc.Content, lucene.STORED)

	// Add the document to the index
	if err := p.index.AddDocument(indexDoc); err != nil {
		return fmt.Errorf("error adding document to index: %w", err)
	}
	return nil
}

func main() {
	// Create a new primary shard with a slice of replicas
	primaryShard := &PrimaryShard{
		replicas: []*ReplicaShard{
			&ReplicaShard{},
			&ReplicaShard{},
			&ReplicaShard{},
		},
	}

	// Start a goroutine to listen for incoming operations
	go func() {
		for {
			// Wait for an incoming operation
			op := <-incomingOperations

			// Execute the operation on the primary shard
			if err := primaryShard.ExecuteOperation(op); err != nil {
				log.Printf("error executing operation: %v", err)
			}
		}
	}()

	// Wait for SIGINT or SIGTERM signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
