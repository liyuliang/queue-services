package services

import (
	"runtime"
	"sync"
	"time"
	"gitee.com/liyuliang/utils/format"
)

var workerNum int = 5

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

func (s *service) SetWorkerNum(num int) *service {
	workerNum = num
	return s
}

func (s *service) Start() {

	go run()
}

func run() {

	for _, multiProcessTask := range multiProcessTasks {

		multiProcessRun(multiProcessTask)
	}

	for _, singleProcessTask := range singleProcessTasks {

		singleProcessRun(singleProcessTask)
	}
}

func multiProcessRun(method taskMethod) {

	go func() {

		wg := sync.WaitGroup{}
		wg.Add(workerNum)

		for i := 0; i < workerNum; i++ {

			go func(workerNum int) {

				for {
					sleepSecond := 1

					err := method(workerNum)
					if err != nil {

						sleepSecond = 3
					}

					time.Sleep(format.IntToTimeSecond(sleepSecond))
				}
			}(i)
		}
		wg.Wait()
	}()
}

func singleProcessRun(method taskMethod) {

	go func() {

		for {
			_sleepSecond := 1

			err := method(0)
			if err != nil {

				_sleepSecond = 3
			}

			time.Sleep(format.IntToTimeSecond(_sleepSecond))
		}
	}()
}
