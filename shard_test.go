package gopherdb


import (
	"testing"
	"os"
	"log"
)
func TestShard(t *testing.T) {
	db_name := "shard"
	defer os.RemoveAll(db_name) // cleanup
	g, err := New(db_name)
	if err != nil {
		log.Fatal(err)
	}
	g.shard(5)
}
