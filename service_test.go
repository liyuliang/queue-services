package services

import (
	"testing"
	"time"
)

func TestService(t *testing.T) {

	AddMultiProcessTask("test println hello", func(workerNum int) (err error) {
		println("hello ~")

		return err
	})
	AddMultiProcessTask("test println world", func(workerNum int) (err error) {
		println("world ~")

		return err
	})

	AddSingleProcessTask("test pringln --- ", func(workerNum int) (err error) {
		println("---")

		return err
	})

	// set up 5 goroutine to run
	Service().SetWorkerNum(5).Start()


	// after 3 second, all goroutine will stop
	time.Sleep(time.Second * 3)
}
