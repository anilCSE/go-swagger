package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// AddPetHandlerFunc turns a function with the right signature into a add pet handler
type AddPetHandlerFunc func(AddPetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddPetHandlerFunc) Handle(params AddPetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddPetHandler interface for that can handle valid add pet params
type AddPetHandler interface {
	Handle(AddPetParams, interface{}) middleware.Responder
}

// NewAddPet creates a new http.Handler for the add pet operation
func NewAddPet(ctx *middleware.Context, handler AddPetHandler) *AddPet {
	return &AddPet{Context: ctx, Handler: handler}
}

/*AddPet swagger:route POST /pets pet addPet

Add a new pet to the store

*/
type AddPet struct {
	Context *middleware.Context
	Params  AddPetParams
	Handler AddPetHandler
}

func (o *AddPet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
