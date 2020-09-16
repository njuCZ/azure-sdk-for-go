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

// VpnSiteLinksOperations contains the methods for the VpnSiteLinks group.
type VpnSiteLinksOperations interface {
	// Get - Retrieves the details of a VPN site link.
	Get(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string) (*VpnSiteLinkResponse, error)
	// ListByVpnSite - Lists all the vpnSiteLinks in a resource group for a vpn site.
	ListByVpnSite(resourceGroupName string, vpnSiteName string) ListVpnSiteLinksResultPager
}

// VpnSiteLinksClient implements the VpnSiteLinksOperations interface.
// Don't use this type directly, use NewVpnSiteLinksClient() instead.
type VpnSiteLinksClient struct {
	*Client
	subscriptionID string
}

// NewVpnSiteLinksClient creates a new instance of VpnSiteLinksClient with the specified values.
func NewVpnSiteLinksClient(c *Client, subscriptionID string) VpnSiteLinksOperations {
	return &VpnSiteLinksClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *VpnSiteLinksClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Retrieves the details of a VPN site link.
func (client *VpnSiteLinksClient) Get(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string) (*VpnSiteLinkResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteLinkName)
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
func (client *VpnSiteLinksClient) GetCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks/{vpnSiteLinkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteLinkName}", url.PathEscape(vpnSiteLinkName))
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
func (client *VpnSiteLinksClient) GetHandleResponse(resp *azcore.Response) (*VpnSiteLinkResponse, error) {
	result := VpnSiteLinkResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnSiteLink)
}

// GetHandleError handles the Get error response.
func (client *VpnSiteLinksClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByVpnSite - Lists all the vpnSiteLinks in a resource group for a vpn site.
func (client *VpnSiteLinksClient) ListByVpnSite(resourceGroupName string, vpnSiteName string) ListVpnSiteLinksResultPager {
	return &listVpnSiteLinksResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByVpnSiteCreateRequest(ctx, resourceGroupName, vpnSiteName)
		},
		responder: client.ListByVpnSiteHandleResponse,
		errorer:   client.ListByVpnSiteHandleError,
		advancer: func(ctx context.Context, resp *ListVpnSiteLinksResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVpnSiteLinksResult.NextLink)
		},
	}
}

// ListByVpnSiteCreateRequest creates the ListByVpnSite request.
func (client *VpnSiteLinksClient) ListByVpnSiteCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
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

// ListByVpnSiteHandleResponse handles the ListByVpnSite response.
func (client *VpnSiteLinksClient) ListByVpnSiteHandleResponse(resp *azcore.Response) (*ListVpnSiteLinksResultResponse, error) {
	result := ListVpnSiteLinksResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnSiteLinksResult)
}

// ListByVpnSiteHandleError handles the ListByVpnSite error response.
func (client *VpnSiteLinksClient) ListByVpnSiteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
