module github.com/b2wdigital/fxstack

go 1.13

require (
	github.com/aws/aws-lambda-go v1.13.3
	github.com/aws/aws-sdk-go-v2 v0.28.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v0.28.0
	github.com/aws/aws-sdk-go-v2/service/sns v0.28.0
	github.com/b2wdigital/goignite v1.8.0
	github.com/cloudevents/sdk-go/v2 v2.0.0-preview8
	github.com/google/uuid v1.1.1
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/labstack/echo/v4 v4.1.16
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2
	github.com/nats-io/gnatsd v1.4.1
	github.com/nats-io/go-nats v1.7.2 // indirect
	github.com/nats-io/nats-server v1.4.1
	github.com/nats-io/nats.go v1.9.2
	github.com/newrelic/go-agent/v3 v3.9.0
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.5.1
	github.com/yudai/gojsondiff v1.0.0
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	go.uber.org/automaxprocs v1.3.0
	go.uber.org/fx v1.13.0
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
)

// replace github.com/b2wdigital/goignite => ../goignite
