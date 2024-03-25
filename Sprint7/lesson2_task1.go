package main

import (
	"os"
	"sync"
)

var (
	dumpFile *os.File
	once2    sync.Once
)

func Dump(data []byte) error {
	once2.Do(func() {
		fname, err := os.Executable()
		if err == nil {
			dumpFile, err = os.Create(fname + `.dump`)
		}
		if err != nil {
			panic(err)
		}
	})
	_, err := dumpFile.Write(data)
	return err
}
