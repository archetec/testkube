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

// Represents a source location of a volume to mount, managed by an external CSI driver
type CsiVolumeSource struct {
	// driver is the name of the CSI driver that handles this volume.
	Driver   string        `json:"driver,omitempty"`
	ReadOnly *BoxedBoolean `json:"readOnly,omitempty"`
	FsType   *BoxedString  `json:"fsType,omitempty"`
	// volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.
	VolumeAttributes     map[string]string     `json:"volumeAttributes,omitempty"`
	NodePublishSecretRef *LocalObjectReference `json:"nodePublishSecretRef,omitempty"`
}