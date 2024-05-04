package resources

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"



var namespaceTbdOperatorsSystem = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "Namespace",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
				"control-plane":                        "controller-manager",
			},
			"name": "tbd-operators-system",
		},
	},
}



var customResourceDefinitionPlatformconfigsDeployPlatformTbdIo = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "apiextensions.k8s.io/v1",
		"kind":       "CustomResourceDefinition",
		"metadata": map[string]interface{}{
			"annotations": map[string]interface{}{
				"controller-gen.kubebuilder.io/version": "v0.14.0",
			},
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name": "platformconfigs.deploy.platform.tbd.io",
		},
		"spec": map[string]interface{}{
			"group": "deploy.platform.tbd.io",
			"names": map[string]interface{}{
				"kind":     "PlatformConfig",
				"listKind": "PlatformConfigList",
				"plural":   "platformconfigs",
				"singular": "platformconfig",
			},
			"scope": "Cluster",
			"versions": []interface{}{
				map[string]interface{}{
					"name": "v1alpha1",
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"description": "PlatformConfig is the Schema for the platformconfigs API.",
							"properties": map[string]interface{}{
								"apiVersion": map[string]interface{}{
									"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
									"type": "string",
								},
								"kind": map[string]interface{}{
									"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
									"type": "string",
								},
								"metadata": map[string]interface{}{
									"type": "object",
								},
								"spec": map[string]interface{}{
									"description": "PlatformConfigSpec defines the desired state of PlatformConfig.",
									"properties": map[string]interface{}{
										"cloud": map[string]interface{}{
											"properties": map[string]interface{}{
												"local": map[string]interface{}{
													"default": true,
													"description": `(Default: true)


	Whether this cloud is deployed as a local cloud to use for testing scenarios.`,
													"type": "boolean",
												},
												"type": map[string]interface{}{
													"default": "aws",
													"description": `(Default: "aws")


	Underlying cloud type this platform is deployed upon.  Currently, only AWS is supported.`,
													"enum": []interface{}{
														"aws",
													},
													"type": "string",
												},
											},
											"type": "object",
										},
										"platform": map[string]interface{}{
											"properties": map[string]interface{}{
												"certificates": map[string]interface{}{
													"properties": map[string]interface{}{
														"deploymentSize": map[string]interface{}{
															"default": "small",
															"description": `(Default: "small")
Size of the


	deployment for the underlying capability.  Must be one of small, medium, or large.`,
															"type": "string",
														},
														"namespace": map[string]interface{}{
															"default": "tbd-certificates-system",
															"description": `(Default: "tbd-certificates-system")
Namespace where


	the capability components will be deployed.`,
															"type": "string",
														},
													},
													"type": "object",
												},
												"identity": map[string]interface{}{
													"properties": map[string]interface{}{
														"deploymentSize": map[string]interface{}{
															"default": "small",
															"description": `(Default: "small")
Size of the


	deployment for the underlying capability.  Must be one of small, medium, or large.`,
															"type": "string",
														},
														"namespace": map[string]interface{}{
															"default": "tbd-identity-system",
															"description": `(Default: "tbd-identity-system")
Namespace where


	the capability components will be deployed.`,
															"type": "string",
														},
													},
													"type": "object",
												},
											},
											"type": "object",
										},
									},
									"type": "object",
								},
								"status": map[string]interface{}{
									"description": "PlatformConfigStatus defines the observed state of PlatformConfig.",
									"properties": map[string]interface{}{
										"conditions": map[string]interface{}{
											"items": map[string]interface{}{
												"description": `PhaseCondition describes an event that has occurred during a phase
of the controller reconciliation loop.`,
												"properties": map[string]interface{}{
													"lastModified": map[string]interface{}{
														"description": "LastModified defines the time in which this component was updated.",
														"type":        "string",
													},
													"message": map[string]interface{}{
														"description": "Message defines a helpful message from the phase.",
														"type":        "string",
													},
													"phase": map[string]interface{}{
														"description": "Phase defines the phase in which the condition was set.",
														"type":        "string",
													},
													"state": map[string]interface{}{
														"description": "PhaseState defines the current state of the phase.",
														"enum": []interface{}{
															"Complete",
															"Reconciling",
															"Failed",
															"Pending",
														},
														"type": "string",
													},
												},
												"required": []interface{}{
													"lastModified",
													"message",
													"phase",
													"state",
												},
												"type": "object",
											},
											"type": "array",
										},
										"created": map[string]interface{}{
											"type": "boolean",
										},
										"dependenciesSatisfied": map[string]interface{}{
											"type": "boolean",
										},
										"resources": map[string]interface{}{
											"items": map[string]interface{}{
												"description": "ChildResource is the resource and its condition as stored on the workload custom resource's status field.",
												"properties": map[string]interface{}{
													"condition": map[string]interface{}{
														"description": "ResourceCondition defines the current condition of this resource.",
														"properties": map[string]interface{}{
															"created": map[string]interface{}{
																"description": "Created defines whether this object has been successfully created or not.",
																"type":        "boolean",
															},
															"lastModified": map[string]interface{}{
																"description": "LastModified defines the time in which this resource was updated.",
																"type":        "string",
															},
															"message": map[string]interface{}{
																"description": "Message defines a helpful message from the resource phase.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"created",
														},
														"type": "object",
													},
													"group": map[string]interface{}{
														"description": "Group defines the API Group of the resource.",
														"type":        "string",
													},
													"kind": map[string]interface{}{
														"description": "Kind defines the kind of the resource.",
														"type":        "string",
													},
													"name": map[string]interface{}{
														"description": "Name defines the name of the resource from the metadata.name field.",
														"type":        "string",
													},
													"namespace": map[string]interface{}{
														"description": "Namespace defines the namespace in which this resource exists in.",
														"type":        "string",
													},
													"version": map[string]interface{}{
														"description": "Version defines the API Version of the resource.",
														"type":        "string",
													},
												},
												"required": []interface{}{
													"group",
													"kind",
													"name",
													"namespace",
													"version",
												},
												"type": "object",
											},
											"type": "array",
										},
									},
									"type": "object",
								},
							},
							"type": "object",
						},
					},
					"served":  true,
					"storage": true,
					"subresources": map[string]interface{}{
						"status": map[string]interface{}{},
					},
				},
			},
		},
	},
}


