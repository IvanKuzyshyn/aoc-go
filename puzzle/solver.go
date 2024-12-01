package puzzle

type Opts struct {
	Input string
	Part  int8
}

type Result struct {
	Output string
}

type Solver interface {
	Solve(opts Opts) (Result, error)
}
