package data

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

func (c *PuzzleConfig) WriteCache(data []byte) error {
	cacheFolder, cacheFilePath := c.buildCacheFilePath()
	var err error
	err = os.MkdirAll(cacheFolder, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(cacheFilePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *PuzzleConfig) ReadCache() ([]byte, error) {
	_, cacheFilePath := c.buildCacheFilePath()

	return os.ReadFile(cacheFilePath)
}

func (c *PuzzleConfig) CleanCache() error {
	return filepath.Walk(c.CacheDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == c.CacheDir {
			return nil
		}
		if info.Name() == ".gitkeep" {
			return nil
		}
		return os.RemoveAll(path)
	})
}

func (c *PuzzleConfig) buildCacheFilePath() (string, string) {
	folder := path.Join(c.CacheDir, fmt.Sprintf("%d", c.Year))
	fullPath := path.Join(folder, fmt.Sprintf("%d.txt", c.Day))
	return folder, fullPath
}
