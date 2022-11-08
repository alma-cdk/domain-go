// Domain with certificate
package almacdkdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// Experimental.
type ICloudFrontConfiguration interface {
	// Certificate Manager certificate.
	// Experimental.
	Certificate() awscertificatemanager.ICertificate
	// Alternative domain names for this distribution.
	// Experimental.
	DomainNames() *[]*string
	// Has IPv6 AAAA records been created.
	//
	// Can be used to conditionally configure IPv6 support
	// to CloudFront distribution.
	// Experimental.
	EnableIpv6() *bool
}

// The jsii proxy for ICloudFrontConfiguration
type jsiiProxy_ICloudFrontConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_ICloudFrontConfiguration) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudFrontConfiguration) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudFrontConfiguration) EnableIpv6() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIpv6",
		&returns,
	)
	return returns
}

