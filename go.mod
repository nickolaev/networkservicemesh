module github.com/networkservicemesh/networkservicemesh

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190820101407-c8dc69f8a8bf

require (
	github.com/aws/aws-sdk-go v1.22.0
	github.com/caddyserver/caddy v1.0.1
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/coredns/coredns v1.5.2
	github.com/dnstap/golang-dnstap v0.1.0
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/ghodss/yaml v1.0.0
	github.com/go-errors/errors v1.0.1
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/ligato/cn-infra v2.0.0+incompatible // indirect
	github.com/ligato/vpp-agent v2.1.1+incompatible
	github.com/miekg/dns v1.1.15
	github.com/onsi/gomega v1.5.1-0.20190520121345-efe19c39ca10
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/packethost/packngo v0.1.1-0.20190507131943-1343be729ca2
	github.com/pkg/errors v0.8.1
	github.com/rs/xid v1.2.1
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible // indirect
	github.com/vishvananda/netlink v1.0.0
	github.com/vishvananda/netns v0.0.0-20190625233234-7109fa855b0f
	golang.org/x/net v0.0.0-20190812203447-cdfb69ac37fc
	golang.org/x/sys v0.0.0-20190618155005-516e3c20635f
	google.golang.org/grpc v1.23.0
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190820101039-d651a1528133
	k8s.io/apimachinery v0.0.0-20190820100750-21ddcbbef9e1
	k8s.io/apiserver v0.0.0-20190111033246-d50e9ac5404f // indirect
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/cluster-bootstrap v0.0.0-20190313124217-0fa624df11e9 // indirect
	k8s.io/component-base v0.0.0-20190409093041-e34633071963 // indirect
	k8s.io/kubernetes v1.14.6
)
