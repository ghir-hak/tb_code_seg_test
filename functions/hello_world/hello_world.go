package lib

import (
	"github.com/taubyte/go-sdk/event"
)

//export helloworld
func ping(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	h.Write([]byte("Hello world"))

	return 0
}
