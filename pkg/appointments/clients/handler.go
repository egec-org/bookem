package clients

import "fmt"

type handler struct {
	svc ClientService
}

func (h *handler) FetchByUsername() {
	fmt.Println("panic")
}
