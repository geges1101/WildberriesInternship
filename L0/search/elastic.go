package search

import (
	"context"
	"encoding/json"
	"github.com/geges1101/l0/model"
	"log"

	"github.com/olivere/elastic"
)

type ElasticRepository struct {
	client *elastic.Client
}

func (r *ElasticRepository) InsertMessage(ctx context.Context, data model.Data) error {
	//TODO implement me
	panic("implement me")
}

func (r *ElasticRepository) SearchMessage(ctx context.Context, query string, skip uint64, take uint64) ([]model.Data, error) {
	//TODO implement me
	panic("implement me")
}

func NewElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {
}

func (r *ElasticRepository) InsertMeow(ctx context.Context, data model.Data) error {
	_, err := r.client.Index().
		Index("data").
		Type("message").
		Id(data.ID).
		BodyJson(data).
		Refresh("wait_for").
		Do(ctx)
	return err
}

func (r *ElasticRepository) SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]model.Data, error) {
	result, err := r.client.Search().
		Index("meows").
		Query(
			elastic.NewMultiMatchQuery(query, "body").
				Fuzziness("3").
				PrefixLength(1).
				CutoffFrequency(0.0001),
		).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	meows := []model.Data{}
	for _, hit := range result.Hits.Hits {
		var meow model.Data
		if err = json.Unmarshal(*hit.Source, &meow); err != nil {
			log.Println(err)
		}
		meows = append(meows, meow)
	}
	return meows, nil
}
