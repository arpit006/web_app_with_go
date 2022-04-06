package factory

import "github.com/gorilla/mux"

var ROUTER_INSTANCE *mux.Router = nil

type RouterSingletonFactory interface {
	GetRouter() *mux.Router
}

type MuxRouter struct {

}

func (mx MuxRouter) GetRouter() *mux.Router {
	if ROUTER_INSTANCE != nil {
		return ROUTER_INSTANCE
	}
	ROUTER_INSTANCE = mux.NewRouter()
	return ROUTER_INSTANCE
}
