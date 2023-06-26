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
	go func() {
		for {
			for _, r := range s.SpecialChar {
				if !s.IsSpinning {
					// Clear the line before printing "Done!".
					fmt.Printf("\r%-*s", 20, "") // assuming line can be 20 chars long.
					fmt.Printf("\rDone!\n")      // Override the spinner when done.
					return
				}
				fmt.Printf("\rWorking.... %c ", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.IsSpinning = false
}
