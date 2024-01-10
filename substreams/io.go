package substreams

import (
	"io"
	"os"
)

func readAll(r io.Reader, allocSize int) ([]byte, error) {
	if allocSize <= 0 {
		allocSize = 1024
	}

	b := make([]byte, allocSize, allocSize)
	_, err := r.Read(b)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return b, err
	}

	return b, nil
}

func WriteOutput(data []byte) (int, error) {
	err := os.WriteFile("/sys/substreams/output", data, 0644)
	return len(data), err
}

func ReadInput(allocSize int) ([]byte, error) {
	return readAll(os.Stdin, allocSize)
}

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

type FileWriter interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

type FileReader interface {
	ReadFile(filename string) ([]byte, error)
}

type FileReadWriter interface {
	FileWriter
	FileReader
}

type OSFileReadWriter struct{}

func (r *OSFileReadWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

func (r *OSFileReadWriter) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