var customResourceDefinitionPlatformoperatorsDeployPlatformTbdIo = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "apiextensions.k8s.io/v1",
		"kind":       "CustomResourceDefinition",
		"metadata": map[string]interface{}{
			"annotations": map[string]interface{}{
				"controller-gen.kubebuilder.io/version": "v0.14.0",
			},
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name": "platformoperators.deploy.platform.tbd.io",
		},
		"spec": map[string]interface{}{
			"group": "deploy.platform.tbd.io",
			"names": map[string]interface{}{
				"kind":     "PlatformOperators",
				"listKind": "PlatformOperatorsList",
				"plural":   "platformoperators",
				"singular": "platformoperators",
			},
			"scope": "Cluster",
			"versions": []interface{}{
				map[string]interface{}{
					"name": "v1alpha1",
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"description": "PlatformOperators is the Schema for the platformoperators API.",
							"properties": map[string]interface{}{
								"apiVersion": map[string]interface{}{
									"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
									"type": "string",
								},
								"kind": map[string]interface{}{
									"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
									"type": "string",
								},
								"metadata": map[string]interface{}{
									"type": "object",
								},
								"spec": map[string]interface{}{
									"description": "PlatformOperatorsSpec defines the desired state of PlatformOperators.",
									"properties": map[string]interface{}{
										"namespace": map[string]interface{}{
											"default":     "tbd-operators-system",
											"description": "(Default: \"tbd-operators-system\")",
											"type":        "string",
										},
									},
									"type": "object",
								},
								"status": map[string]interface{}{
									"description": "PlatformOperatorsStatus defines the observed state of PlatformOperators.",
									"properties": map[string]interface{}{
										"conditions": map[string]interface{}{
											"items": map[string]interface{}{
												"description": `PhaseCondition describes an event that has occurred during a phase
of the controller reconciliation loop.`,
												"properties": map[string]interface{}{
													"lastModified": map[string]interface{}{
														"description": "LastModified defines the time in which this component was updated.",
														"type":        "string",
													},
													"message": map[string]interface{}{
														"description": "Message defines a helpful message from the phase.",
														"type":        "string",
													},
													"phase": map[string]interface{}{
														"description": "Phase defines the phase in which the condition was set.",
														"type":        "string",
													},
													"state": map[string]interface{}{
														"description": "PhaseState defines the current state of the phase.",
														"enum": []interface{}{
															"Complete",
															"Reconciling",
															"Failed",
															"Pending",
														},
														"type": "string",
													},
												},
												"required": []interface{}{
													"lastModified",
													"message",
													"phase",
													"state",
												},
												"type": "object",
											},
											"type": "array",
										},
										"created": map[string]interface{}{
											"type": "boolean",
										},
										"dependenciesSatisfied": map[string]interface{}{
											"type": "boolean",
										},
										"resources": map[string]interface{}{
											"items": map[string]interface{}{
												"description": "ChildResource is the resource and its condition as stored on the workload custom resource's status field.",
												"properties": map[string]interface{}{
													"condition": map[string]interface{}{
														"description": "ResourceCondition defines the current condition of this resource.",
														"properties": map[string]interface{}{
															"created": map[string]interface{}{
																"description": "Created defines whether this object has been successfully created or not.",
																"type":        "boolean",
															},
															"lastModified": map[string]interface{}{
																"description": "LastModified defines the time in which this resource was updated.",
																"type":        "string",
															},
															"message": map[string]interface{}{
																"description": "Message defines a helpful message from the resource phase.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"created",
														},
														"type": "object",
													},
													"group": map[string]interface{}{
														"description": "Group defines the API Group of the resource.",
														"type":        "string",
													},
													"kind": map[string]interface{}{
														"description": "Kind defines the kind of the resource.",
														"type":        "string",
													},
													"name": map[string]interface{}{
														"description": "Name defines the name of the resource from the metadata.name field.",
														"type":        "string",
													},
													"namespace": map[string]interface{}{
														"description": "Namespace defines the namespace in which this resource exists in.",
														"type":        "string",
													},
													"version": map[string]interface{}{
														"description": "Version defines the API Version of the resource.",
														"type":        "string",
													},
												},
												"required": []interface{}{
													"group",
													"kind",
													"name",
													"namespace",
													"version",
												},
												"type": "object",
											},
											"type": "array",
										},
									},
									"type": "object",
								},
							},
							"type": "object",
						},
					},
					"served":  true,
					"storage": true,
					"subresources": map[string]interface{}{
						"status": map[string]interface{}{},
					},
				},
			},
		},
	},
}


