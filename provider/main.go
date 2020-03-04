package main

import (
	"context"
	"math"

	"github.com/MrGossett/plugger/shared"
)

func Solve(ctx context.Context, start shared.Node) (shared.Node, error) {
	var (
		best shared.Node
		cost = math.MaxInt64
		q    = start.Branch()
	)

outer:
	for {
		if len(q) == 0 {
			break outer
		}

		select {
		case <-ctx.Done():
			break outer
		default:
		}

		n := q[0]
		q = q[1:]

		if n.IsSolution() && n.Cost() < cost {
			best = n
			cost = n.Cost()
			continue
		}
		for _, bn := range n.Branch() {
			if bn.Cost() >= cost {
				continue
			}
			q = append(q, bn)
		}
	}

	return best, nil
}
