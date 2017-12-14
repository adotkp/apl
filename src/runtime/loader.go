package runtime

import (
    "os"
)

type Loader interface {
  Exists(path string) bool
  Open(path string) (io.RuneScanner, error)
}

type StringLoader struct {
    m map[string]string
}

func (s *StringLoader) Add(path, data string) error {
    if _, ok := s.m[path]; ok {
        return errors.New("string loader got conflicting writes")
    }
    s.m[path] = data
}

func (s *StringLoader) Exists(path string) bool {
    _, ok := s.m[path]
    return ok
}

func (s *StringLoader) Open(path string) (io.RuneScanner, error) {
   data, ok := s.m[path]
   if !ok {
    return nil, fmt.Errorf()
   }
}

type FileLoader interface {
}
