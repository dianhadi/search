package elastic

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type Elastic struct {
	client *elasticsearch.Client
}

func New(host string, port int) (*Elastic, error) {
	address := fmt.Sprintf("http://%s:%d", host, port)
	cfg := elasticsearch.Config{
		Addresses: []string{address},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	elastic := Elastic{
		client: es,
	}
	return &elastic, nil
}
