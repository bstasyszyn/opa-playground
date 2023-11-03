/*
   Copyright SecureKey Technologies Inc.

   This file contains software code that is the intellectual property of SecureKey.
   SecureKey reserves all rights in the code and you may not use it without
	 written permission from SecureKey.
*/

package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/open-policy-agent/opa/sdk"
	sdktest "github.com/open-policy-agent/opa/sdk/test"
)

//go:embed "policy/trustregistry/issuer_check.rego"
var issuerCheckPolicy string

//go:embed "policy/trustregistry/profiles.rego"
var profilesPolicy string

//go:embed "data/check_issuer_credtype_allow.json"
var checkIssuerCredTypeAllowInput []byte

//go:embed "data/check_issuer_credtype_disallow.json"
var checkIssuerCredTypeDisallowInput []byte

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

	var input map[string]interface{}
	err = json.Unmarshal(checkIssuerCredTypeAllowInput, &input)
	if err != nil {
		panic(err)
	}

	result, err := opa.Decision(ctx, sdk.DecisionOptions{
		Path:  "/trustregistry/allow",
		Input: input},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result for issuer check: %v\n", result.Result)

	var input2 map[string]interface{}
	err = json.Unmarshal(checkIssuerCredTypeDisallowInput, &input2)
	if err != nil {
		panic(err)
	}

	result, err = opa.Decision(ctx, sdk.DecisionOptions{
		Path:  "/trustregistry/allow",
		Input: input2},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result for issuer check: %v\n", result.Result)
}
