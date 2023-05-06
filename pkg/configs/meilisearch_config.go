package configs

import (
	"os"

	"github.com/meilisearch/meilisearch-go"
)

var MSClient *meilisearch.Client

func MeiliSearchConfig() {
	MEILISEARCH_HOST := os.Getenv("MEILISEARCH_HOST")

	MSClient = meilisearch.NewClient(meilisearch.ClientConfig{
		Host: MEILISEARCH_HOST,
	})
}
