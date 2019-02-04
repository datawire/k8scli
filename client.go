package k8scli

import (
	"net/http"

	"github.com/ericchiang/k8s"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// NewOutOfClusterClientFactory is a counterpart to
// k8s.NewInClusterClient().  It adds the appropriate kubeconfig flags
// to the FlagSet, and returns a function to call after
// FlagSet.Parse() that will read those flags and the kubeconfig file
// to construct a client.
func NewOutOfClusterClientFactory(flags *pflag.FlagSet) func() (*k8s.Client, error) {
	configFlags := genericclioptions.NewConfigFlags()

	// We can disable or enable flags by setting them to
	// nil/non-nil prior to calling .AddFlags().
	//
	// .Username and .Password are already disabled by default in
	// genericclioptions.NewConfigFlags().
	//
	// Unlike client-go/discovery.CachedDiscoveryClient,
	// ericchiang/k8s doesn't support caching to the filesystem,
	// so disable the '--cache-dir' flag.
	configFlags.CacheDir = nil

	configFlags.AddFlags(flags)

	return func() (*k8s.Client, error) {
		configLoader := configFlags.ToRawKubeConfigLoader()

		namespace, _, err := configLoader.Namespace()
		if err != nil {
			return nil, err
		}

		restconfig, err := configLoader.ClientConfig()
		if err != nil {
			return nil, err
		}

		endpoint, _, err := rest.DefaultServerURL(restconfig.Host,
			"", schema.GroupVersion{},
			rest.IsConfigTransportTLS(*restconfig))
		if err != nil {
			return nil, err
		}

		transport, err := rest.TransportFor(restconfig)
		if err != nil {
			return nil, err
		}

		client := &k8s.Client{
			Endpoint:  endpoint.String(),
			Namespace: namespace,
			Client: &http.Client{
				Transport: transport,
			},
		}
		return client, nil
	}
}
