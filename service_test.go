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

	// Async
	// build up 5 goroutine under a clide goroutine to run
	Service().SetWorkerNum(5).Start(false)
	// after 3 second, all goroutine will stop
	time.Sleep(time.Second * 3)

	// Sync
	//or you can run forever to setting isBlock:true
	//Service().SetWorkerNum(5).Start(true)


}
