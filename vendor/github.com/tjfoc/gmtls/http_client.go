/*
Copyright Suzhou Tongji Fintech Research Institute 2017 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gmtls

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/tjfoc/gmsm/sm2"
)

// NewHTTPSClient 创建国密HTTPS客户端，只对服务端进行身份认证（验证服务端证书）。
// pool: 根证书链
func NewHTTPSClient(pool *sm2.CertPool, workMode string) *http.Client {
	return &http.Client{
		Transport: NewSimpleRoundTripper(&Config{
			GMSupport:  &GMSupport{workMode},
			RootCAs:    pool,
			MinVersion: VersionGMSSL,
		}),
	}
}

// NewAuthHTTPSClient 创建双向身份认证国密HTTPS客户端
// pool: 根证书链
// clientAuthCert: 客户端认证密钥对和证书
func NewAuthHTTPSClient(pool *sm2.CertPool, clientAuthCert *Certificate, workMode string) *http.Client {
	return &http.Client{
		Transport: NewSimpleRoundTripper(&Config{
			GMSupport:    &GMSupport{workMode},
			MinVersion:   VersionGMSSL,
			RootCAs:      pool,
			Certificates: []Certificate{*clientAuthCert},
		}),
	}
}

// NewCustomHTTPSClient 创建自定义国密HTTPS客户端
// 通过自定义TLS参数定制TLS实现细节，如进行双向身份认证等。
func NewCustomHTTPSClient(config *Config) *http.Client {
	if config == nil {
		return &http.Client{}
	}

	return &http.Client{
		Transport: newDefaultGMHttpClientTransport(config),
	}
}

/**
 * @description: 构建默认国密HTTP客户端Transport信息
 * @author: Andycao
 * @date: 2023/7/4
 * @params
 * @return:
 */
func newDefaultGMHttpClientTransport(tlsConfig *Config) http.RoundTripper {
	return &http.Transport{
		DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {

			dialer := &net.Dialer{}

			conn, err := DialWithDialer(dialer, network, addr, tlsConfig)
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 60 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		IdleConnTimeout:     30 * time.Second,
	}
}
