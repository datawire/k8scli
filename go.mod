module github.com/datawire/k8scli

require (
	contrib.go.opencensus.io/exporter/ocagent v0.4.3 // indirect
	github.com/Azure/go-autorest v11.4.0+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/ericchiang/k8s v1.2.0
	github.com/evanphx/json-patch v4.1.0+incompatible // indirect
	github.com/gogo/protobuf v1.2.0 // indirect
	github.com/google/btree v0.0.0-20180813153112-4030bb1f1f0c // indirect
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gophercloud/gophercloud v0.0.0-20190204021847-a2b0ad6ce68c // indirect
	github.com/gregjones/httpcache v0.0.0-20190203031600-7a902570cb17 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/spf13/cobra v0.0.3 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.3.0 // indirect
	go.opencensus.io v0.19.0 // indirect
	golang.org/x/crypto v0.0.0-20190131182504-b8fe1690c613 // indirect
	golang.org/x/net v0.0.0-20190125091013-d26f9f9a57f3 // indirect
	golang.org/x/oauth2 v0.0.0-20190130055435-99b60b757ec1 // indirect
	golang.org/x/sys v0.0.0-20190204103248-980327fe3c65 // indirect
	golang.org/x/time v0.0.0-20181108054448-85acf8d2951c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/klog v0.1.0 // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)

// The "v0.0.0-2019…" versions are the tag "kubernetes-1.13.3", but
// `go build` (in its infinite wisdom) wants to edit the file to not
// be useful to humans.  <https://github.com/golang/go/issues/27271>
// <https://github.com/golang/go/issues/25898>
//
// clieng-go v10 is the version corresponding to Kubernetes 1.13.
// These 4 packages should all be upgraded together (for example,
// client-go v10 won't build with the other packages using
// v1.14.0-alpha versions
// <https://github.com/kubernetes/client-go/issues/551>)
require (
	k8s.io/api v0.0.0-20190202010724-74b699b93c15 // indirect
	k8s.io/apimachinery v0.0.0-20190117220443-572dfc7bdfcb
	k8s.io/cli-runtime v0.0.0-20190202014047-491c94071cfa
	k8s.io/client-go v10.0.0+incompatible
)
