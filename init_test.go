package gopherdb

import (
	"log"
	"fmt"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	db_name := "first"
	defer os.Remove(db_name) // cleanup

	gopher, err := New(db_name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gopher.name)
}