var serviceAccountPlatformConfigOperatorcontrollerManager = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "ServiceAccount",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name":      "platform-config-operatorcontroller-manager",
			"namespace": "tbd-operators-system",
		},
	},
}


var rolePlatformConfigOperatorleaderElectionRole = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "rbac.authorization.k8s.io/v1",
		"kind":       "Role",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name":      "platform-config-operatorleader-election-role",
			"namespace": "tbd-operators-system",
		},
		"rules": []interface{}{
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"configmaps",
				},
				"verbs": []interface{}{
					"get",
					"list",
					"watch",
					"create",
					"update",
					"patch",
					"delete",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"coordination.k8s.io",
				},
				"resources": []interface{}{
					"leases",
				},
				"verbs": []interface{}{
					"get",
					"list",
					"watch",
					"create",
					"update",
					"patch",
					"delete",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"events",
				},
				"verbs": []interface{}{
					"create",
					"patch",
				},
			},
		},
	},
}


var clusterRolePlatformConfigOperatormanagerRole = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "rbac.authorization.k8s.io/v1",
		"kind":       "ClusterRole",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name": "platform-config-operatormanager-role",
		},
		"rules": []interface{}{
			map[string]interface{}{
				"nonResourceURLs": []interface{}{
					"/metrics",
				},
				"verbs": []interface{}{
					"get",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"acme.cert-manager.io",
				},
				"resources": []interface{}{
					"challenges",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"deletecollection",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"acme.cert-manager.io",
				},
				"resources": []interface{}{
					"challenges/finalizers",
				},
				"verbs": []interface{}{
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"acme.cert-manager.io",
				},
				"resources": []interface{}{
					"challenges/status",
				},
				"verbs": []interface{}{
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"acme.cert-manager.io",
				},
				"resources": []interface{}{
					"orders",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"deletecollection",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"acme.cert-manager.io",
				},
				"resources": []interface{}{
					"orders/finalizers",
				},
				"verbs": []interface{}{
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"acme.cert-manager.io",
				},
				"resources": []interface{}{
					"orders/status",
				},
				"verbs": []interface{}{
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"admissionregistration.k8s.io",
				},
				"resources": []interface{}{
					"mutatingwebhookconfigurations",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"admissionregistration.k8s.io",
				},
				"resources": []interface{}{
					"validatingwebhookconfigurations",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"apiextensions.k8s.io",
				},
				"resources": []interface{}{
					"customresourcedefinitions",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"apiregistration.k8s.io",
				},
				"resources": []interface{}{
					"apiservices",
				},
				"verbs": []interface{}{
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"apps",
				},
				"resources": []interface{}{
					"deployments",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"authentication.k8s.io",
				},
				"resources": []interface{}{
					"tokenreviews",
				},
				"verbs": []interface{}{
					"create",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"authorization.k8s.io",
				},
				"resources": []interface{}{
					"subjectaccessreviews",
				},
				"verbs": []interface{}{
					"create",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"certificaterequests",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"deletecollection",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"certificaterequests/finalizers",
				},
				"verbs": []interface{}{
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"certificaterequests/status",
				},
				"verbs": []interface{}{
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"certificates",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"deletecollection",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"certificates/finalizers",
				},
				"verbs": []interface{}{
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"certificates/status",
				},
				"verbs": []interface{}{
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"clusterissuers",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"clusterissuers/status",
				},
				"verbs": []interface{}{
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"issuers",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"deletecollection",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"issuers/status",
				},
				"verbs": []interface{}{
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"cert-manager.io",
				},
				"resources": []interface{}{
					"signers",
				},
				"verbs": []interface{}{
					"approve",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"certificates.k8s.io",
				},
				"resources": []interface{}{
					"certificatesigningrequests",
				},
				"verbs": []interface{}{
					"create",
					"get",
					"list",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"certificates.k8s.io",
				},
				"resources": []interface{}{
					"certificatesigningrequests/status",
				},
				"verbs": []interface{}{
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"certificates.k8s.io",
				},
				"resources": []interface{}{
					"signers",
				},
				"verbs": []interface{}{
					"sign",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"certificates.platform.tbd.io",
				},
				"resources": []interface{}{
					"certmanagers",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"certificates.platform.tbd.io",
				},
				"resources": []interface{}{
					"certmanagers/status",
				},
				"verbs": []interface{}{
					"get",
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"certificates.platform.tbd.io",
				},
				"resources": []interface{}{
					"trustmanagers",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"certificates.platform.tbd.io",
				},
				"resources": []interface{}{
					"trustmanagers/status",
				},
				"verbs": []interface{}{
					"get",
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"coordination.k8s.io",
				},
				"resources": []interface{}{
					"leases",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"configmaps",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"endpoints",
				},
				"verbs": []interface{}{
					"get",
					"list",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"events",
				},
				"verbs": []interface{}{
					"create",
					"get",
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"namespaces",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"pods",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"secrets",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"serviceaccounts",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"",
				},
				"resources": []interface{}{
					"services",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"deploy.platform.tbd.io",
				},
				"resources": []interface{}{
					"platformconfigs",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"deploy.platform.tbd.io",
				},
				"resources": []interface{}{
					"platformconfigs/status",
				},
				"verbs": []interface{}{
					"get",
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"deploy.platform.tbd.io",
				},
				"resources": []interface{}{
					"platformoperators",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"deploy.platform.tbd.io",
				},
				"resources": []interface{}{
					"platformoperators/status",
				},
				"verbs": []interface{}{
					"get",
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"gateway.networking.k8s.io",
				},
				"resources": []interface{}{
					"gateways",
				},
				"verbs": []interface{}{
					"get",
					"list",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"gateway.networking.k8s.io",
				},
				"resources": []interface{}{
					"gateways/finalizers",
				},
				"verbs": []interface{}{
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"gateway.networking.k8s.io",
				},
				"resources": []interface{}{
					"httproutes",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"gateway.networking.k8s.io",
				},
				"resources": []interface{}{
					"httproutes/finalizers",
				},
				"verbs": []interface{}{
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"identity.platform.tbd.io",
				},
				"resources": []interface{}{
					"awspodidentitywebhooks",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"identity.platform.tbd.io",
				},
				"resources": []interface{}{
					"awspodidentitywebhooks/status",
				},
				"verbs": []interface{}{
					"get",
					"patch",
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"networking.k8s.io",
				},
				"resources": []interface{}{
					"ingresses",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"networking.k8s.io",
				},
				"resources": []interface{}{
					"ingresses/finalizers",
				},
				"verbs": []interface{}{
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"rbac.authorization.k8s.io",
				},
				"resources": []interface{}{
					"clusterrolebindings",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"rbac.authorization.k8s.io",
				},
				"resources": []interface{}{
					"clusterroles",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"rbac.authorization.k8s.io",
				},
				"resources": []interface{}{
					"rolebindings",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"rbac.authorization.k8s.io",
				},
				"resources": []interface{}{
					"roles",
				},
				"verbs": []interface{}{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"route.openshift.io",
				},
				"resources": []interface{}{
					"routes/custom-host",
				},
				"verbs": []interface{}{
					"create",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"trust.cert-manager.io",
				},
				"resources": []interface{}{
					"bundles",
				},
				"verbs": []interface{}{
					"get",
					"list",
					"update",
					"watch",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"trust.cert-manager.io",
				},
				"resources": []interface{}{
					"bundles/finalizers",
				},
				"verbs": []interface{}{
					"update",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"trust.cert-manager.io",
				},
				"resources": []interface{}{
					"bundles/status",
				},
				"verbs": []interface{}{
					"patch",
				},
			},
		},
	},
}


