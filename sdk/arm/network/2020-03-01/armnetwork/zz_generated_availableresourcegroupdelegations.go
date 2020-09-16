// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// AvailableResourceGroupDelegationsOperations contains the methods for the AvailableResourceGroupDelegations group.
type AvailableResourceGroupDelegationsOperations interface {
	// List - Gets all of the available subnet delegations for this resource group in this region.
	List(location string, resourceGroupName string) AvailableDelegationsResultPager
}

// AvailableResourceGroupDelegationsClient implements the AvailableResourceGroupDelegationsOperations interface.
// Don't use this type directly, use NewAvailableResourceGroupDelegationsClient() instead.
type AvailableResourceGroupDelegationsClient struct {
	*Client
	subscriptionID string
}

// NewAvailableResourceGroupDelegationsClient creates a new instance of AvailableResourceGroupDelegationsClient with the specified values.
func NewAvailableResourceGroupDelegationsClient(c *Client, subscriptionID string) AvailableResourceGroupDelegationsOperations {
	return &AvailableResourceGroupDelegationsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *AvailableResourceGroupDelegationsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// List - Gets all of the available subnet delegations for this resource group in this region.
func (client *AvailableResourceGroupDelegationsClient) List(location string, resourceGroupName string) AvailableDelegationsResultPager {
	return &availableDelegationsResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, location, resourceGroupName)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *AvailableDelegationsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AvailableDelegationsResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *AvailableResourceGroupDelegationsClient) ListCreateRequest(ctx context.Context, location string, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableDelegations"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
func (client *AvailableResourceGroupDelegationsClient) ListHandleResponse(resp *azcore.Response) (*AvailableDelegationsResultResponse, error) {
	result := AvailableDelegationsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailableDelegationsResult)
}

// ListHandleError handles the List error response.
func (client *AvailableResourceGroupDelegationsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
