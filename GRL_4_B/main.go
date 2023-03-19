package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}

type AdjacencyList struct {
	nodes   map[int][]int
	visited []bool
}

func NewAdlist(v_count int) AdjacencyList {
	g := AdjacencyList{
		nodes:   make(map[int][]int, v_count),
		visited: make([]bool, v_count+1), // 1 indexed
	}
	return g
}
func (g AdjacencyList) PrintWithSpace() {
	for i, l := range g.nodes {
		fmt.Printf("i=%d:", i)
		for j, v := range l {
			if j >= 1 {
				fmt.Printf(" ")
			}
			fmt.Printf("j=%d", v)
		}
		fmt.Printf("\n")
	}
}
func (g AdjacencyList) Push(from int, to int) {
	g.nodes[from] = append(g.nodes[from], to)
}
func (g *AdjacencyList) TopoSort(v int) []int {
	order := make([]int, v)
	i := 0
	var dfsPath func(node int)
	dfsPath = func(node int) {
		if g.visited[node] {
			return
		}
		g.visited[node] = true
		for _, neighbor := range g.nodes[node] {
			dfsPath(neighbor)
		}
		order[i] = node
		i++
	}
	for node := range g.nodes {
		if !g.visited[node] {
			dfsPath(node)
		}
	}
	for i := 0; i < len(order)/2; i++ {
		order[i], order[len(order)-i-1] = order[len(order)-i-1], order[i]
	}
	return order
}
func (g *AdjacencyList) BuildDirectedGraph(m int) {
	for i := 0; i < m; i++ {
		a := scanInt()
		b := scanInt()
		g.Push(a, b)
	}
}
func (g *AdjacencyList) InitGraph(v int) {
	for i := 0; i < v; i++ {
		if _, exists := g.nodes[i]; !exists {
			g.nodes[i] = []int{}
		}
	}
}
func main() {
	bufInit()

	v, e := scanInt(), scanInt()
	g := NewAdlist(v)
	g.BuildDirectedGraph(e)
	g.InitGraph(v)
	order := g.TopoSort(v)
	for i := 0; i < v; i++ {
		fmt.Printf("%d\n", order[i])
	}
}

var sc = bufio.NewScanner(os.Stdin)

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInt() int {
	sc.Scan()
	return atoi(sc.Text())
}
func atoi(nStr string) int {
	i, err := strconv.Atoi(nStr)
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}
