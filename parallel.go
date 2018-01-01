package parallel

import (
	"sync"
)

// Func is the function to run concurrently.
type Func func() error

// Run calls the passed functions in a goroutine, returns a chan of errors.
func Run(functions ...Func) chan error {
	total := len(functions)
	errs := make(chan error, total)

	var wg sync.WaitGroup
	wg.Add(total)

	go func(errs chan error) {
		wg.Wait()
		close(errs)
	}(errs)

	for _, fn := range functions {
		go func(fn Func, errs chan error) {
			defer wg.Done()
			errs <- fn()
		}(fn, errs)
	}

	return errs
}

// RunLimit calls the passed functions in a goroutine, limiting the number of goroutines running at the same time,
// returns a chan of errors.
func RunLimit(concurrency int, functions ...Func) chan error {
	total := len(functions)

	if concurrency <= 0 {
		concurrency = 1
	}

	if concurrency > total {
		concurrency = total
	}

	var wg sync.WaitGroup
	wg.Add(total)

	errs := make(chan error, total)
	go func(errs chan error) {
		wg.Wait()
		close(errs)
	}(errs)

	sem := make(chan struct{}, concurrency)
	defer func(sem chan<- struct{}) { close(sem) }(sem)

	for _, fn := range functions {
		go func(fn Func, sem <-chan struct{}, errs chan error) {
			defer wg.Done()
			defer func(sem <-chan struct{}) { <-sem }(sem)

			errs <- fn()
		}(fn, sem, errs)
	}

	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}

	return errs
}
