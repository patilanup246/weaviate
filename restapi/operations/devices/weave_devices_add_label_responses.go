package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*WeaveDevicesAddLabelOK Successful response

swagger:response weaveDevicesAddLabelOK
*/
type WeaveDevicesAddLabelOK struct {
}

// NewWeaveDevicesAddLabelOK creates WeaveDevicesAddLabelOK with default headers values
func NewWeaveDevicesAddLabelOK() *WeaveDevicesAddLabelOK {
	return &WeaveDevicesAddLabelOK{}
}

// WriteResponse to the client
func (o *WeaveDevicesAddLabelOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}