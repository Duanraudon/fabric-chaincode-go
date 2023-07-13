package gmtls

import "net/http"

func newServer(listenAddress string, sm2SigCert, sm2EncCert, stdCert *Certificate) (error, error) {
	config, err := NewBasicAutoSwitchConfig(sm2SigCert, sm2EncCert, stdCert)
	if err != nil {

	}
	listen, err := Listen("tcp", listenAddress, config)
	if err != nil {

	}
	err = http.Serve(listen, nil)
	if err != nil {

	}
	return nil, nil
}
