// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteConfigTunnelVxlanVxlanIDParams creates a new DeleteConfigTunnelVxlanVxlanIDParams object
//
// There are no default values defined in the spec.
func NewDeleteConfigTunnelVxlanVxlanIDParams() DeleteConfigTunnelVxlanVxlanIDParams {

	return DeleteConfigTunnelVxlanVxlanIDParams{}
}

// DeleteConfigTunnelVxlanVxlanIDParams contains all the bound params for the delete config tunnel vxlan vxlan ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteConfigTunnelVxlanVxlanID
type DeleteConfigTunnelVxlanVxlanIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*vxlan id (24-bit). Allows to remove routes with defined vnid only. Applicable for routes with nexthop_type 'vxlan-tunnel'. Otherwise '400' error will be returned
	  Required: true
	  In: path
	*/
	VxlanID int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteConfigTunnelVxlanVxlanIDParams() beforehand.
func (o *DeleteConfigTunnelVxlanVxlanIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rVxlanID, rhkVxlanID, _ := route.Params.GetOK("vxlanID")
	if err := o.bindVxlanID(rVxlanID, rhkVxlanID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindVxlanID binds and validates parameter VxlanID from path.
func (o *DeleteConfigTunnelVxlanVxlanIDParams) bindVxlanID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("vxlanID", "path", "int32", raw)
	}
	o.VxlanID = value

	return nil
}
