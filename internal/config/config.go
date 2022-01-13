package config

import (
	"io/ioutil"

	_ "github.com/candy12t/jarvis/internal/statik"
	"github.com/rakyll/statik/fs"
)

func Setup() error {
	if err := NewTwitter(); err != nil {
		return err
	}

	if err := NewJustwatch(); err != nil {
		return err
	}

	return nil
}

func ReadFile(filename string) ([]byte, error) {
	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	r, err := statikFS.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return contents, nil
}
