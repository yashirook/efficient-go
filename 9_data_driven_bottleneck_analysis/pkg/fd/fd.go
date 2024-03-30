package fd

import (
	"os"
	"runtime/pprof"
)

var fdProfile = pprof.NewProfile("fs.inuse")

type File struct {
	*os.File
}

func Open(name string) (*File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	fdProfile.Add(f, 2)
	return &File{File: f}, nil
}

func (f *File) Close() error {
	defer fdProfile.Remove(f.File)
	return f.File.Close()
}

func Write(profileOutPath string) error {
	out, err := os.Create(profileOutPath)
	if err != nil {
		return err
	}
	if err := fdProfile.WriteTo(out, 0); err != nil {
		_ = out.Close()
		return err
	}
	return out.Close()
}
