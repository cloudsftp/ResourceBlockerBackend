package dir

import (
	"os"
	"path"
)

func GetStoreDir(name string) string {
	prefix := path.Join(os.Getenv("HOME"), ".resbloccker")
	err := os.MkdirAll(prefix, os.ModeDir|0700)
	if err != nil {
		panic(err) // Should not happen, because $HOME exists always
	}

	return path.Join(prefix, name)
}
