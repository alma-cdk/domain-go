// Domain with certificate
package almacdkdomain

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// Properties to configure the domain (zone and certificate).
// Experimental.
type DomainProps struct {
	// Provide either a fully-qualified domain name as string to perform a hosted zone lookup or a previously defined hosted zone as `route53.IHostedZone`.
	// Experimental.
	Zone interface{} `field:"required" json:"zone" yaml:"zone"`
	// Provide your own pre-existing certificate.
	//
	// If not provided, a new certificate will be created
	// by default.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Set to false to disable IPv6 `AAAA` record creation.
	// Experimental.
	EnableIpv6 *bool `field:"optional" json:"enableIpv6" yaml:"enableIpv6"`
	// AWS Region to deploy the certificate into.
	//
	// Defaults to `us-east-1` which is the only region where
	// ACM certificates can be deployed to CloudFront.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Provide subdomain or leave undefined to use the zone apex domain.
	//
	// If subdomain provided, the resulting FQDN will be `subdomain.zone`.
	// Experimental.
	Subdomain *string `field:"optional" json:"subdomain" yaml:"subdomain"`
}

