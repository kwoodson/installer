package apimanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// QuotaByCounterKeysClient is the apiManagement Client
type QuotaByCounterKeysClient struct {
	BaseClient
}

// NewQuotaByCounterKeysClient creates an instance of the QuotaByCounterKeysClient client.
func NewQuotaByCounterKeysClient(subscriptionID string) QuotaByCounterKeysClient {
	return NewQuotaByCounterKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQuotaByCounterKeysClientWithBaseURI creates an instance of the QuotaByCounterKeysClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewQuotaByCounterKeysClientWithBaseURI(baseURI string, subscriptionID string) QuotaByCounterKeysClient {
	return QuotaByCounterKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByService lists a collection of current quota counter periods associated with the counter-key configured in the
// policy on the specified service instance. The api does not support paging yet.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// quotaCounterKey - quota counter key identifier.This is the result of expression defined in counter-key
// attribute of the quota-by-key policy.For Example, if you specify counter-key="boo" in the policy, then it’s
// accessible by "boo" counter key. But if it’s defined as counter-key="@("b"+"a")" then it will be accessible
// by "ba" key
func (client QuotaByCounterKeysClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string) (result QuotaCounterCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaByCounterKeysClient.ListByService")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.QuotaByCounterKeysClient", "ListByService", err.Error())
	}

	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, quotaCounterKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.QuotaByCounterKeysClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.QuotaByCounterKeysClient", "ListByService", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.QuotaByCounterKeysClient", "ListByService", resp, "Failure responding to request")
		return
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client QuotaByCounterKeysClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"quotaCounterKey":   autorest.Encode("path", quotaCounterKey),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaByCounterKeysClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client QuotaByCounterKeysClient) ListByServiceResponder(resp *http.Response) (result QuotaCounterCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates all the quota counter values specified with the existing quota counter key to a value in the
// specified service instance. This should be used for reset of the quota counter values.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// quotaCounterKey - quota counter key identifier.This is the result of expression defined in counter-key
// attribute of the quota-by-key policy.For Example, if you specify counter-key="boo" in the policy, then it’s
// accessible by "boo" counter key. But if it’s defined as counter-key="@("b"+"a")" then it will be accessible
// by "ba" key
// parameters - the value of the quota counter to be applied to all quota counter periods.
func (client QuotaByCounterKeysClient) Update(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, parameters QuotaCounterValueUpdateContract) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaByCounterKeysClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.QuotaByCounterKeysClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, quotaCounterKey, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.QuotaByCounterKeysClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.QuotaByCounterKeysClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.QuotaByCounterKeysClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client QuotaByCounterKeysClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, parameters QuotaCounterValueUpdateContract) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"quotaCounterKey":   autorest.Encode("path", quotaCounterKey),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaByCounterKeysClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client QuotaByCounterKeysClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
