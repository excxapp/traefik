package acmev2

// Challenge is a string that identifies a particular type and version of ACME challenge.
type Challenge string

const (
	// HTTP01 is the "http-01" ACME challenge https://github.com/ietf-wg-acme/acme/blob/master/draft-ietf-acme-acmev2.md#http
	// Note: HTTP01ChallengePath returns the URL path to fulfill this challenge
	HTTP01 = Challenge("http-01")
	// DNS01 is the "dns-01" ACME challenge https://github.com/ietf-wg-acme/acme/blob/master/draft-ietf-acme-acmev2.md#dns
	// Note: DNS01Record returns a DNS record which will fulfill this challenge
	DNS01 = Challenge("dns-01")
)