package gopherdb

import (
	"errors"
	"os"
	"strconv"
	"path/filepath"
	"log"
)

type gopher struct {
	name string
	keys []string
}

func New(filePath string) (gopher, error) {
	if _, err := os.Stat("./" + filePath); err == nil {
		return gopher{}, errors.New("Directory already exists with the same name")
	}
	os.Mkdir(filePath, os.ModePerm)
	return gopher{name: filePath}, nil
}

func Load(filePath string) (gopher, error) {
	if _, err := os.Stat("." + filePath); os.IsNotExist(err) { // directory does not exist
		return gopher{}, errors.New("Database does not exist, consider using the New command to construct a new database")
	}
	_ = os.Chdir(filePath)
	files, err := filepath.Glob("*.hole")
	if err != nil {
		log.Fatal(err)
	}
	//
	return gopher{filePath, files}, nil
}

func (g *gopher) shard(numShards int) error {
	for i:=0;i<numShards;i++ {
		os.Create("./"+g.name+"/shard"+strconv.Itoa(i+1)+".log")
	}
	return nil
}

