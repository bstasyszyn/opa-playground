/*
   Copyright Gen Digital Inc.

   This file contains software code that is the intellectual property of Gen Digital.
   Gen Digital reserves all rights in the code and you may not use it without
	 written permission from Gen Digital.
*/

package httpreader

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

//go:embed "testdata/policies.json"
var policies []byte

func TestHTTPPolicyReader_Read(t *testing.T) {
	const cfgURL = "https://example.com/policies.json"

	httpClient := NewMockHttpClient(gomock.NewController(t))
	httpClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		require.Equal(t, http.MethodGet, req.Method)
		require.Equal(t, cfgURL, req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(policies)),
		}, nil
	})

	r := New(cfgURL, httpClient)
	require.NotNil(t, r)

	policies, err := r.Read(context.TODO())
	require.NoError(t, err)
	require.Len(t, policies, 2)
	require.Equal(t, policies[0].Name, "Trusted Issuers")
}
