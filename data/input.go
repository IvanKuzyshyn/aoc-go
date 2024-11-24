package data

type PuzzleData struct {
	Day      int16
	Year     int16
	CacheDir string
}

const (
	cacheDir = "cache"
)

func NewPuzzleData(day, year int16) PuzzleData {
	return PuzzleData{
		Day:      day,
		Year:     year,
		CacheDir: cacheDir,
	}
}

func NewEmptyData() PuzzleData {
	return PuzzleData{
		CacheDir: cacheDir,
	}
}
