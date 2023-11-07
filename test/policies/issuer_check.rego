package trustregistry

import future.keywords.in
import data.verifiers

default allow = false

# Allow if no trust list for verifier
allow {
    some verifier in data.verifiers
    verifier.verifier.id == input.verifierId
    not verifier.verifier.checks.credential.issuerTrustList
}

# Allow if issuer contained in trust list for verifier
# for any credential type.
allow {
    some verifier in data.verifiers
    verifier.verifier.id == input.verifierId
    count(verifier.verifier.checks.credential.issuerTrustList[input.issuerId]) == 0
}


# Allow if issuer contained in trust list for verifier
# for specific credential types.
allow {
    some verifier in data.verifiers
    verifier.verifier.id == input.verifierId
    some cred_type in verifier.verifier.checks.credential.issuerTrustList[input.issuerId].credentialTypes
	input.credentialType == cred_type
}
