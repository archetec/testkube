/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

type TestWorkflowSignature struct {
	// step reference
	Ref string `json:"ref,omitempty"`
	// step name
	Name string `json:"name,omitempty"`
	// step category, that may be used as name fallback
	Category string `json:"category,omitempty"`
	// is the step/group meant to be optional
	Optional bool `json:"optional,omitempty"`
	// is the step/group meant to be negative
	Negative bool                    `json:"negative,omitempty"`
	Children []TestWorkflowSignature `json:"children,omitempty"`
}