var clusterRolePlatformConfigOperatormetricsReader = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "rbac.authorization.k8s.io/v1",
		"kind":       "ClusterRole",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name": "platform-config-operatormetrics-reader",
		},
		"rules": []interface{}{
			map[string]interface{}{
				"nonResourceURLs": []interface{}{
					"/metrics",
				},
				"verbs": []interface{}{
					"get",
				},
			},
		},
	},
}


var clusterRolePlatformConfigOperatorproxyRole = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "rbac.authorization.k8s.io/v1",
		"kind":       "ClusterRole",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name": "platform-config-operatorproxy-role",
		},
		"rules": []interface{}{
			map[string]interface{}{
				"apiGroups": []interface{}{
					"authentication.k8s.io",
				},
				"resources": []interface{}{
					"tokenreviews",
				},
				"verbs": []interface{}{
					"create",
				},
			},
			map[string]interface{}{
				"apiGroups": []interface{}{
					"authorization.k8s.io",
				},
				"resources": []interface{}{
					"subjectaccessreviews",
				},
				"verbs": []interface{}{
					"create",
				},
			},
		},
	},
}


var roleBindingPlatformConfigOperatorleaderElectionRolebinding = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "rbac.authorization.k8s.io/v1",
		"kind":       "RoleBinding",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name":      "platform-config-operatorleader-election-rolebinding",
			"namespace": "tbd-operators-system",
		},
		"roleRef": map[string]interface{}{
			"apiGroup": "rbac.authorization.k8s.io",
			"kind":     "Role",
			"name":     "platform-config-operatorleader-election-role",
		},
		"subjects": []interface{}{
			map[string]interface{}{
				"kind":      "ServiceAccount",
				"name":      "platform-config-operatorcontroller-manager",
				"namespace": "tbd-operators-system",
			},
		},
	},
}


