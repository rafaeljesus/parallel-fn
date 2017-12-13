package parallel

// Run calls the passed functions in a goroutine, returns a chan of errors.
func Run(functions ...func() error) chan error {
	errs := make(chan error, len(functions))

	for _, fn := range functions {
		go func(fn func() error) {
			errs <- fn()
		}(fn)
	}

	return errs
}

// RunLimit calls the passed functions in a goroutine, limiting the number of goroutines running at the same time,
// returns a chan of errors.
func RunLimit(concurrency int, functions ...func() error) chan error {
	total := len(functions)

	if concurrency <= 0 {
		concurrency = 1
	}

	if concurrency > total {
		concurrency = total
	}

	sem := make(chan struct{}, concurrency)
	errs := make(chan error, total)

	for _, fn := range functions {
		go func(fn func() error, sem <-chan struct{}, errs chan error) {
			defer func() { <-sem }()
			errs <- fn()
		}(fn, sem, errs)
	}

	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}

	return errs
}
