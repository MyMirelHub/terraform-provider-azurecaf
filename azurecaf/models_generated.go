// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-10-17 18:23:35.123569 +0100 BST m=+0.016720334
// using data from
// resourceDefinition.json and resourceDefinition_out_of_docs.json

package azurecaf

// ResourceDefinitions are a map of difinitions for the resources supported
var ResourceDefinitions = map[string]ResourceStructure{
	"gke_kubernetes_cluster": {"gke_kubernetes_cluster", "gke", 1, 63, false, "[^0-9A-Za-z_-]", "^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]$", true, "global"},
}

// ResourceMaps are a map from the slug to the resource definition
var ResourceMaps = map[string]string{
	"gke": "gke_kubernetes_cluster",
}
