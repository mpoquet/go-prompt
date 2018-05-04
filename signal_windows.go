// +build windows

package prompt

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (p *Prompt) handleSignals(exitCh chan int, winSizeCh chan *WinSize, stop chan struct{}) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(
		sigCh,
		syscall.SIGQUIT,
	)

	for {
		select {
		case <-stop:
			log.Println("[INFO] stop handleSignals")
			return
		case s := <-sigCh:
			switch s {

			case syscall.SIGQUIT: // kill -SIGQUIT XXXX
				log.Println("[SIGNAL] Catch SIGQUIT")
				exitCh <- 0
			}
		}
	}
}
