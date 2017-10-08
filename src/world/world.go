package world

import (
	"sync"
)

//Request - a request to get information from the world dispatcher
type Request struct {
	Type string       //this will determine what the return will be
	Comm *chan string //This is the channel to use when sending response
}

//EventDispatcher - responds to requests submitted via the communications channel.
type EventDispatcher struct {
	Comm chan Request //chanel for sending requests to the world dispatcher
	Mux  sync.Mutex
}

//Universe - The universal event handler, handles/coordinates time based events.
type Universe struct {
	EventChannel EventDispatcher
	//TODO: out going event stream may be needed.
}

func (u *Universe) listenForEvents() {
	event := u.EventChannel
	for {
		event.Mux.Lock()

		event.Mux.Unlock()
	}
}

//CreateUniverse - Creates the universe in which the actors will live.
func (u *Universe) CreateUniverse() {
	//TODO: create an Event Dispatcher, and make it available globally
	eventDispatcher := new(EventDispatcher)
	eventDispatcher.Comm = make(chan Request)

	u.EventChannel = eventDispatcher
}
