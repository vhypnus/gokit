package graph


type Graph struct {

	//total ver
	tv int

	//total edge
	te int

	adj map[interface{}]*ver

}

type ver struct {

	//value
	v interface {}

	// edge weight
	w float32

	n *node
}

type NULL interface{}

func NewGraph() *Graph{

	var adj = make(map[interface]*ver)
	var g = &Graph{0,0,adj}

	return g
}


func AddEdge(src interface{},weight float32,target interface{}) {

	var s,e ver

	if g.adj[src] == nil {
		s = &Vertex{v}
	}

	if g.adj[target]

	if g.adj[value] == nil {
		
	} else {
		//panic
	}
}