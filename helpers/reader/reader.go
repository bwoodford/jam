package reader

import "os"

type FileReader interface {
	ReadFile(string) ([]byte, error)
}

type ReaderWrapper struct {
}

func (read *ReaderWrapper) ReadFile(val string) ([]byte, error) {
	return os.ReadFile(val)
}
