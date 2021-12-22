module github.com/openshift-online/ocm-sdk-go/examples

go 1.16

// We don't want to use the latest released versio of the SDK, but exactly the same version that
// is in the parent directory.
replace github.com/openshift-online/ocm-sdk-go => ../

require (
	github.com/openshift-online/ocm-sdk-go v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.9.0
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
)
