package blevedemo

import (
	"fmt"

	"github.com/blevesearch/bleve"
	"github.com/bwmarrin/snowflake"
)

type Message struct {
	ID      int64
	Title   string
	Content string
}

func WriteMessage() error {
	indexPath := "example.bleve"
	index, err := bleve.Open(indexPath)
	if err != nil {
		return err
	}

	defer index.Close()

	n, err := snowflake.NewNode(67)
	if err != nil {
		return err
	}

	id := n.Generate()
	m := Message{id.Int64(), "test", "testing content."}

	err = index.Index(id.String(), m)
	if err != nil {
		return err
	}

	fmt.Println("create index successed.")
	return nil
}
