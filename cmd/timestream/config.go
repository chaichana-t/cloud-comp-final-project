package timestream

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"time"
)

var writeSvc *timestreamwrite.TimestreamWrite

func init() {
	tr := &http.Transport{
		ResponseHeaderTimeout: 20 * time.Second,
		// Using DefaultTransport values for other parameters: https://golang.org/pkg/net/http/#RoundTripper
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: 30 * time.Second,
			DualStack: true,
			Timeout:   30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	
	// So client makes HTTP/2 requests
	http2.ConfigureTransport(tr)
	
	sess, err := session.NewSession(&aws.Config{ Region: aws.String("us-west-2"), MaxRetries: aws.Int(10), HTTPClient: &http.Client{ Transport: tr }})
	if err != nil {
		panic(err)
	}
	writeSvc = timestreamwrite.New(sess)
}
