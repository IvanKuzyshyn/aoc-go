package api

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type Cache struct {
	Dir   string
	Input Input
}

func (c *Cache) Write(data []byte) error {
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

func (c *Cache) Read() ([]byte, error) {
	_, cacheFilePath := c.buildCacheFilePath()

	return os.ReadFile(cacheFilePath)
}

func (c *Cache) Clean() error {
	return filepath.Walk(c.Dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == c.Dir {
			return nil
		}
		if info.Name() == ".gitkeep" {
			return nil
		}
		return os.RemoveAll(path)
	})
}

func (c *Cache) buildCacheFilePath() (string, string) {
	folder := path.Join(c.Dir, fmt.Sprintf("%d", c.Input.Year))
	fullPath := path.Join(folder, fmt.Sprintf("%d.txt", c.Input.Day))
	return folder, fullPath
}
