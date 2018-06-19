package gopherdb

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"fmt"
	"os"
)

type gopher struct {
	name string
	keys []string
}

func New(filepath string) gopher {
	os.Mkdir(filepath)
	return gopher{name: filepath}
}

func Load(filepath string) (gopher, err) {
	if _, err := os.Stat("./filepath"); os.IsNotExist(err) { // directory does not exist
		return nil, errors.New("Database does not exist, consider using the New command to construct a new database")
	}
	_ = os.Chdir(filepath)
	files, err := filepath.Glob("*.hole")
	if err != nil {
		log.Fatal(err)
	}
	// 
	return gopher{}
}


