package configs

import "github.com/meilisearch/meilisearch-go"

var MSClient *meilisearch.Client

func MeiliSearchConfig() {
	MSClient = meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://localhost:7700",
	})
}
