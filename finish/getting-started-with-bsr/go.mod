module github.com/bufbuild/buf-tour

go 1.22.2

require (
	buf.build/gen/go/<USERNAME>/petapis/connectrpc/go v1.16.1-20240424002215-e3c35074dc9a.1
	buf.build/gen/go/<USERNAME>/petapis/protocolbuffers/go v1.33.0-20240424002215-e3c35074dc9a.1
	connectrpc.com/connect v1.16.1
	golang.org/x/net v0.24.0
	google.golang.org/genproto v0.0.0-20240429193739-8cf5692501f6
	google.golang.org/protobuf v1.33.0
)

require golang.org/x/text v0.14.0 // indirect
