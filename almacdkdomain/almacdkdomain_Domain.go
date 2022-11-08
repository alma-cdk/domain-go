// Domain with certificate
package almacdkdomain

import (
	_init_ "github.com/alma-cdk/domain-go/almacdkdomain/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alma-cdk/domain-go/almacdkdomain/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type Domain interface {
	constructs.Construct
	IDomain
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Route53 hosted zone used to assign the domain into.
	// Experimental.
	Zone() awsroute53.IHostedZone
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
	// Helper method to configure CloudFront distribution with the domain, certificate and IPv6 support.
	//
	// Returns: CloudFront configuration for certificate, domainNames and IPv6.
	// Experimental.
	ConfigureCloudFront() ICloudFrontConfiguration
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Domain
type jsiiProxy_Domain struct {
	internal.Type__constructsConstruct
	jsiiProxy_IDomain
}

func (j *jsiiProxy_Domain) Certificate() awscertificatemanager.ICertificate {
	var returns awscertificatemanager.ICertificate
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) EnableIpv6() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Zone() awsroute53.IHostedZone {
	var returns awsroute53.IHostedZone
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}


// Initializing a `new Domain` construct instance will lookup the Route53 hosted zone and define ACM DNS-validated certificate.
//
// After initialization you must use `assign(alias)` method to to configure `A`/`AAAA` records
// with the `alias` as the record value.
// Experimental.
func NewDomain(scope constructs.Construct, id *string, props *DomainProps) Domain {
	_init_.Initialize()

	if err := validateNewDomainParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Domain{}

	_jsii_.Create(
		"@alma-cdk/domain.Domain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Initializing a `new Domain` construct instance will lookup the Route53 hosted zone and define ACM DNS-validated certificate.
//
// After initialization you must use `assign(alias)` method to to configure `A`/`AAAA` records
// with the `alias` as the record value.
// Experimental.
func NewDomain_Override(d Domain, scope constructs.Construct, id *string, props *DomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/domain.Domain",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Domain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/domain.Domain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) AddTarget(alias awsroute53.IAliasRecordTarget) {
	if err := d.validateAddTargetParameters(alias); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addTarget",
		[]interface{}{alias},
	)
}

func (d *jsiiProxy_Domain) ConfigureCloudFront() ICloudFrontConfiguration {
	var returns ICloudFrontConfiguration

	_jsii_.Invoke(
		d,
		"configureCloudFront",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

