package adjgraph

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	g := CreateGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	log.Println(g)

}
