package main

import (
	"context"
	"flag"
	"log"
	"math"
	"time"

	"github.com/MrGossett/plugger/shared"
)

func main() {
	// get the filepath of the .so from a flag
	var filepath string
	flag.StringVar(&filepath, "plugin", "", "path to .so that provides a Solver")
	flag.Parse()

	if filepath == "" {
		log.Fatal("plugin path must be provided")
	}

	// setup our timeout
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// initialize our problem instance
	n := &node{
		worker: -1, job: -1,
		assignments: make([]bool, len(costs)),
	}

	// do the thing
	solution, err := shared.Solve(ctx, filepath, n)
	if err != nil {
		log.Fatal(err)
	}
	if solution == nil {
		log.Fatal("solution is nil")
	}
	n = solution.(*node)
	cost := n.Cost()

	if cost != math.MaxInt64 {
		log.Printf("total cost is %d\n", cost)
		for i := n; i.parent != nil; i = i.parent {
			log.Printf("assign worker %d to job %d\n", i.worker, i.job)
		}
	}
}

var _ shared.Node = &node{}

type node struct {
	parent      *node
	worker, job int
	cost        int
	assignments []bool
}

func (n *node) Branch() []shared.Node {
	children := []shared.Node{}
	for i, assigned := range n.assignments {
		if assigned {
			continue
		}
		child := &node{
			parent:      n,
			worker:      n.worker + 1,
			job:         i,
			assignments: make([]bool, len(n.assignments)),
		}
		copy(child.assignments, n.assignments)
		child.assignments[i] = true
		children = append(children, child)
	}
	return children
}

func (n *node) Cost() int {
	if n.cost == 0 {
		n.cost = cost(n.worker, n.job)
		if n.parent != nil {
			n.cost += n.parent.cost
		}
	}
	return n.cost
}

func (n *node) IsSolution() bool {
	for _, b := range n.assignments {
		if !b {
			return false
		}
	}
	return true
}

// cols: workers, rows: jobs
var costs = [][]int{
	{2, 7, 4, 5, 7, 6, 6, 4, 9, 7},
	{8, 5, 2, 1, 2, 1, 8, 2, 2, 4},
	{7, 5, 8, 9, 4, 8, 6, 9, 7, 4},
	{9, 9, 6, 3, 2, 1, 9, 6, 7, 1},
	{9, 7, 1, 8, 2, 3, 8, 3, 6, 9},
	{7, 2, 1, 5, 4, 1, 8, 3, 4, 5},
	{3, 2, 5, 4, 2, 7, 2, 5, 9, 1},
	{4, 1, 3, 2, 1, 4, 5, 7, 6, 5},
	{7, 8, 9, 1, 8, 9, 3, 1, 7, 7},
	{9, 6, 3, 3, 9, 3, 5, 8, 8, 2},
}

func cost(worker, job int) int {
	if job >= len(costs) || worker >= len(costs[0]) {
		return math.MaxInt64
	}
	return costs[job][worker]
}
