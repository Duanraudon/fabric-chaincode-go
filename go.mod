module github.com/hyperledger/fabric-chaincode-go

go 1.17

require (
	github.com/golang/protobuf v1.5.2
	github.com/hyperledger/fabric-protos-go v0.0.0-20220613214546-bf864f01d75e
	github.com/stretchr/testify v1.8.0
	github.com/tjfoc/gmsm v1.4.1
	github.com/tjfoc/gmtls v1.2.1
	google.golang.org/grpc v1.48.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/tjfoc/gmtls v1.2.1 => gitea.com/tjfoc/gmtls v1.2.3

replace github.com/tjfoc/gmsm v1.4.1 => gitea.com/tjfoc/gmsm v1.1.3
