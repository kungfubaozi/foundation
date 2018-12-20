package main

import (
	"zskparker.com/foundation/base/logger/cmd/loggercli"
	"zskparker.com/foundation/base/logger/cmd/loggersvc"
)

func main() {
	go func() {
		loggersvc.StartService()
	}()

	loggercli.StartClient()
}
