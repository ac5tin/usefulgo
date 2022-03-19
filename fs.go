package utils

import (
	"io/ioutil"
	"os"
)

// FS - filesystem
type FS struct{}

// NewFS - returns new FS
func NewFS() FS {
	return FS{}
}

// Write - write to file
func (fs FS) Write(data *[]byte, path string) error {
	err := os.Truncate(path, 0)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	f.Write(*data)
	return nil
}

// Read - read file
func (fs FS) Read(path string) (*[]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FileExist returns bool value of whether file exists
func (fs FS) FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// CreateFile creates empty file
func (fs FS) CreateFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
