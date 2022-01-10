module github.com/mintel/eventrouter

go 1.14

require (
	github.com/Shopify/sarama v1.23.1
	github.com/crewjam/rfc5424 v0.0.0-20180723152949-c25bdd3a0ba2
	github.com/eapache/channels v1.1.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/prometheus/client_golang v1.1.0
	github.com/sethgrid/pester v0.0.0-20190127155807-68a33a018ad0
	github.com/spf13/viper v1.4.0
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0 // indirect
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
	k8s.io/api v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v12.0.0+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.18.3
