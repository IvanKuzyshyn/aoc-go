package data

type PuzzleConfig struct {
	Day      int16
	Year     int16
	Part     int8
	CacheDir string
}

const (
	cacheDir = "cache"
)

func NewPuzzleConfig(day, year int16, part int8) PuzzleConfig {
	return PuzzleConfig{
		Day:      day,
		Year:     year,
		Part:     part,
		CacheDir: cacheDir,
	}
}

func NewEmptyConfig() PuzzleConfig {
	return PuzzleConfig{
		CacheDir: cacheDir,
	}
}
