module github.com/b2wdigital/fxstack

go 1.13

replace github.com/b2wdigital/goignite => ../goignite

require (
	github.com/aws/aws-lambda-go v1.22.0
	github.com/aws/aws-sdk-go-v2 v1.2.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.1.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.1.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.1.1
	github.com/b2wdigital/goignite v0.0.0-00010101000000-000000000000
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/google/uuid v1.2.0
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/labstack/echo/v4 v4.1.17
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2
	github.com/nats-io/gnatsd v1.4.1
	github.com/nats-io/go-nats v1.7.2 // indirect
	github.com/nats-io/nats-server v1.4.1
	github.com/nats-io/nats.go v1.10.0
	github.com/newrelic/go-agent/v3 v3.10.0
	github.com/prometheus/client_golang v1.9.0
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/yudai/gojsondiff v1.0.0
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	go.uber.org/automaxprocs v1.4.0
	go.uber.org/fx v1.13.1
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
)
