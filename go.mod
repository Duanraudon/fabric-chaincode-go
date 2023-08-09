module github.com/hyperledger/fabric-chaincode-go

go 1.20

require (
	github.com/golang/protobuf v1.5.3
	github.com/hyperledger/fabric v2.1.1+incompatible
	github.com/hyperledger/fabric-protos-go v0.3.0
	github.com/tjfoc/gmsm v1.4.1
	github.com/tjfoc/gmtls v1.2.1
	google.golang.org/grpc v1.57.0
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace (
	github.com/tjfoc/gmsm v1.4.1 => gitea.com/bsn-go-library/gmsm v0.0.0-20230809075355-477db5e34743
	github.com/tjfoc/gmtls v1.2.1 => gitea.com/bsn-go-library/gmtls v0.0.0-20230809075814-c2a6c3a07fc8
)
