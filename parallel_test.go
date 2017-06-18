package parallel

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	var err error
	var count int
	fn1 := func() error { return nil }
	fn2 := func() error { return errors.New("BOOM!") }

	errs := Run(fn1, fn2)
	for e := range errs {
		count++
		if e != nil {
			err = e
		}
		if count == 2 {
			break
		}
	}

	if count != 2 {
		t.Errorf("parallel.Run() failed, got '%v', expected '%v'", count, 2)
	}

	if err == nil {
		t.Errorf("parallel.Run() failed, got '%v', expected '%v'", err, nil)
	}
}

func TestRunLimit(t *testing.T) {
	var err error
	var count int
	fn1 := func() error { return nil }
	fn2 := func() error { return errors.New("BOOM!") }

	errs := RunLimit(2, fn1, fn2)
	for e := range errs {
		count++
		if e != nil {
			err = e
		}
		if count == 2 {
			break
		}
	}

	if count != 2 {
		t.Errorf("parallel.Run() failed, got '%v', expected '%v'", count, 2)
	}

	if err == nil {
		t.Errorf("parallel.Run() failed, got '%v', expected '%v'", err, nil)
	}
}
