package solver

type Opts struct {
	Input string
}

type Result struct {
	Output string
}

type Solver interface {
	Solve(opts Opts) (Result, error)
}
