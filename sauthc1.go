package sauthc1

import (
	"net/http"
	"time"
)

type ApiKey string

type Sauthc1Signer struct {
	DefaultAlgorithm     string
	HostHeader           string
	AuthorizationHeader  string
	StorpathDateHeader   string
	IdTerminator         string
	Algorithm            string
	AuthenticationScheme string
	Sauthc1Id            string
	Sauthc1SignedHeaders string
	Sauthc1Signature     string
	DateFormat           string
	TimestampFormat      string
	Newline              string
}

// DefaultSigner returns a new Sauthc1Signer object initialized with the default
// values.
func DefaultSigner() Sauthc1Signer {
	s := Sauthc1Signer{
		DefaultAlgorithm:     "SHA256",
		HostHeader:           "Host",
		AuthorizationHeader:  "Authorization",
		StorpathDateHeader:   "X-Stormpath-Date",
		IdTerminator:         "sauthc1_request",
		Algorithm:            "HMAC-SHA-256",
		AuthenticationScheme: "SAuthc1",
		Sauthc1Id:            "sauthc1Id",
		Sauthc1SignedHeaders: "sauthc1SignedHeaders",
		Sauthc1Signature:     "sauthc1Signature",
		DateFormat:           "%Y%m%d",
		TimestampFormat:      "%Y%m%dT%H%M%SZ",
		Newline:              "\n",
	}
	return s
}

func (s *Sauthc1Signer) Sign(req *http.Request, key ApiKey) error {
	now := time.Now()
	timeStamp := now.Format(s.TimestampFormat)
	dateStamp := now.Format(s.DateFormat)
	print(timeStamp)
	print("\n")
	print(dateStamp)
	print("\n")
	return nil // Success!
}
