package store

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"log"
	"net/http"
	"os"
	"time"
)

const eventStoreName = "eventstore"
const partitionKeyMetadata = "partitionKey"

type Store struct {
	dc dapr.Client
}

func New(daprClient dapr.Client) *Store {
	return &Store{
		dc: daprClient,
	}
}

func (s *Store) GetAggregate(ctx context.Context, partitionKey string, aggregate string, aggregateId string, fromVersion int) any {
	c, err := dapr.NewClient()
	if err != nil {
		return nil
	}

	query := fmt.Sprintf(
		"SELECT * FROM c WHERE c.bucketId = '%s' AND c.aggregateId = '%s' AND c.aggregate = '%s' AND c.version >= %d ORDER BY c.version",
		partitionKey,
		aggregateId, aggregate,
		fromVersion)

	response, err := c.QueryStateAlpha1(ctx, eventStoreName, query, map[string]string{partitionKey: partitionKey})
	if err != nil {
		return err
	}
	log.Println(response.Results)
	return nil
}

func (s *Store) Save(ctx context.Context, id string, partitionKey string, event any) error {
	daprHost := os.Getenv("DAPR_HOST")
	if daprHost == "" {
		daprHost = "http://localhost"
	}
	daprHttpPort := os.Getenv("DAPR_HTTP_PORT")
	if daprHttpPort == "" {
		daprHttpPort = "3500"
	}

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	b, err := json.Marshal([]any{event})
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", daprHost+":"+daprHttpPort+"/v1.0/state/"+eventStoreName, bytes.NewReader(b))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	log.Println("id", id)
	log.Println("partitionKey", partitionKey)
	log.Println("payload", string(b))

	resp, err := client.Do(request)
	if err != nil {
		json.NewEncoder(os.Stdout).Encode(resp.Body)
		return err
	}

	d := map[string]any{}
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return err
	}

	log.Printf("%v", d)
	if err != nil {
		return err
	}

	return s.dc.SaveState(ctx, eventStoreName, id, b, nil)
}
