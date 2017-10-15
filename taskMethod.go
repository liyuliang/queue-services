package services

type taskMethod func(workerNum int) error

var multiProcessTasks map[string]taskMethod
var singleProcessTasks map[string]taskMethod

func init() {
	multiProcessTasks = make(map[string]taskMethod)
	singleProcessTasks = make(map[string]taskMethod)
}

func AddMultiProcessTask(methodName string, method taskMethod) {

	multiProcessTasks[methodName] = method
}

func AddSingleProcessTask(methodName string, method taskMethod) {
	singleProcessTasks[methodName] = method
}
