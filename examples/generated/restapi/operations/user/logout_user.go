package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// LogoutUserHandlerFunc turns a function with the right signature into a logout user handler
type LogoutUserHandlerFunc func() middleware.Responder

// Handle executing the request and returning a response
func (fn LogoutUserHandlerFunc) Handle() middleware.Responder {
	return fn()
}

// LogoutUserHandler interface for that can handle valid logout user params
type LogoutUserHandler interface {
	Handle() middleware.Responder
}

// NewLogoutUser creates a new http.Handler for the logout user operation
func NewLogoutUser(ctx *middleware.Context, handler LogoutUserHandler) *LogoutUser {
	return &LogoutUser{Context: ctx, Handler: handler}
}

/*LogoutUser swagger:route GET /users/logout user logoutUser

Logs out current logged in user session

*/
type LogoutUser struct {
	Context *middleware.Context
	Handler LogoutUserHandler
}

func (o *LogoutUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, nil); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle() // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
