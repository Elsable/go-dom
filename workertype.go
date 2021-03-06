// Code generated DO NOT EDIT
// workertype.go
package dom

import "syscall/js"

type WorkerType string

const (
	WorkerTypeClassic WorkerType = "classic"
	WorkerTypeModule  WorkerType = "module"
)

func JSValueToWorkerType(val js.Value) WorkerType {
	return WorkerType(val.String())
}
