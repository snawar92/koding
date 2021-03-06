package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJGroupFetchModeratorsWithEmailIDReader is a Reader for the PostRemoteAPIJGroupFetchModeratorsWithEmailID structure.
type PostRemoteAPIJGroupFetchModeratorsWithEmailIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupFetchModeratorsWithEmailIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupFetchModeratorsWithEmailIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupFetchModeratorsWithEmailIDOK creates a PostRemoteAPIJGroupFetchModeratorsWithEmailIDOK with default headers values
func NewPostRemoteAPIJGroupFetchModeratorsWithEmailIDOK() *PostRemoteAPIJGroupFetchModeratorsWithEmailIDOK {
	return &PostRemoteAPIJGroupFetchModeratorsWithEmailIDOK{}
}

/*PostRemoteAPIJGroupFetchModeratorsWithEmailIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJGroupFetchModeratorsWithEmailIDOK struct {
	Payload PostRemoteAPIJGroupFetchModeratorsWithEmailIDOKBody
}

func (o *PostRemoteAPIJGroupFetchModeratorsWithEmailIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.fetchModeratorsWithEmail/{id}][%d] postRemoteApiJGroupFetchModeratorsWithEmailIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupFetchModeratorsWithEmailIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJGroupFetchModeratorsWithEmailIDOKBody post remote API j group fetch moderators with email ID o k body
swagger:model PostRemoteAPIJGroupFetchModeratorsWithEmailIDOKBody
*/
type PostRemoteAPIJGroupFetchModeratorsWithEmailIDOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJGroupFetchModeratorsWithEmailIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO0

	var postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJGroupFetchModeratorsWithEmailIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO0)

	postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupFetchModeratorsWithEmailIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j group fetch moderators with email ID o k body
func (o *PostRemoteAPIJGroupFetchModeratorsWithEmailIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JGroup.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
