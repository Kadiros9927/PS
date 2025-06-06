package verify

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

const storageFile = "data.json"

var mu sync.Mutex

type Record struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func SaveRecord(rec Record) error {
	mu.Lock()
	defer mu.Unlock()

	records, _ := loadAll()
	records = append(records, rec)

	data, err := json.MarshalIndent(records, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, data, 0644)
}

func loadAll() ([]Record, error) {
	data, err := os.ReadFile(storageFile)
	if err != nil {
		return []Record{}, err
	}

	var records []Record
	_ = json.Unmarshal(data, &records)
	return records, nil
}

func DeleteByHash(hash string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()

	records, _ := loadAll()
	updated := make([]Record, 0, len(records))
	foundEmail := ""
	matched := false

	for _, rec := range records {
		if rec.Hash == hash {
			foundEmail = rec.Email
			matched = true
			continue
		}
		updated = append(updated, rec)
	}

	newData, _ := json.MarshalIndent(updated, "", "  ")
	err := os.WriteFile(storageFile, newData, 0644)
	if err != nil {
		log.Printf("ошибка при записи файла: %v", err)
		return "", false
	}
	return foundEmail, matched
}
