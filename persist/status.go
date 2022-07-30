package persist

import (
	"encoding/json"

	"github.com/cloudsftp/ResourceBlockerBackend/resource"
	"github.com/dgraph-io/badger"
)

func UpdateStatus(id string, status *resource.ResourceStatus) error {
	return db.Update(func(txn *badger.Txn) error {
		key := statusKey(id)
		val, err := json.Marshal(status)
		if err != nil {
			return err
		}

		return txn.Set(key, val)
	})
}

func GetStatus(id string) (*resource.ResourceStatus, error) {
	status := &resource.ResourceStatus{}

	err := db.View(func(txn *badger.Txn) error {
		key := statusKey(id)
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			return json.Unmarshal(val, status)
		})
		return err
	})

	return status, err
}

func statusKey(id string) []byte {
	return []byte("status_" + id)
}
