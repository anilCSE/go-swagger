package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/go-swagger/go-swagger/examples/generated/models"
)

// CreateUsersWithArrayInputParams contains all the bound params for the create users with array input operation
// typically these are obtained from a http.Request
//
// swagger:parameters createUsersWithArrayInput
type CreateUsersWithArrayInputParams struct {
	/*List of user object
	  In: body
	*/
	Body []*models.User
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateUsersWithArrayInputParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	var body []*models.User
	if err := route.Consumer.Consume(r.Body, &body); err != nil {
		res = append(res, errors.NewParseError("body", "body", "", err))
	} else {
		for _, io := range o.Body {
			if err := io.Validate(route.Formats); err != nil {
				res = append(res, err)
				break
			}
		}

		if len(res) == 0 {
			o.Body = body
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
