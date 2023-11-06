/*
Copyright Gen Digital Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"

	"github.com/open-policy-agent/opa/sdk"
	sdktest "github.com/open-policy-agent/opa/sdk/test"
)

//go:embed "policy/trustregistry/issuer_check.rego"
var issuerCheckPolicy string

//go:embed "policy/trustregistry/profiles.rego"
var profilesPolicy string

func main() {
	ctx := context.Background()

	// create a mock HTTP bundle server
	server, err := sdktest.NewServer(sdktest.MockBundle("/bundles/bundle.tar.gz",
		map[string]string{
			"issuer_check.rego": issuerCheckPolicy,
			"profiles.rego":     profilesPolicy,
		},
	))
	if err != nil {
		panic(err)
	}

	defer server.Stop()

	// provide the OPA configuration which specifies
	// fetching policy bundles from the mock server
	// and logging decisions locally to the console
	config := []byte(fmt.Sprintf(`{
		"services": {
			"test": {
				"url": %q
			}
		},
		"bundles": {
			"test": {
				"resource": "/bundles/bundle.tar.gz"
			}
		},
		"decision_logs": {
			"console": true
		}
	}`, server.URL()))

	ready := make(chan struct{})

	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()

	// create an instance of the OPA object
	opa, err := sdk.New(ctx, sdk.Options{
		ID:     "opa-test-1",
		Config: bytes.NewReader(config),
		Ready:  ready,
	})
	if err != nil {
		panic(err)
	}

	defer opa.Stop(ctx)

	<-ready

	input := map[string]interface{}{
		"verifierId":     "v_myprofile_jwt_whitelist",
		"issuerId":       "bank_issuer_sdjwt_v5",
		"credentialType": "CrudeProductCredential",
	}

	result, err := opa.Decision(ctx, sdk.DecisionOptions{
		Path:  "/trustregistry/allow",
		Input: input},
	)
	if err != nil {
		panic(err)
	}

	allowed, ok := result.Result.(bool)
	if !ok {
		panic("expecting result to be bool")
	}

	fmt.Printf("Result for issuer check: %t\n", allowed)

	input = map[string]interface{}{
		"verifierId":     "v_myprofile_jwt_whitelist",
		"issuerId":       "bank_issuer_sdjwt_v5",
		"credentialType": "NotWhiteListed",
	}

	result, err = opa.Decision(ctx, sdk.DecisionOptions{
		Path:  "/trustregistry/allow",
		Input: input},
	)
	if err != nil {
		panic(err)
	}

	allowed, ok = result.Result.(bool)
	if !ok {
		panic("expecting result to be bool")
	}

	fmt.Printf("Result for issuer check: %t\n", allowed)
}
