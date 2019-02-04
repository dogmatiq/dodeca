package config

import (
	"io"
	"io/ioutil"
	"os"
	"sync"
)

// tempfile is a ref-counted temporary file that is deleted when there are no
// more references to it.
type tempfile struct {
	m    sync.Mutex
	path string
	refs uint64
}

// addRef increases the ref count for the temporary file used by Path(). The
// temporary file is created the first time addRef() is called, and populated by
// calling fn().
func (t *tempfile) addRef(
	fn func(io.Writer) error,
) (string, error) {
	t.m.Lock()
	defer t.m.Unlock()

	if t.refs == 0 {
		if err := t.create(fn); err != nil {
			return "", err
		}
	}

	t.refs++

	return t.path, nil
}

// decRef decreases the ref count for the temporary file used by Path().
// The temporary file is deleted when decRef() has been called the same number
// of times as addRef().
func (t *tempfile) decRef() error {
	t.m.Lock()
	defer t.m.Unlock()

	t.refs--

	if t.refs == 0 {
		return os.Remove(t.path)
	}

	return nil
}

// create writes a temporary file containing the content of the string to be
// used by Path().
func (t *tempfile) create(fn func(io.Writer) error) error {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	defer f.Close()

	t.path = f.Name()

	if err := fn(f); err != nil {
		os.Remove(t.path)
		return err
	}

	if err := f.Sync(); err != nil {
		os.Remove(t.path)
		return err
	}

	return nil
}
