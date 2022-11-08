// Domain with certificate
package almacdkdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Interface contract implemented by Domain construct.
// Experimental.
type IDomain interface {
	// Assign an alias as record target with the fully-qualified domain name.
	//
	// This will create both `A` & `AAAA` DNS records, unless `disableIpV6` was set to `true`
	// during initialization of `Domain` construct (resulting in only `A` record being created).
	//
	// Example:
	//   domain.addTarget(new targets.CloudFrontTarget(distribution))
	//
	// Experimental.
	AddTarget(alias awsroute53.IAliasRecordTarget)
	// Certificate Manager certificate.
	// Experimental.
	Certificate() awscertificatemanager.ICertificate
	// Has IPv6 AAAA records been created.
	//
	// Can be used to conditionally configure IPv6 support
	// to CloudFront distribution.
	// Experimental.
	EnableIpv6() *bool
	// Fully-qualified domain name.
	// Experimental.
	Fqdn() *string
	// Route53 hosted zone used to assign the domain into.
	// Experimental.
	Zone() awsroute53.IHostedZone
}

// The jsii proxy for IDomain
type jsiiProxy_IDomain struct {
	_ byte // padding
}

func (i *jsiiProxy_IDomain) AddTarget(alias awsroute53.IAliasRecordTarget) {
	if err := i.validateAddTargetParameters(alias); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addTarget",
		[]interface{}{alias},
	)
}

func (j *jsiiProxy_IDomain) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomain) EnableIpv6() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomain) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomain) Zone() awsroute53.IHostedZone {
	var returns awsroute53.IHostedZone
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

