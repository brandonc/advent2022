package solution

import "io"

type Solver interface {
	Solve(input io.Reader) (int, int, error)
}

type SolutionFactory func() Solver
