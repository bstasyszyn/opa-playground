package trustregistry

profiles = {
  "issuers": [],
  "verifiers": [
    {
      "verifier": {
        "id": "v_myprofile_ldp",
        "version": "v1.0",
        "name": "v_myprofile_ldp",
        "organizationID": "00000000-0000-0000-0000-000000000001",
        "url": "https://test-verifier.com",
        "active": true,
        "webHook": "http://vcs.webhook.example.com:8180",
        "checks": {
          "credential": {
            "format": [
              "ldp"
            ],
            "proof": true,
            "status": true
          },
          "presentation": {
            "format": [
              "ldp"
            ],
            "vcSubject": true,
            "proof": true
          }
        },
        "oidcConfig": {
          "roSigningAlgorithm": "EcdsaSecp256k1Signature2019",
          "keyType": "ECDSASecp256k1DER",
          "didMethod": "ion"
        },
        "presentationDefinitions": [
          {
            "id": "32f54163-no-limit-disclosure-single-field",
            "input_descriptors": [
              {
                "id": "degree",
                "name": "degree",
                "purpose": "We can only hire with bachelor degree.",
                "schema": [
                  {
                    "uri": "https://www.w3.org/2018/credentials#VerifiableCredential"
                  }
                ],
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.degree.type",
                        "$.vc.credentialSubject.degree.type"
                      ],
                      "id": "degree_type_id",
                      "purpose": "We can only hire with bachelor degree.",
                      "filter": {
                        "type": "string",
                        "const": "BachelorDegree"
                      }
                    }
                  ]
                }
              }
            ]
          },
          {
            "id": "3c8b1d9a-limit-disclosure-optional-fields",
            "input_descriptors": [
              {
                "id": "uom",
                "name": "uom",
                "purpose": "Crude oil stream specification.",
                "constraints": {
                  "limit_disclosure": "required",
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.physicalSpecs.uom",
                        "$.vc.credentialSubject.physicalSpecs.uom"
                      ],
                      "id": "unit_of_measure_barrel",
                      "purpose": "We can only use barrel UoM.",
                      "filter": {
                        "type": "string",
                        "const": "barrel"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.physicalSpecs.apiGravity",
                        "$.vc.credentialSubject.physicalSpecs.apiGravity"
                      ],
                      "id": "api_gravity",
                      "purpose": "Min API Gravity.",
                      "filter": {
                        "type": "integer",
                        "minimum": 20
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.category",
                        "$.vc.credentialSubject.category"
                      ],
                      "id": "category",
                      "purpose": "Category.",
                      "optional": true,
                      "filter": {
                        "type": "string"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.supplierAddress",
                        "$.vc.credentialSubject.supplierAddress"
                      ],
                      "id": "supplier_address",
                      "purpose": "Supplier Address.",
                      "optional": true,
                      "filter": {
                        "type": "object"
                      }
                    }
                  ]
                }
              }
            ]
          },
          {
            "id": "lp403pb9-schema-match",
            "input_descriptors": [
              {
                "id": "schema",
                "name": "schema",
                "purpose": "Match credentials using specific schema.",
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$[\"@context\"]"
                      ],
                      "id": "schema_id",
                      "purpose": "Match credentials using specific schema.",
                      "filter": {
                        "type": "array",
                        "contains": {
                          "type": "string",
                          "pattern": "https://trustbloc.github.io/context/vc/examples-crude-product-v1.jsonld"
                        }
                      }
                    }
                  ]
                }
              }
            ]
          },
          {
            "id": "062759b1-no-limit-disclosure-optional-fields",
            "input_descriptors": [
              {
                "id": "lprCategory",
                "name": "lprCategory",
                "purpose": "Permanent Resident Card specification",
                "schema": [
                  {
                    "uri": "https://www.w3.org/2018/credentials#VerifiableCredential"
                  }
                ],
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.lprCategory",
                        "$.vc.credentialSubject.lprCategory"
                      ],
                      "id": "lpr_category_id",
                      "purpose": "Specific LPR category.",
                      "filter": {
                        "type": "string",
                        "const": "C09"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.commuterClassification",
                        "$.vc.credentialSubject.commuterClassification"
                      ],
                      "id": "commuter_classification",
                      "purpose": "Specific commuter classification.",
                      "filter": {
                        "type": "string",
                        "const": "C1"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.registrationCity",
                        "$.vc.credentialSubject.registrationCity"
                      ],
                      "id": "registration_city",
                      "purpose": "Specific registration city.",
                      "optional": true,
                      "filter": {
                        "type": "string",
                        "const": "Albuquerque"
                      }
                    }
                  ]
                }
              }
            ]
          }
        ]
      },
      "createDID": true
    },
    {
      "verifier": {
        "id": "v_myprofile_jwt",
        "version": "v1.0",
        "name": "v_myprofile_jwt",
        "organizationID": "00000000-0000-0000-0000-000000000001",
        "url": "https://test-verifier.com",
        "active": true,
        "webHook": "http://vcs.webhook.example.com:8180",
        "checks": {
          "credential": {
            "format": [
              "jwt"
            ],
            "proof": true,
            "status": true,
            "strict": true
          },
          "presentation": {
            "format": [
              "jwt"
            ],
            "vcSubject": true,
            "proof": true
          }
        },
        "oidcConfig": {
          "roSigningAlgorithm": "EcdsaSecp256k1Signature2019",
          "keyType": "ECDSASecp256k1DER",
          "didMethod": "ion"
        },
        "presentationDefinitions": [
          {
            "id": "32f54163-no-limit-disclosure-single-field",
            "input_descriptors": [
              {
                "id": "degree",
                "name": "degree",
                "purpose": "We can only hire with bachelor degree.",
                "schema": [
                  {
                    "uri": "https://www.w3.org/2018/credentials#VerifiableCredential"
                  }
                ],
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.degree.type",
                        "$.vc.credentialSubject.degree.type"
                      ],
                      "id": "degree_type_id",
                      "purpose": "We can only hire with bachelor degree.",
                      "filter": {
                        "type": "string",
                        "const": "BachelorDegree"
                      }
                    }
                  ]
                }
              }
            ]
          },
          {
            "id": "32f54163-no-limit-disclosure-optional-fields",
            "input_descriptors": [
              {
                "id": "lprCategory",
                "name": "lprCategory",
                "purpose": "Permanent Resident Card specification",
                "schema": [
                  {
                    "uri": "https://www.w3.org/2018/credentials#VerifiableCredential"
                  }
                ],
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.lprCategory",
                        "$.vc.credentialSubject.lprCategory"
                      ],
                      "id": "lpr_category_id",
                      "purpose": "Specific LPR category.",
                      "filter": {
                        "type": "string",
                        "const": "C09"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.commuterClassification",
                        "$.vc.credentialSubject.commuterClassification"
                      ],
                      "id": "commuter_classification",
                      "purpose": "Specific commuter classification.",
                      "filter": {
                        "type": "string",
                        "const": "C1"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.registrationCity",
                        "$.vc.credentialSubject.registrationCity"
                      ],
                      "id": "registration_city",
                      "purpose": "Specific registration city.",
                      "optional": true,
                      "filter": {
                        "type": "string",
                        "const": "Albuquerque"
                      }
                    }
                  ]
                }
              }
            ]
          },
          {
            "id": "3c8b1d9a-limit-disclosure-optional-fields",
            "input_descriptors": [
              {
                "id": "uom",
                "name": "uom",
                "purpose": "Crude oil stream specification.",
                "constraints": {
                  "limit_disclosure": "required",
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.physicalSpecs.uom",
                        "$.vc.credentialSubject.physicalSpecs.uom"
                      ],
                      "id": "unit_of_measure_barrel",
                      "purpose": "We can only use barrel UoM.",
                      "filter": {
                        "type": "string",
                        "const": "barrel"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.physicalSpecs.apiGravity",
                        "$.vc.credentialSubject.physicalSpecs.apiGravity"
                      ],
                      "id": "api_gravity",
                      "purpose": "Min API Gravity.",
                      "filter": {
                        "type": "integer",
                        "minimum": 20
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.category",
                        "$.vc.credentialSubject.category"
                      ],
                      "id": "category",
                      "purpose": "Category.",
                      "optional": true,
                      "filter": {
                        "type": "string"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.supplierAddress",
                        "$.vc.credentialSubject.supplierAddress"
                      ],
                      "id": "supplier_address",
                      "purpose": "Supplier Address.",
                      "optional": true,
                      "filter": {
                        "type": "object"
                      }
                    }
                  ]
                }
              }
            ]
          },
          {
            "id": "lp403pb9-schema-match",
            "input_descriptors": [
              {
                "id": "schema",
                "name": "schema",
                "purpose": "Match credentials using specific schema.",
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$[\"@context\"]"
                      ],
                      "id": "schema_id",
                      "purpose": "Match credentials using specific schema.",
                      "filter": {
                        "type": "array",
                        "contains": {
                          "type": "string",
                          "pattern": "https://trustbloc.github.io/context/vc/examples-crude-product-v1.jsonld"
                        }
                      }
                    }
                  ]
                }
              }
            ]
          }
        ]
      },
      "createDID": true
    },
    {
      "verifier": {
        "id": "v_myprofile_multivp_jwt",
        "version": "v1.0",
        "name": "v_myprofile_multivp_jwt",
        "organizationID": "00000000-0000-0000-0000-000000000001",
        "url": "https://test-verifier.com",
        "logoURL": "https://test-verifier.com/logo.png",
        "active": true,
        "webHook": "http://vcs.webhook.example.com:8180",
        "checks": {
          "credential": {
            "format": [
              "jwt"
            ],
            "proof": true,
            "status": true,
            "strict": true
          },
          "presentation": {
            "format": [
              "jwt"
            ],
            "vcSubject": true,
            "proof": true
          }
        },
        "oidcConfig": {
          "roSigningAlgorithm": "EcdsaSecp256k1Signature2019",
          "keyType": "ECDSASecp256k1DER",
          "didMethod": "ion"
        },
        "presentationDefinitions": [
          {
            "id": "32f54163-7166-48f1-93d8-ff217bdb0654",
            "input_descriptors": [
              {
                "id": "type",
                "name": "type",
                "purpose": "We can only interact with specific status information for Verifiable Credentials",
                "schema": [
                  {
                    "uri": "https://www.w3.org/2018/credentials#VerifiableCredential"
                  }
                ],
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$.credentialStatus.type",
                        "$.vc.credentialStatus.type"
                      ],
                      "purpose": "We can only interact with specific status information for Verifiable Credentials",
                      "filter": {
                        "type": "string",
                        "enum": [
                          "StatusList2021Entry",
                          "RevocationList2021Status",
                          "RevocationList2020Status"
                        ]
                      }
                    }
                  ]
                }
              },
              {
                "id": "degree",
                "name": "degree",
                "purpose": "We can only hire with bachelor degree.",
                "schema": [
                  {
                    "uri": "https://www.w3.org/2018/credentials#VerifiableCredential"
                  }
                ],
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.degree.type",
                        "$.vc.credentialSubject.degree.type"
                      ],
                      "purpose": "We can only hire with bachelor degree.",
                      "filter": {
                        "type": "string",
                        "const": "BachelorDegree"
                      }
                    }
                  ]
                }
              }
            ]
          }
        ]
      },
      "createDID": true
    },
    {
      "verifier": {
        "id": "v_myprofile_jwt_whitelist",
        "version": "v1.0",
        "name": "v_myprofile_jwt_whitelist",
        "organizationID": "00000000-0000-0000-0000-000000000001",
        "url": "https://test-verifier.com",
        "active": true,
        "webHook": "http://vcs.webhook.example.com:8180",
        "checks": {
          "credential": {
            "format": [
              "jwt"
            ],
            "issuerTrustList": {
              "bank_issuer" : {},
              "bank_issuer_sdjwt_v5": {
                "credentialTypes": [
                  "CrudeProductCredential"
                ]
              }
            },
            "proof": true,
            "status": true,
            "strict": true
          },
          "presentation": {
            "format": [
              "jwt"
            ],
            "vcSubject": true,
            "proof": true
          }
        },
        "oidcConfig": {
          "roSigningAlgorithm": "EcdsaSecp256k1Signature2019",
          "keyType": "ECDSASecp256k1DER",
          "didMethod": "ion"
        },
        "presentationDefinitions": [
          {
            "id": "32f54163-no-limit-disclosure-single-field",
            "input_descriptors": [
              {
                "id": "degree",
                "name": "degree",
                "purpose": "We can only hire with bachelor degree.",
                "schema": [
                  {
                    "uri": "https://www.w3.org/2018/credentials#VerifiableCredential"
                  }
                ],
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.degree.type",
                        "$.vc.credentialSubject.degree.type"
                      ],
                      "id": "degree_type_id",
                      "purpose": "We can only hire with bachelor degree.",
                      "filter": {
                        "type": "string",
                        "const": "BachelorDegree"
                      }
                    }
                  ]
                }
              }
            ]
          },
          {
            "id": "32f54163-no-limit-disclosure-optional-fields",
            "input_descriptors": [
              {
                "id": "lprCategory",
                "name": "lprCategory",
                "purpose": "Permanent Resident Card specification",
                "schema": [
                  {
                    "uri": "https://www.w3.org/2018/credentials#VerifiableCredential"
                  }
                ],
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.lprCategory",
                        "$.vc.credentialSubject.lprCategory"
                      ],
                      "id": "lpr_category_id",
                      "purpose": "Specific LPR category.",
                      "filter": {
                        "type": "string",
                        "const": "C09"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.commuterClassification",
                        "$.vc.credentialSubject.commuterClassification"
                      ],
                      "id": "commuter_classification",
                      "purpose": "Specific commuter classification.",
                      "filter": {
                        "type": "string",
                        "const": "C1"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.registrationCity",
                        "$.vc.credentialSubject.registrationCity"
                      ],
                      "id": "registration_city",
                      "purpose": "Specific registration city.",
                      "optional": true,
                      "filter": {
                        "type": "string",
                        "const": "Albuquerque"
                      }
                    }
                  ]
                }
              }
            ]
          },
          {
            "id": "3c8b1d9a-limit-disclosure-optional-fields",
            "input_descriptors": [
              {
                "id": "uom",
                "name": "uom",
                "purpose": "Crude oil stream specification.",
                "constraints": {
                  "limit_disclosure": "required",
                  "fields": [
                    {
                      "path": [
                        "$.credentialSubject.physicalSpecs.uom",
                        "$.vc.credentialSubject.physicalSpecs.uom"
                      ],
                      "id": "unit_of_measure_barrel",
                      "purpose": "We can only use barrel UoM.",
                      "filter": {
                        "type": "string",
                        "const": "barrel"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.physicalSpecs.apiGravity",
                        "$.vc.credentialSubject.physicalSpecs.apiGravity"
                      ],
                      "id": "api_gravity",
                      "purpose": "Min API Gravity.",
                      "filter": {
                        "type": "integer",
                        "minimum": 20
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.category",
                        "$.vc.credentialSubject.category"
                      ],
                      "id": "category",
                      "purpose": "Category.",
                      "optional": true,
                      "filter": {
                        "type": "string"
                      }
                    },
                    {
                      "path": [
                        "$.credentialSubject.supplierAddress",
                        "$.vc.credentialSubject.supplierAddress"
                      ],
                      "id": "supplier_address",
                      "purpose": "Supplier Address.",
                      "optional": true,
                      "filter": {
                        "type": "object"
                      }
                    }
                  ]
                }
              }
            ]
          },
          {
            "id": "lp403pb9-schema-match",
            "input_descriptors": [
              {
                "id": "schema",
                "name": "schema",
                "purpose": "Match credentials using specific schema.",
                "constraints": {
                  "fields": [
                    {
                      "path": [
                        "$[\"@context\"]"
                      ],
                      "id": "schema_id",
                      "purpose": "Match credentials using specific schema.",
                      "filter": {
                        "type": "array",
                        "contains": {
                          "type": "string",
                          "pattern": "https://trustbloc.github.io/context/vc/examples-crude-product-v1.jsonld"
                        }
                      }
                    }
                  ]
                }
              }
            ]
          }
        ]
      },
      "createDID": true
    }
  ]
}
