package runtime

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Loadable represents a reader for the source code.
type Loadable interface {
	io.RuneScanner
	io.Closer
}

type nopCloser struct {
	io.RuneScanner
}

func (n *nopCloser) Close() error {
	return nil
}

type fileCloser struct {
	*bufio.Reader
	f *os.File
}

func (f *fileCloser) Close() error {
	return f.f.Close()
}

// Loader represents something that can load source code.
type Loader interface {
	Load(path string) (Loadable, error)
}

// StringLoader is a Loader backed by an in-memory map. It is primarily useful
// for testing.
type StringLoader struct {
	m map[string]string
}

// Load returns the data at path or error if not found.
func (s *StringLoader) Load(path string) (Loadable, error) {
	data, ok := s.m[path]
	if !ok {
		return nil, fmt.Errorf("unknown import: %s", path)
	}
	return &nopCloser{strings.NewReader(data)}, nil
}

// FileLoader is a Loader backed by the file system.
type FileLoader struct {
	SearchPaths []string
}

// Load searches for the path along each SearchPath and returns the first
// reader found. If none found, returns an error.
func (f *FileLoader) Load(path string) (Loadable, error) {
	for _, searchPath := range f.SearchPaths {
		absPath := filepath.Join(searchPath, path)
		f, err := os.Open(absPath)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return &fileCloser{bufio.NewReader(f), f}, nil
	}
	return nil, fmt.Errorf("unknown import: %s", path)
}
