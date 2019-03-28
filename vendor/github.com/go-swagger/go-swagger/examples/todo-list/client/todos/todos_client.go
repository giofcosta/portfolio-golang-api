// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new todos API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for todos API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddOne add one API
*/
func (a *Client) AddOne(params *AddOneParams, authInfo runtime.ClientAuthInfoWriter) (*AddOneCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOne",
		Method:             "POST",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/io.swagger.examples.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.examples.todo-list.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddOneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddOneCreated), nil

}

/*
DestroyOne destroy one API
*/
func (a *Client) DestroyOne(params *DestroyOneParams, authInfo runtime.ClientAuthInfoWriter) (*DestroyOneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDestroyOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "destroyOne",
		Method:             "DELETE",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/io.swagger.examples.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.examples.todo-list.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DestroyOneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DestroyOneNoContent), nil

}

/*
Find find API
*/
func (a *Client) Find(params *FindParams, authInfo runtime.ClientAuthInfoWriter) (*FindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "find",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/io.swagger.examples.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.examples.todo-list.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindOK), nil

}

/*
UpdateOne update one API
*/
func (a *Client) UpdateOne(params *UpdateOneParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOne",
		Method:             "PUT",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/io.swagger.examples.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.examples.todo-list.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOneOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
