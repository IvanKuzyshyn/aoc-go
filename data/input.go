package data

type PuzzleData struct {
	Day      int16
	Year     int16
	Part     int8
	CacheDir string
}

const (
	cacheDir = "cache"
)

func NewPuzzleData(day, year int16, part int8) PuzzleData {
	return PuzzleData{
		Day:      day,
		Year:     year,
		Part:     part,
		CacheDir: cacheDir,
	}
}

func NewEmptyData() PuzzleData {
	return PuzzleData{
		CacheDir: cacheDir,
	}
}