var clusterRoleBindingPlatformConfigOperatormanagerRolebinding = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "rbac.authorization.k8s.io/v1",
		"kind":       "ClusterRoleBinding",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name": "platform-config-operatormanager-rolebinding",
		},
		"roleRef": map[string]interface{}{
			"apiGroup": "rbac.authorization.k8s.io",
			"kind":     "ClusterRole",
			"name":     "platform-config-operatormanager-role",
		},
		"subjects": []interface{}{
			map[string]interface{}{
				"kind":      "ServiceAccount",
				"name":      "platform-config-operatorcontroller-manager",
				"namespace": "tbd-operators-system",
			},
		},
	},
}


var clusterRoleBindingPlatformConfigOperatorproxyRolebinding = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "rbac.authorization.k8s.io/v1",
		"kind":       "ClusterRoleBinding",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
			},
			"name": "platform-config-operatorproxy-rolebinding",
		},
		"roleRef": map[string]interface{}{
			"apiGroup": "rbac.authorization.k8s.io",
			"kind":     "ClusterRole",
			"name":     "platform-config-operatorproxy-role",
		},
		"subjects": []interface{}{
			map[string]interface{}{
				"kind":      "ServiceAccount",
				"name":      "platform-config-operatorcontroller-manager",
				"namespace": "tbd-operators-system",
			},
		},
	},
}


