# Go Spinner 

## Description

Go Spinner is a simple CLI progress spinner for Go. It's easy to use in any Go application where you need a progress indicator for long-running tasks.

The spinner animation is a simple textual representation of a progress indicator in a Command Line Interface (CLI). The spinner itself cycles through a sequence of characters that create the illusion of a spinning effect when printed one after the other in the same position on the command line. 

    Working.... |
    Working.... /
    Working.... -
    Working.... \

## Features

- Simple API: just `Start` and `Stop` methods.
- Configurable spinner characters.
- Works concurrently with other tasks.

## Getting Started

You can add Spinner to your Go projects with the `go get` command:
## Usage

```go
package main

import (
    "github.com/ajanes780/goSpinner"
    "time"
)

func main() {
    s := spinner.New()
    s.Start()
    time.Sleep(3 * time.Second) // replace with your long running code
    s.Stop()
}
```

## License



MIT License.

