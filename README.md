## Parallel fn

* Run functions in parallel.
* Limit the number of goroutines running at the same time.

## Installation
```bash
go get -u github.com/rafaeljesus/parallel-fn
```

## Usage
### Run
```go
package main

import (
  	"errors"

  	"github.com/rafaeljesus/parallel-fn"
)

func main() {
	timeout := time.After(2 * time.Second)
        fn1 := func() error { return nil }
        fn2 := func() error { return errors.New("BOOM!") }

	for {
		select {
		case err := <-Run(fn1, fn2):
                	// catch errors
		case <-timeout:
      			// handle timeout
		}
	}
}
```

### RunLimit
```go
package main

import (
  	"errors"

  	"github.com/rafaeljesus/parallel-fn"
)

func main() {
	timeout := time.After(2 * time.Second)
  	fn1 := func() error { return nil }
  	fn2 := func() error { return errors.New("BOOM!") }
  	fn3 := func() error { nil }
  	fn4 := func() error { nil }

	for {
		select {
		case err := <-RunLimit(2, fn1, fn2, fn3, fn4):
      			// catch errors
		case <-timeout:
      			// handle timeout
		}
	}
}
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![Build Status](https://circleci.com/gh/rafaeljesus/parallel-fn.svg?style=svg)](https://circleci.com/gh/rafaeljesus/parallel-fn)
[![Go Report Card](https://goreportcard.com/badge/github.com/rafaeljesus/parallel-fn)](https://goreportcard.com/report/github.com/rafaeljesus/parallel-fn)
[![Go Doc](https://godoc.org/github.com/rafaeljesus/parallel-fn?status.svg)](https://godoc.org/github.com/rafaeljesus/parallel-fn)

---

> GitHub [@rafaeljesus](https://github.com/rafaeljesus) &nbsp;&middot;&nbsp;
> Medium [@_jesus_rafael](https://medium.com/@_jesus_rafael) &nbsp;&middot;&nbsp;
> Twitter [@_jesus_rafael](https://twitter.com/_jesus_rafael)
