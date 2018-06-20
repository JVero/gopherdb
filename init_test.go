package gopherdb

import (
	"log"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	db_name := "first"
	defer os.RemoveAll(db_name) // cleanup
	g, err := New(db_name)
	if err != nil {
		log.Fatal(err)
	}
	g.shard(1)
}
