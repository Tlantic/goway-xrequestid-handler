package goway_xrequestid_handler

import (
	"net/http"
	"github.com/twinj/uuid"
	"github.com/andrepinto/goway/router"
	. "github.com/andrepinto/goway/handlers"
)

//noinspection GoUnusedConst
const (
	NAME = "REQUEST_ID"
	DEFAULT_HEADER_KEY = "X-Request-Id"
)


type XRequestIDHandler struct {
	HeaderKey	string
}
//noinspection GoUnusedExportedFunction
func NewXRequestIDHandler( key string ) (*XRequestIDHandler) {

	if (key == "") {
		key = DEFAULT_HEADER_KEY
	}
	return &XRequestIDHandler{
		HeaderKey: key,
	}
}

func (handler *XRequestIDHandler) Handle(route *router.Route, req *http.Request) (*HandlerError){
	req.Header.Add(handler.HeaderKey, uuid.NewV4().String())
	return nil
}