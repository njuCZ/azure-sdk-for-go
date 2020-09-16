// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualNetworkTapsOperations contains the methods for the VirtualNetworkTaps group.
type VirtualNetworkTapsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a Virtual Network Tap.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap) (*VirtualNetworkTapPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualNetworkTapPoller, error)
	// BeginDelete - Deletes the specified virtual network tap.
	BeginDelete(ctx context.Context, resourceGroupName string, tapName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets information about the specified virtual network tap.
	Get(ctx context.Context, resourceGroupName string, tapName string) (*VirtualNetworkTapResponse, error)
	// ListAll - Gets all the VirtualNetworkTaps in a subscription.
	ListAll() VirtualNetworkTapListResultPager
	// ListByResourceGroup - Gets all the VirtualNetworkTaps in a subscription.
	ListByResourceGroup(resourceGroupName string) VirtualNetworkTapListResultPager
	// UpdateTags - Updates an VirtualNetworkTap tags.
	UpdateTags(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject) (*VirtualNetworkTapResponse, error)
}

// VirtualNetworkTapsClient implements the VirtualNetworkTapsOperations interface.
// Don't use this type directly, use NewVirtualNetworkTapsClient() instead.
type VirtualNetworkTapsClient struct {
	*Client
	subscriptionID string
}

// NewVirtualNetworkTapsClient creates a new instance of VirtualNetworkTapsClient with the specified values.
func NewVirtualNetworkTapsClient(c *Client, subscriptionID string) VirtualNetworkTapsOperations {
	return &VirtualNetworkTapsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *VirtualNetworkTapsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdate - Creates or updates a Virtual Network Tap.
func (client *VirtualNetworkTapsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap) (*VirtualNetworkTapPollerResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, tapName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("VirtualNetworkTapsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualNetworkTapPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualNetworkTapResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualNetworkTapsClient) ResumeCreateOrUpdate(token string) (VirtualNetworkTapPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkTapsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkTapPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkTapsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualNetworkTapsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VirtualNetworkTapPollerResponse, error) {
	return &VirtualNetworkTapPollerResponse{RawResponse: resp.Response}, nil
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualNetworkTapsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified virtual network tap.
func (client *VirtualNetworkTapsClient) BeginDelete(ctx context.Context, resourceGroupName string, tapName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, tapName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	result, err := client.DeleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("VirtualNetworkTapsClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualNetworkTapsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkTapsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *VirtualNetworkTapsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, tapName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleResponse handles the Delete response.
func (client *VirtualNetworkTapsClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *VirtualNetworkTapsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets information about the specified virtual network tap.
func (client *VirtualNetworkTapsClient) Get(ctx context.Context, resourceGroupName string, tapName string) (*VirtualNetworkTapResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, tapName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *VirtualNetworkTapsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, tapName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *VirtualNetworkTapsClient) GetHandleResponse(resp *azcore.Response) (*VirtualNetworkTapResponse, error) {
	result := VirtualNetworkTapResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkTap)
}

// GetHandleError handles the Get error response.
func (client *VirtualNetworkTapsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all the VirtualNetworkTaps in a subscription.
func (client *VirtualNetworkTapsClient) ListAll() VirtualNetworkTapListResultPager {
	return &virtualNetworkTapListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAllCreateRequest(ctx)
		},
		responder: client.ListAllHandleResponse,
		errorer:   client.ListAllHandleError,
		advancer: func(ctx context.Context, resp *VirtualNetworkTapListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkTapListResult.NextLink)
		},
	}
}

// ListAllCreateRequest creates the ListAll request.
func (client *VirtualNetworkTapsClient) ListAllCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworkTaps"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListAllHandleResponse handles the ListAll response.
func (client *VirtualNetworkTapsClient) ListAllHandleResponse(resp *azcore.Response) (*VirtualNetworkTapListResultResponse, error) {
	result := VirtualNetworkTapListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkTapListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *VirtualNetworkTapsClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Gets all the VirtualNetworkTaps in a subscription.
func (client *VirtualNetworkTapsClient) ListByResourceGroup(resourceGroupName string) VirtualNetworkTapListResultPager {
	return &virtualNetworkTapListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *VirtualNetworkTapListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkTapListResult.NextLink)
		},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualNetworkTapsClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualNetworkTapsClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*VirtualNetworkTapListResultResponse, error) {
	result := VirtualNetworkTapListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkTapListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VirtualNetworkTapsClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates an VirtualNetworkTap tags.
func (client *VirtualNetworkTapsClient) UpdateTags(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject) (*VirtualNetworkTapResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, tapName, tapParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualNetworkTapsClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(tapParameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualNetworkTapsClient) UpdateTagsHandleResponse(resp *azcore.Response) (*VirtualNetworkTapResponse, error) {
	result := VirtualNetworkTapResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkTap)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *VirtualNetworkTapsClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
