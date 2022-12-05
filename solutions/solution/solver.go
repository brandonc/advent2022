package solution

import "io"

type Solver interface {
	Solve(input io.Reader) (interface{}, interface{}, error)
}

type SolutionFactory func() Solver
