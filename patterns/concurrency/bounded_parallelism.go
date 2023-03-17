package concurrency

import (
	"crypto/md5"
	"errors"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"sync"
)

type result2 struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func walkFiles2(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()

	return paths, errc
}

func digester(done <-chan struct{}, paths <-chan string, c chan<- result2) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result2{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

func MD5All2(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles2(done, root)

	c := make(chan result2)

	const numDigesters = 20
	var wg sync.WaitGroup
	wg.Add(numDigesters)

	for i := 0; i < numDigesters; i++ {
		go func() {
			defer wg.Done()
			digester(done, paths, c)
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}
