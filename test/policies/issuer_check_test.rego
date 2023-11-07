#
#   Copyright Gen Digital Inc. All Rights Reserved.
#
#   SPDX-License-Identifier: Apache-2.0
#

package trustregistry

import future.keywords

test_no_trustlist_allow if {
  allow
    with data.verifiers as [
      {
        "verifier": {
          "id": "verifier1",
          "checks": {
            "credential": {
            }
          }
        }
      }
    ]
    with input as {
      "verifierId": "verifier1",
      "issuerId": "issuer1",
      "credentialType": "Credentialx"
    }
}

test_trustlist_issuer_allow if {
  allow
    with data.verifiers as [
      {
        "verifier": {
          "id": "verifier1",
          "checks": {
            "credential": {
              "issuerTrustList": {
                "issuer1" : {},
                "issuer2": {
                  "credentialTypes": [
                    "credential1",
                    "credential2",
                  ]
                }
              }
            }
          }
        }
      }
    ]
    with input as {
      "verifierId": "verifier1",
      "issuerId": "issuer1",
      "credentialType": "credential3"
    }
}

test_trustlist_issuer_disallow if {
  not allow
    with data.verifiers as [
      {
        "verifier": {
          "id": "verifier1",
          "checks": {
            "credential": {
              "issuerTrustList": {
                "issuer2": {
                  "credentialTypes": [
                    "credential1",
                    "credential2",
                  ]
                }
              }
            }
          }
        }
      }
    ]
    with input as {
      "verifierId": "verifier1",
      "issuerId": "issuer1",
      "credentialType": "credential3"
    }
}

test_trustlist_issuer_and_cred_type_allow if {
  allow
    with data.verifiers as [
      {
        "verifier": {
          "id": "verifier1",
          "checks": {
            "credential": {
              "issuerTrustList": {
                "issuer1" : {},
                "issuer2": {
                  "credentialTypes": [
                    "credential1",
                    "credential2",
                  ]
                }
              }
            }
          }
        }
      }
    ]
    with input as {
      "verifierId": "verifier1",
      "issuerId": "issuer2",
      "credentialType": "credential1"
    }
}

test_trustlist_issuer_and_cred_type_disallow if {
  not allow
    with data.verifiers as [
      {
        "verifier": {
          "id": "verifier1",
          "checks": {
            "credential": {
             "issuerTrustList": {
                "issuer1" : {},
                "issuer2": {
                  "credentialTypes": [
                    "credential1",
                    "credential2",
                  ]
                }
              }
            }
          }
        }
      }
    ]
    with input as {
      "verifierId": "verifier1",
      "issuerId": "issuer2",
      "credentialType": "credential3"
    }
}
