package meili

import (
	"log"

	"github.com/meilisearch/meilisearch-go"
)

type MeiliSearchStore struct {
	DB *meilisearch.Client
}

func New() *MeiliSearchStore {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "85d7c22efd3395fd5ed83c546f514585f6fa7e8ff4ed982fdf66c1b1cac24aef",
	})

	return &MeiliSearchStore{
		DB: client,
	}
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