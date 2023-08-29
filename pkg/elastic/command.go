package elastic

import (
	"context"
	"encoding/json"
	"strings"
)

func (e *Elastic) Put(ctx context.Context, index string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = e.client.Index(index, strings.NewReader(string(data)))
	if err != nil {
		return err
	}

	return nil
}

func (e *Elastic) Search(ctx context.Context, index string, query map[string]interface{}) ([]interface{}, error) {
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	searchResponse, err := e.client.Search(
		e.client.Search.WithContext(ctx),
		e.client.Search.WithIndex(index),
		e.client.Search.WithBody(strings.NewReader(string(queryJSON))),
	)

	if err != nil {
		return nil, err
	}

	defer searchResponse.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(searchResponse.Body).Decode(&response); err != nil {
		return nil, err
	}

	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})

	return hits, nil
}
