package main

import (
	"io"
	"log"
	"sync"

	"github.com/yashirook/efficient-go/9_data_driven_bottleneck_analysis/pkg/fd"
)

type TestApp struct {
	files []io.ReadCloser
}

func (a *TestApp) Close() {
	for i, cl := range a.files {
		_ = cl.Close()
		a.files[i] = nil
	}
	a.files = a.files[:0]
}

func (a *TestApp) open(name string) {
	f, _ := fd.Open(name)
	a.files = append(a.files, f)
}

func (a *TestApp) OpenSingleFile(name string) {
	a.open(name)
}

func (a *TestApp) OpenTenFiles(name string) {
	for i := 0; i < 10; i++ {
		a.open(name)
	}
}

func (a *TestApp) Open100FilesConcurrently(name string) {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			a.OpenTenFiles(name)
			wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	a := &TestApp{}
	defer a.Close()

	for i := 0; i < 10; i++ {
		a.OpenTenFiles("/dev/null")
		a.Close()
	}

	f, _ := fd.Open("/dev/null")
	a.files = append(a.files, f)

	a.OpenSingleFile("/dev/null")
	a.OpenTenFiles("/dev/null")
	a.Open100FilesConcurrently("/dev/null")

	if err := fd.Write("fd.pprof"); err != nil {
		log.Fatal(err)
	}
}
