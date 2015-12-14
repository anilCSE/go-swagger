package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/strfmt"
)

// FindPetsByStatusParams contains all the bound params for the find pets by status operation
// typically these are obtained from a http.Request
//
// swagger:parameters findPetsByStatus
type FindPetsByStatusParams struct {
	/*Status values that need to be considered for filter
	  In: query
	  Collection Format: multi
	*/
	Status []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FindPetsByStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	qs := httpkit.Values(r.URL.Query())

	qStatus, qhkStatus, _ := qs.GetOK("status")
	if err := o.bindStatus(qStatus, qhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindPetsByStatusParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {

	raw := rawData
	size := len(raw)

	if size == 0 {
		return nil
	}

	ic := raw
	isz := size
	var ir []string
	iValidateElement := func(i int, statusI string) *errors.Validation {

		return nil
	}

	for i := 0; i < isz; i++ {

		if err := iValidateElement(i, ic[i]); err != nil {
			return err
		}
		ir = append(ir, ic[i])
	}

	o.Status = ir

	return nil
}
