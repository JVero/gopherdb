package gopherdb


import (
	"testing"
	"fmt"
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
	fmt.Println(g.name)
	g.shard(5)
}
