package httptracer

import "time"

type TracerResult struct {
	NameLookup   time.Duration `json:"name_lookup" xml:"name_lookup"`
	Connect      time.Duration `json:"connect" xml:"connect"`
	TLSHandshake time.Duration `json:"tls_handshake" xml:"tls_handshake"`
	FirstByte    time.Duration `json:"first_byte" xml:"first_byte"`
	FullResponse time.Duration `json:"full_response" xml:"full_response"`
	BodySize     int           `json:"body_size" xml:"body_size"`
}
