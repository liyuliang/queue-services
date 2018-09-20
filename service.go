package services

import (
	"runtime"
	"sync"
	"time"
	"log"
)

var workerNum int = 5

var isLog bool = true

type service struct {
}

var _service *service

func Service() *service {

	runtime.GOMAXPROCS(runtime.NumCPU())

	if _service == nil {
		_service = new(service)
	}

	return _service
}

func (s *service) setIsLog(is bool) *service {
	isLog = is
	return s
}

func (s *service) SetWorkerNum(num int) *service {
	workerNum = num
	return s
}

func (s *service) Start(isBlock bool) {

	if isBlock {
		wg := sync.WaitGroup{}
		wg.Add(1)

		run()

		wg.Wait()

	} else {

		run()
	}
}

func run() {

	for taskName, multiProcessTask := range multiProcessTasks {

		multiProcessRun(taskName, multiProcessTask)
	}

	for taskName, singleProcessTask := range singleProcessTasks {

		singleProcessRun(taskName, singleProcessTask)
	}
}

func multiProcessRun(taskName string, method taskMethod) {

	go func(name string) {

		wg := sync.WaitGroup{}
		wg.Add(workerNum)

		for i := 0; i < workerNum; i++ {

			go func(workerNum int) {
				if isLog {
					log.Printf("Task \"%s\" is running, current worker is %d, ", name, workerNum)
				}
				for {
					sleepSecond := 1

					err := method(workerNum)
					if err != nil {
						log.Printf("but it get error :%s", err.Error())

						sleepSecond = 3
					}

					time.Sleep(time.Second * time.Duration(sleepSecond))
				}
			}(i)
		}
		wg.Wait()
	}(taskName)
}

func singleProcessRun(taskName string, method taskMethod) {

	go func(name string) {

		for {
			sleepSecond := 1
			if isLog {
				log.Printf("Task \"%s\" is running, ", name)
			}
			err := method(0)
			if err != nil {
				log.Printf("but it get error :%s", err.Error())

				sleepSecond = 3
			}

			time.Sleep(time.Second * time.Duration(sleepSecond))
		}
	}(taskName)
}
