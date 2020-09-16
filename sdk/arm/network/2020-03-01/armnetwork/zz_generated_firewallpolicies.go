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

// FirewallPoliciesOperations contains the methods for the FirewallPolicies group.
type FirewallPoliciesOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified Firewall Policy.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy) (*FirewallPolicyPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (FirewallPolicyPoller, error)
	// BeginDelete - Deletes the specified Firewall Policy.
	BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified Firewall Policy.
	Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, firewallPoliciesGetOptions *FirewallPoliciesGetOptions) (*FirewallPolicyResponse, error)
	// List - Lists all Firewall Policies in a resource group.
	List(resourceGroupName string) FirewallPolicyListResultPager
	// ListAll - Gets all the Firewall Policies in a subscription.
	ListAll() FirewallPolicyListResultPager
}

// FirewallPoliciesClient implements the FirewallPoliciesOperations interface.
// Don't use this type directly, use NewFirewallPoliciesClient() instead.
type FirewallPoliciesClient struct {
	*Client
	subscriptionID string
}

// NewFirewallPoliciesClient creates a new instance of FirewallPoliciesClient with the specified values.
func NewFirewallPoliciesClient(c *Client, subscriptionID string) FirewallPoliciesOperations {
	return &FirewallPoliciesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *FirewallPoliciesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdate - Creates or updates the specified Firewall Policy.
func (client *FirewallPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy) (*FirewallPolicyPollerResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, parameters)
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
	pt, err := armcore.NewPoller("FirewallPoliciesClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &firewallPolicyPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*FirewallPolicyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *FirewallPoliciesClient) ResumeCreateOrUpdate(token string) (FirewallPolicyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPoliciesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &firewallPolicyPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FirewallPoliciesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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
func (client *FirewallPoliciesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*FirewallPolicyPollerResponse, error) {
	return &FirewallPolicyPollerResponse{RawResponse: resp.Response}, nil
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FirewallPoliciesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified Firewall Policy.
func (client *FirewallPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, firewallPolicyName)
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
	pt, err := armcore.NewPoller("FirewallPoliciesClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *FirewallPoliciesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPoliciesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *FirewallPoliciesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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
func (client *FirewallPoliciesClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *FirewallPoliciesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified Firewall Policy.
func (client *FirewallPoliciesClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, firewallPoliciesGetOptions *FirewallPoliciesGetOptions) (*FirewallPolicyResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, firewallPolicyName, firewallPoliciesGetOptions)
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
func (client *FirewallPoliciesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, firewallPoliciesGetOptions *FirewallPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if firewallPoliciesGetOptions != nil && firewallPoliciesGetOptions.Expand != nil {
		query.Set("$expand", *firewallPoliciesGetOptions.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *FirewallPoliciesClient) GetHandleResponse(resp *azcore.Response) (*FirewallPolicyResponse, error) {
	result := FirewallPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicy)
}

// GetHandleError handles the Get error response.
func (client *FirewallPoliciesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all Firewall Policies in a resource group.
func (client *FirewallPoliciesClient) List(resourceGroupName string) FirewallPolicyListResultPager {
	return &firewallPolicyListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *FirewallPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FirewallPolicyListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *FirewallPoliciesClient) ListCreateRequest(ctx context.Context, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies"
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

// ListHandleResponse handles the List response.
func (client *FirewallPoliciesClient) ListHandleResponse(resp *azcore.Response) (*FirewallPolicyListResultResponse, error) {
	result := FirewallPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyListResult)
}

// ListHandleError handles the List error response.
func (client *FirewallPoliciesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all the Firewall Policies in a subscription.
func (client *FirewallPoliciesClient) ListAll() FirewallPolicyListResultPager {
	return &firewallPolicyListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAllCreateRequest(ctx)
		},
		responder: client.ListAllHandleResponse,
		errorer:   client.ListAllHandleError,
		advancer: func(ctx context.Context, resp *FirewallPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FirewallPolicyListResult.NextLink)
		},
	}
}

// ListAllCreateRequest creates the ListAll request.
func (client *FirewallPoliciesClient) ListAllCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/firewallPolicies"
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
func (client *FirewallPoliciesClient) ListAllHandleResponse(resp *azcore.Response) (*FirewallPolicyListResultResponse, error) {
	result := FirewallPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *FirewallPoliciesClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
