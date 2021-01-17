package task

import (
	"io/ioutil"
	"os"
)

type Store struct {
	path string
}
type File interface {
	read() []byte
	write([]byte)
}

func (s Store) read() []byte {
	jsonFile, err := os.Open(s.path)
	if err != nil {
		return nil
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	return byteValue
}
func (s Store) write(b []byte) {
	_ = ioutil.WriteFile(s.path, b, 0644)
}
