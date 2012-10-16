package sauthc1

import (
	"net/url"
	"strings"
)


// encodeUrl encodes a URL to be signed by SAuthc1.  
//
// Based on https://github.com/stormpath/stormpath-sdk-ruby/blob/master/lib/stormpath-sdk/util/request_utils.rb#L35
func encodeUrl (u *url.URL, path, canonical bool) *url.URL {
	
}


// defaultPort checks whether a URL contains the default port for its scheme: 80
// for http, 443 for https, or no port specified.
func defaultPort(u *url.URL) bool {
	scheme := u.Scheme
	parts := strings.Split(u.Host, ":")
	var port string
	switch len(parts) {
		case 1:
			port = ""
		case 2:
			port = parts[1]
		default:
			// Would it be better to return an error here?
			return false
	}
	result := port == "" || port == "0" || (port == "80" && scheme == "http") || (port == "443" && scheme == "https")
	return result // Success!
}