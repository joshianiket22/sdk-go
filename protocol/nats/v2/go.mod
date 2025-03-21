module github.com/cloudevents/sdk-go/protocol/nats/v2

go 1.22
toolchain go1.24.1

replace github.com/cloudevents/sdk-go/v2 => ../../../v2

require (
	github.com/cloudevents/sdk-go/v2 v2.14.0
	github.com/nats-io/nats.go v1.31.0
)

require (
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nkeys v0.4.6 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)
