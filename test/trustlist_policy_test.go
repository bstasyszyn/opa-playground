/*
   Copyright Gen Digital Inc. All Rights Reserved.

   SPDX-License-Identifier: Apache-2.0
*/

package test

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/stretchr/testify/require"

	"opa/pkg/evaluator"
)

const (
	allowPath = "data.trustregistry.allow"
)

func TestTrustlistPolicy(t *testing.T) {
	data, err := os.ReadFile("./data/profiles.json")
	require.NoError(t, err)

	e, err := evaluator.New(
		evaluator.WithModuleFile("./policies/issuer_check.rego"),
		// evaluator.WithModuleFile("./policies/profiles.rego"),
		evaluator.WithStore(inmem.NewFromReader(bytes.NewBuffer(data))),
	)
	require.NoError(t, err)

	input := map[string]interface{}{
		"verifierId":     "v_myprofile_jwt_whitelist",
		"issuerId":       "bank_issuer_sdjwt_v5",
		"credentialType": "CrudeProductCredential",
	}

	t.Run("allow -> true", func(t *testing.T) {
		allowed, err := e.Evaluate(context.TODO(), allowPath, input)
		require.NoError(t, err)
		require.True(t, allowed)
	})
}
