package httptracer

import (
	"crypto/tls"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"
)

func Trace(url, method string) (*TracerResult, error) {
	req, _ := http.NewRequest(strings.ToUpper(method), url, nil)

	var startTime, connectStartTime, nameLookupStartTime, tlsHandshakeStartTime time.Time
	httpStatData := TracerResult{}

	trace := &httptrace.ClientTrace{
		// NameLookup time
		DNSStart: func(dsi httptrace.DNSStartInfo) { nameLookupStartTime = time.Now() },
		DNSDone:  func(ddi httptrace.DNSDoneInfo) { httpStatData.NameLookup = time.Since(nameLookupStartTime) },

		// TLSHandshake time
		TLSHandshakeStart: func() { tlsHandshakeStartTime = time.Now() },
		TLSHandshakeDone:  func(cs tls.ConnectionState, err error) { httpStatData.TLSHandshake = time.Since(tlsHandshakeStartTime) },

		// Connect time
		ConnectStart: func(network, addr string) { connectStartTime = time.Now() },
		ConnectDone:  func(network, addr string, err error) { httpStatData.Connect = time.Since(connectStartTime) },

		// FirstByte time
		GotFirstResponseByte: func() { httpStatData.FirstByte = time.Since(startTime) },
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	startTime = time.Now()
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// FullResponse time
	httpStatData.FullResponse = time.Since(startTime)

	// BodySize
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	httpStatData.BodySize = binary.Size(bodyBytes)

	return &httpStatData, nil
}

func Tracer(url, method string) (*TracerResult, error) {
	return Trace(url, method)
}

func (d *TracerResult) ToJSON() ([]byte, error) {
	return json.Marshal(d)
}

func (d *TracerResult) ToXML() ([]byte, error) {
	return xml.Marshal(d)
}
