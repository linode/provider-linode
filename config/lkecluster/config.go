package lkecluster

import (
	"encoding/base64"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
)

const (
	errKubeconfigNotString = "kubeconfig attribute is not a string"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_lke_cluster", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = DecodeKubeconfig
		r.UseAsync = true
	})
}

// DecodeKubeconfig takes "kubeconfig" attribute and decodes it so that it can
// be used without additional processing.
func DecodeKubeconfig(attr map[string]interface{}) (map[string][]byte, error) {
	if attr["kubeconfig"] == nil {
		return nil, nil
	}
	s, ok := attr["kubeconfig"].(string)
	if !ok {
		return nil, errors.New(errKubeconfigNotString)
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode kubeconfig")
	}
	return map[string][]byte{
		"kubeconfig": decoded,
	}, nil
}
