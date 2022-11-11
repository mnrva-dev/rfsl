package rfsl

import (
	"os"
	"sync"
	"time"
)

type writer struct {
	mu         sync.Mutex
	filename   string
	fp         *os.File
	size       int64
	timeformat string
	level      LogLevel
}

var rfsl *writer

func startrfsl(filename string) error {
	r := &writer{
		filename: filename,
	}
	err := r.move()
	if err != nil {
		return err
	}
	rfsl = r
	return nil
}

func (r *writer) write(o []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	i, err := os.Stat(r.filename)
	if err != nil {
		return 0, err
	}
	if i.Size() >= r.size && r.size != 0 {
		r.mu.Unlock()
		err = r.move()
		if err != nil {
			return 0, err
		}
		r.mu.Lock()
	}
	return r.fp.Write(o)
}

func (r *writer) move() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// close current logging file
	if r.fp != nil {
		err := r.fp.Close()
		r.fp = nil
		if err != nil {
			// handle error closing file
			return err
		}
	}

	// if file already exists, rename it
	_, err := os.Stat(r.filename)
	if err == nil {
		err = os.Rename(r.filename, r.filename+"."+time.Now().Format("2006-01-02@15-04-05"))
		if err != nil {
			// handle error renaming file
			return err
		}
	}

	r.fp, err = os.Create(r.filename)
	if err != nil {
		// handle error starting new log file
		return err
	}
	return nil
}
