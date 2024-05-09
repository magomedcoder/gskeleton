package test

import (
	"path/filepath"
	"runtime"

	"github.com/magomedcoder/gskeleton/internal/config"
)

func GetConfig() *config.Config {
	_, file, _, _ := runtime.Caller(0)

	paths := []string{filepath.Dir(filepath.Dir(file)), "./configs/gskeleton.yaml"}

	return config.New(filepath.Join(paths...))
}
