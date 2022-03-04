module github.com/sbwise01/cwmetricdecode

go 1.17

require (
	github.com/golang/protobuf v1.5.2
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/open-telemetry/opentelemetry-proto v0.7.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5 // indirect
)

require google.golang.org/protobuf v1.27.1 // indirect

replace github.com/open-telemetry/opentelemetry-proto => ./pkg/github.com/open-telemetry/opentelemetry-proto
