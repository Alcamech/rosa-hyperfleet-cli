package cluster

import (
	"testing"

	v1alpha1 "github.com/openshift-online/rosa-seaplane-api/api/v1alpha1/public"
	clientsetscheme "github.com/openshift-online/rosa-seaplane-api/clientset/generated/scheme"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestClientsetSchemeDecodesClusterResponse(t *testing.T) {
	data := []byte(`{"apiVersion":"seaplane.io/v1alpha1","kind":"Cluster","metadata":{"name":"test-cluster"}}`)

	obj, gvk, err := clientsetscheme.Codecs.UniversalDeserializer().Decode(data, nil, nil)
	if err != nil {
		t.Fatalf("decode Cluster response: %v", err)
	}

	wantGVK := schema.GroupVersionKind{Group: "seaplane.io", Version: "v1alpha1", Kind: "Cluster"}
	if gvk == nil || *gvk != wantGVK {
		t.Fatalf("decoded GVK = %v, want %v", gvk, wantGVK)
	}
	if _, ok := obj.(*v1alpha1.Cluster); !ok {
		t.Fatalf("decoded object has type %T, want *v1alpha1.Cluster", obj)
	}
}
