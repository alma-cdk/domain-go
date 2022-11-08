package almacdkdomain

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@alma-cdk/domain.Domain",
		reflect.TypeOf((*Domain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTarget", GoMethod: "AddTarget"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberMethod{JsiiMethod: "configureCloudFront", GoMethod: "ConfigureCloudFront"},
			_jsii_.MemberProperty{JsiiProperty: "enableIpv6", GoGetter: "EnableIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "fqdn", GoGetter: "Fqdn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "zone", GoGetter: "Zone"},
		},
		func() interface{} {
			j := jsiiProxy_Domain{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDomain)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/domain.DomainProps",
		reflect.TypeOf((*DomainProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@alma-cdk/domain.ICloudFrontConfiguration",
		reflect.TypeOf((*ICloudFrontConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "domainNames", GoGetter: "DomainNames"},
			_jsii_.MemberProperty{JsiiProperty: "enableIpv6", GoGetter: "EnableIpv6"},
		},
		func() interface{} {
			return &jsiiProxy_ICloudFrontConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"@alma-cdk/domain.IDomain",
		reflect.TypeOf((*IDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTarget", GoMethod: "AddTarget"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "enableIpv6", GoGetter: "EnableIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "fqdn", GoGetter: "Fqdn"},
			_jsii_.MemberProperty{JsiiProperty: "zone", GoGetter: "Zone"},
		},
		func() interface{} {
			return &jsiiProxy_IDomain{}
		},
	)
}
