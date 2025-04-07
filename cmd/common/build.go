package common

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"sync"
)

// Sentinel errors
const (
	ErrNotFound     = VersionError("not found")
	ErrNotSupported = VersionError("not supported")
)

type VersionError string

func (e VersionError) Error() string {
	return string(e)
}

func UserAgent() string {
	exePath, err := os.Executable()
	if err != nil {
		exePath = os.Args[0]
	}

	exeBase := filepath.Base(exePath)
	return fmt.Sprintf("%s/%s", exeBase, MustBuildVersion())

}

const BuildUnknown = "unknown"

type Build struct {
	*debug.BuildInfo
}

func (b *Build) Version() string {
	if b == nil {
		return BuildUnknown
	}
	return b.Main.Version
}

var cachedBuildInfo struct {
	once sync.Once
	bi   *debug.BuildInfo
	ok   bool
}

func NewBuild() (Build, error) {
	cachedBuildInfo.once.Do(func() {
		bi, ok := debug.ReadBuildInfo()
		cachedBuildInfo.bi = bi
		cachedBuildInfo.ok = ok
	})
	if !cachedBuildInfo.ok {
		return Build{}, ErrNotSupported
	}
	return Build{BuildInfo: cachedBuildInfo.bi}, nil
}

func MustBuildVersion() string {
	b, err := NewBuild()
	if err != nil {
		return BuildUnknown
	}
	return b.Version()
}
