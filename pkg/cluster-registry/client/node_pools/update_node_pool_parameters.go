// Code generated by go-swagger; DO NOT EDIT.

package node_pools

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

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// NewUpdateNodePoolParams creates a new UpdateNodePoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNodePoolParams() *UpdateNodePoolParams {
	return &UpdateNodePoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNodePoolParamsWithTimeout creates a new UpdateNodePoolParams object
// with the ability to set a timeout on a request.
func NewUpdateNodePoolParamsWithTimeout(timeout time.Duration) *UpdateNodePoolParams {
	return &UpdateNodePoolParams{
		timeout: timeout,
	}
}

// NewUpdateNodePoolParamsWithContext creates a new UpdateNodePoolParams object
// with the ability to set a context for a request.
func NewUpdateNodePoolParamsWithContext(ctx context.Context) *UpdateNodePoolParams {
	return &UpdateNodePoolParams{
		Context: ctx,
	}
}

// NewUpdateNodePoolParamsWithHTTPClient creates a new UpdateNodePoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNodePoolParamsWithHTTPClient(client *http.Client) *UpdateNodePoolParams {
	return &UpdateNodePoolParams{
		HTTPClient: client,
	}
}

/* UpdateNodePoolParams contains all the parameters to send to the API endpoint
   for the update node pool operation.

   Typically these are written to a http.Request.
*/
type UpdateNodePoolParams struct {

	/* ClusterID.

	   ID of the cluster.
	*/
	ClusterID string

	/* NodePool.

	   Node pool to be updated.
	*/
	NodePool *models.NodePoolUpdate

	/* NodePoolName.

	   Name of the node pool.
	*/
	NodePoolName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update node pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNodePoolParams) WithDefaults() *UpdateNodePoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update node pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNodePoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update node pool params
func (o *UpdateNodePoolParams) WithTimeout(timeout time.Duration) *UpdateNodePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update node pool params
func (o *UpdateNodePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update node pool params
func (o *UpdateNodePoolParams) WithContext(ctx context.Context) *UpdateNodePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update node pool params
func (o *UpdateNodePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update node pool params
func (o *UpdateNodePoolParams) WithHTTPClient(client *http.Client) *UpdateNodePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update node pool params
func (o *UpdateNodePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the update node pool params
func (o *UpdateNodePoolParams) WithClusterID(clusterID string) *UpdateNodePoolParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update node pool params
func (o *UpdateNodePoolParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithNodePool adds the nodePool to the update node pool params
func (o *UpdateNodePoolParams) WithNodePool(nodePool *models.NodePoolUpdate) *UpdateNodePoolParams {
	o.SetNodePool(nodePool)
	return o
}

// SetNodePool adds the nodePool to the update node pool params
func (o *UpdateNodePoolParams) SetNodePool(nodePool *models.NodePoolUpdate) {
	o.NodePool = nodePool
}

// WithNodePoolName adds the nodePoolName to the update node pool params
func (o *UpdateNodePoolParams) WithNodePoolName(nodePoolName string) *UpdateNodePoolParams {
	o.SetNodePoolName(nodePoolName)
	return o
}

// SetNodePoolName adds the nodePoolName to the update node pool params
func (o *UpdateNodePoolParams) SetNodePoolName(nodePoolName string) {
	o.NodePoolName = nodePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNodePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}
	if o.NodePool != nil {
		if err := r.SetBodyParam(o.NodePool); err != nil {
			return err
		}
	}

	// path param node_pool_name
	if err := r.SetPathParam("node_pool_name", o.NodePoolName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
