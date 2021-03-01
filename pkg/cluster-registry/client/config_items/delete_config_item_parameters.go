// Code generated by go-swagger; DO NOT EDIT.

package config_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteConfigItemParams creates a new DeleteConfigItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteConfigItemParams() *DeleteConfigItemParams {
	return &DeleteConfigItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConfigItemParamsWithTimeout creates a new DeleteConfigItemParams object
// with the ability to set a timeout on a request.
func NewDeleteConfigItemParamsWithTimeout(timeout time.Duration) *DeleteConfigItemParams {
	return &DeleteConfigItemParams{
		timeout: timeout,
	}
}

// NewDeleteConfigItemParamsWithContext creates a new DeleteConfigItemParams object
// with the ability to set a context for a request.
func NewDeleteConfigItemParamsWithContext(ctx context.Context) *DeleteConfigItemParams {
	return &DeleteConfigItemParams{
		Context: ctx,
	}
}

// NewDeleteConfigItemParamsWithHTTPClient creates a new DeleteConfigItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteConfigItemParamsWithHTTPClient(client *http.Client) *DeleteConfigItemParams {
	return &DeleteConfigItemParams{
		HTTPClient: client,
	}
}

/* DeleteConfigItemParams contains all the parameters to send to the API endpoint
   for the delete config item operation.

   Typically these are written to a http.Request.
*/
type DeleteConfigItemParams struct {

	/* ClusterID.

	   ID of the cluster.
	*/
	ClusterID string

	/* ConfigKey.

	   Key for the config value.
	*/
	ConfigKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete config item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConfigItemParams) WithDefaults() *DeleteConfigItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete config item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConfigItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete config item params
func (o *DeleteConfigItemParams) WithTimeout(timeout time.Duration) *DeleteConfigItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete config item params
func (o *DeleteConfigItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete config item params
func (o *DeleteConfigItemParams) WithContext(ctx context.Context) *DeleteConfigItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete config item params
func (o *DeleteConfigItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete config item params
func (o *DeleteConfigItemParams) WithHTTPClient(client *http.Client) *DeleteConfigItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete config item params
func (o *DeleteConfigItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete config item params
func (o *DeleteConfigItemParams) WithClusterID(clusterID string) *DeleteConfigItemParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete config item params
func (o *DeleteConfigItemParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithConfigKey adds the configKey to the delete config item params
func (o *DeleteConfigItemParams) WithConfigKey(configKey string) *DeleteConfigItemParams {
	o.SetConfigKey(configKey)
	return o
}

// SetConfigKey adds the configKey to the delete config item params
func (o *DeleteConfigItemParams) SetConfigKey(configKey string) {
	o.ConfigKey = configKey
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConfigItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param config_key
	if err := r.SetPathParam("config_key", o.ConfigKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
