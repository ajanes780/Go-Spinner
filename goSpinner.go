package goSpinner

import (
	"fmt"
	"time"
)

var spinChars = `|/-\`

type Spinner struct {
	SpecialChar string
	IsSpinning  bool
}

func New() *Spinner {
	return &Spinner{
		SpecialChar: spinChars,
		IsSpinning:  false,
	}
}

func (s *Spinner) Start() {
	s.IsSpinning = true
	go func() { // Start the spinner in a separate goroutine so it is non-blocking.
		for {
			for _, r := range s.SpecialChar {
				if !s.IsSpinning {
					// Clear the line before printing "Done!".
					fmt.Printf("\r%-*s", 50, "") // assuming line can be 20 chars long.
					fmt.Printf("\rDone!\n")
					return
				}
				fmt.Printf("\rWorking.... %c ", r)
				time.Sleep(115 * time.Millisecond)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.IsSpinning = false
	time.Sleep(1 * time.Second)
}
