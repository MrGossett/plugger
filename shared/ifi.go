package shared

import (
	"context"
	"errors"
	"plugin"
)

type Node interface {
	Branch() []Node
	Cost() int
	IsSolution() bool
}

type Solver func(context.Context, Node) (Node, error)

func load(so string) (Solver, error) {
	plug, err := plugin.Open(so)
	if err != nil {
		return nil, err
	}

	sym, err := plug.Lookup("Solve")
	if err != nil {
		return nil, err
	}

	f, ok := sym.(func(context.Context, Node) (Node, error))
	if !ok {
		return nil, errors.New("symbol is not a Solver")
	}

	return Solver(f), nil
}

func Solve(ctx context.Context, filepath string, instance Node) (Node, error) {
	f, err := load(filepath)
	if err != nil {
		return nil, err
	}

	return f(ctx, instance)
}
