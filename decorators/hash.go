package proto

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func removeTempFiles() error {
	return filepath.Walk("/tmp", func(path string, f os.FileInfo, err error) error {
		matched, err := filepath.Match(".hash*", filepath.Base(path))
		if err != nil {
			return err
		}
		if matched {
			if err = os.Remove(path); err != nil {
				return err
			}
		}
		return nil
	})
}

func HashToTmp(hashValue []byte) (string, error) {
	if err := removeTempFiles(); err != nil {
		return "", err
	}
	tempfile, err := ioutil.TempFile("", ".hash")
	if err != nil {
		return "", err
	}
	if _, err = tempfile.Write(hashValue); err != nil {
		return "", err
	}
	stat, err := tempfile.Stat()
	if err != nil {
		return "", err
	}
	if err = tempfile.Close(); err != nil {
		return "", err
	}
	return stat.Name(), nil
}
