package solution

import "io"

type Solver interface {
	Solve(input io.Reader) (any, any, error)
}

type SolutionFactory func() Solver