var servicePlatformConfigOperatorcontrollerManagerMetricsService = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "Service",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
				"control-plane":                        "controller-manager",
			},
			"name":      "platform-config-operatorcontroller-manager-metrics-service",
			"namespace": "tbd-operators-system",
		},
		"spec": map[string]interface{}{
			"ports": []interface{}{
				map[string]interface{}{
					"name":       "https",
					"port":       8443,
					"protocol":   "TCP",
					"targetPort": "https",
				},
			},
			"selector": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
				"control-plane":                        "controller-manager",
			},
		},
	},
}


var deploymentPlatformConfigOperatorcontrollerManager = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "apps/v1",
		"kind":       "Deployment",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app":                                  "platform-config-operator",
				"app.kubernetes.io/component":          "platform-config-operator",
				"app.kubernetes.io/created-by":         "platform-config-operator",
				"app.kubernetes.io/instance":           "manager",
				"app.kubernetes.io/managed-by":         "platform-bootstrapper",
				"app.kubernetes.io/name":               "platform-config-operator",
				"app.kubernetes.io/part-of":            "platform",
				"app.kubernetes.io/version":            "v0.0.1",
				"capabilities.tbd.io/capability":       "platform-config-operator",
				"capabilities.tbd.io/platform-version": "unstable",
				"capabilities.tbd.io/version":          "v0.0.1",
				"control-plane":                        "controller-manager",
			},
			"name":      "platform-config-operatorcontroller-manager",
			"namespace": "tbd-operators-system",
		},
		"spec": map[string]interface{}{
			"replicas": 1,
			"selector": map[string]interface{}{
				"matchLabels": map[string]interface{}{
					"app":                                  "platform-config-operator",
					"app.kubernetes.io/component":          "platform-config-operator",
					"app.kubernetes.io/instance":           "manager",
					"app.kubernetes.io/managed-by":         "platform-bootstrapper",
					"app.kubernetes.io/name":               "platform-config-operator",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/version":            "v0.0.1",
					"capabilities.tbd.io/capability":       "platform-config-operator",
					"capabilities.tbd.io/platform-version": "unstable",
					"capabilities.tbd.io/version":          "v0.0.1",
					"control-plane":                        "controller-manager",
				},
			},
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{
					"annotations": map[string]interface{}{
						"kubectl.kubernetes.io/default-container": "manager",
					},
					"labels": map[string]interface{}{
						"app":                                  "platform-config-operator",
						"app.kubernetes.io/component":          "platform-config-operator",
						"app.kubernetes.io/instance":           "manager",
						"app.kubernetes.io/managed-by":         "platform-bootstrapper",
						"app.kubernetes.io/name":               "platform-config-operator",
						"app.kubernetes.io/part-of":            "platform",
						"app.kubernetes.io/version":            "v0.0.1",
						"capabilities.tbd.io/capability":       "platform-config-operator",
						"capabilities.tbd.io/platform-version": "unstable",
						"capabilities.tbd.io/version":          "v0.0.1",
						"control-plane":                        "controller-manager",
					},
				},
				"spec": map[string]interface{}{
					"affinity": map[string]interface{}{
						"podAntiAffinity": map[string]interface{}{
							"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
								map[string]interface{}{
									"podAffinityTerm": map[string]interface{}{
										"labelSelector": map[string]interface{}{
											"matchExpressions": []interface{}{
												map[string]interface{}{
													"key":      "app.kubernetes.io/name",
													"operator": "In",
													"values": []interface{}{
														"platform-config-operator",
													},
												},
											},
										},
										"topologyKey": "kubernetes.io/hostname",
									},
									"weight": 100,
								},
							},
						},
					},
					"containers": []interface{}{
						map[string]interface{}{
							"args": []interface{}{
								"--secure-listen-address=0.0.0.0:8443",
								"--upstream=http://127.0.0.1:8080/",
								"--logtostderr=true",
								"--v=0",
							},
							"image": "gcr.io/kubebuilder/kube-rbac-proxy:v0.13.1",
							"name":  "kube-rbac-proxy",
							"ports": []interface{}{
								map[string]interface{}{
									"containerPort": 8443,
									"name":          "https",
									"protocol":      "TCP",
								},
							},
							"resources": map[string]interface{}{
								"limits": map[string]interface{}{
									"memory": "64Mi",
								},
								"requests": map[string]interface{}{
									"cpu":    "5m",
									"memory": "16Mi",
								},
							},
							"securityContext": map[string]interface{}{
								"allowPrivilegeEscalation": false,
								"capabilities": map[string]interface{}{
									"drop": []interface{}{
										"ALL",
									},
								},
								"readOnlyRootFilesystem": true,
							},
						},
						map[string]interface{}{
							"args": []interface{}{
								"--health-probe-bind-address=:8081",
								"--metrics-bind-address=127.0.0.1:8080",
								"--leader-elect",
							},
							"command": []interface{}{
								"/manager",
							},
							"image": "quay.io/tbd-paas/platform-config-operator:latest",
							"livenessProbe": map[string]interface{}{
								"httpGet": map[string]interface{}{
									"path": "/healthz",
									"port": 8081,
								},
								"initialDelaySeconds": 15,
								"periodSeconds":       20,
							},
							"name": "manager",
							"readinessProbe": map[string]interface{}{
								"httpGet": map[string]interface{}{
									"path": "/readyz",
									"port": 8081,
								},
								"initialDelaySeconds": 5,
								"periodSeconds":       10,
							},
							"resources": map[string]interface{}{
								"limits": map[string]interface{}{
									"cpu":    "125m",
									"memory": "64Mi",
								},
								"requests": map[string]interface{}{
									"cpu":    "10m",
									"memory": "16Mi",
								},
							},
							"securityContext": map[string]interface{}{
								"allowPrivilegeEscalation": false,
								"capabilities": map[string]interface{}{
									"drop": []interface{}{
										"ALL",
									},
								},
								"readOnlyRootFilesystem": true,
							},
						},
					},
					"nodeSelector": map[string]interface{}{
						"kubernetes.io/arch": "arm64",
						"kubernetes.io/os":   "linux",
						"tbd.io/node-type":   "platform",
					},
					"securityContext": map[string]interface{}{
						"fsGroup":      1001,
						"runAsGroup":   1001,
						"runAsNonRoot": true,
						"runAsUser":    1001,
					},
					"serviceAccountName":            "platform-config-operatorcontroller-manager",
					"terminationGracePeriodSeconds": 10,
				},
			},
		},
	},
}


var platformOperatorsPlatformoperatorsSample = &unstructured.Unstructured{
	Object: map[string]interface{}{
		"apiVersion": "deploy.platform.tbd.io/v1alpha1",
		"kind":       "PlatformOperators",
		"metadata": map[string]interface{}{
			"name": "platformoperators-sample",
		},
		"spec": map[string]interface{}{
			"namespace": "tbd-operators-system",
		},
	},
}


