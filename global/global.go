package global

import (
	"crypto/tls"
	"net/http"
	"time"
	"wihscan/dataType"
)

var (
	// 版本号
	Version = "1.0"

	// 规则路径
	RulePath string

	// 超时时间
	TimeOut time.Duration

	// 自定义 Tr
	Tr = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			//VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			//	var cert, err = x509.ParseCertificate(rawCerts[0])
			//	if err != nil {
			//		fmt.Println(err)
			//	}
			//	fmt.Println("证书有效期：", cert.NotBefore, " 到 ", cert.NotAfter)
			//	return nil
			//},
			//Renegotiation:      tls.RenegotiateOnceAsClient,
		}, // 跳过https证书

		//DialContext: (&net.Dialer{
		//	Timeout:   TimeOut, // 超时时间
		//	KeepAlive: TimeOut, // keepAlive 超时时间
		//}).DialContext,

		//MaxConnsPerHost:       5,
		//MaxIdleConns:          100, // 最大空闲连接数
		//MaxIdleConnsPerHost:   GoThreadNum * 2,
		//IdleConnTimeout:       TimeOut, // 空闲连接超时
		//TLSHandshakeTimeout:   TimeOut, // TLS 握手超时
		//DisableKeepAlives:     true,
		//ExpectContinueTimeout: TimeOut,
	}

	RuloWIH *dataType.WIH

	Debug bool
)
