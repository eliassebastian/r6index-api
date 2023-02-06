package meili

import (
	"log"

	"github.com/eliassebastian/r6index-api/cmd/api/models"
	"github.com/eliassebastian/r6index-api/pkg/utils"
	"github.com/meilisearch/meilisearch-go"
)

type MeiliSearchStore struct {
	DB *meilisearch.Client
}

func New() (*MeiliSearchStore, error) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   utils.GetEnv("MEILI_URL", "http://127.0.0.1:7700"),
		APIKey: utils.GetEnv("MEILI_API_KEY", ""),
	})

	_, err := client.GetVersion()
	if err != nil {
		return nil, err
	}

	return &MeiliSearchStore{
		DB: client,
	}, nil
}

func (m *MeiliSearchStore) Init() error {
	keys, err := m.DB.GetKeys(nil)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(keys)

	return nil
}

func (m *MeiliSearchStore) CreateIndex() error {
	_, a := m.DB.CreateIndex(&meilisearch.IndexConfig{
		Uid:        "uplay",
		PrimaryKey: "profileId",
	})

	_, b := m.DB.CreateIndex(&meilisearch.IndexConfig{
		Uid:        "psn",
		PrimaryKey: "profileId",
	})

	_, c := m.DB.CreateIndex(&meilisearch.IndexConfig{
		Uid:        "xbl",
		PrimaryKey: "profileId",
	})

	if a != nil || b != nil || c != nil {
		return a
	}

	log.Println("successfully created uplay, psn, and xbox indexes")
	return nil
}

func (m *MeiliSearchStore) CreateKey() error {
	_, e := m.DB.CreateKey(&meilisearch.Key{
		Description: "General Search API Key",
		Actions:     []string{"search", "documents.add", "documents.get"},
		Indexes:     []string{"uplay", "xbl", "psn"},
		Name:        "Search API Key",
	})

	resp, b := m.DB.GetKeys(nil)
	log.Println(resp, e, b)
	return nil
}

func (m *MeiliSearchStore) GetKeys() error {
	resp, e := m.DB.GetKeys(nil)
	log.Println(resp)
	return e

}

// Search database based on platform and query provided
func (m *MeiliSearchStore) Search(platform, query string) (*meilisearch.SearchResponse, error) {
	return m.DB.Index(platform).Search(query, &meilisearch.SearchRequest{})
}

// Fetch complete player profile from database
func (m *MeiliSearchStore) FetchPlayer(platform, uuid string, output *models.Player) error {
	return m.DB.Index(platform).GetDocument(uuid, &meilisearch.DocumentQuery{}, output)
}

// Check whether player uuid exists in database
func (m *MeiliSearchStore) FindPlayer(platform, uuid string, output *models.PlayerFound) error {
	return m.DB.Index(platform).GetDocument(uuid, &meilisearch.DocumentQuery{Fields: []string{"profileId"}}, output)
}
