/*
   Copyright Gen Digital Inc.

   This file contains software code that is the intellectual property of Gen Digital.
   Gen Digital reserves all rights in the code and you may not use it without
	 written permission from Gen Digital.
*/

package httpreader

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"opa/pkg/policy/api"
)

//go:generate mockgen -destination mocks_test.go -package httpreader -source=reader.go -mock_names httpClient=MockHttpClient

const (
	headerAccept    = "Accept"
	contentTypeJSON = "application/json"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type HTTPPolicyReader struct {
	policyConfigURL string
	httpClient      httpClient
}

func New(policyConfigURL string, httpClient httpClient) *HTTPPolicyReader {
	return &HTTPPolicyReader{
		policyConfigURL: policyConfigURL,
		httpClient:      httpClient,
	}
}

func (r *HTTPPolicyReader) Read(ctx context.Context) ([]api.Policy, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.policyConfigURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create http get request for URL [%s]: %w", r.policyConfigURL, err)
	}

	req.Header.Set(headerAccept, contentTypeJSON)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send get request to URL [%s]: %w", r.policyConfigURL, err)
	}

	defer func() {
		if e := resp.Body.Close(); e != nil {
			fmt.Printf("Failed to close response body: %s", e)
		}
	}()

	policies := &api.Policies{}

	err = json.NewDecoder(resp.Body).Decode(&policies)
	if err != nil {
		return nil, fmt.Errorf("decode policies: %w", err)
	}

	return policies.Policies, nil
}
