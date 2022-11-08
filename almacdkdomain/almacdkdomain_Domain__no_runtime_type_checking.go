//go:build no_runtime_type_checking

// Domain with certificate
package almacdkdomain

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Domain) validateAddTargetParameters(alias awsroute53.IAliasRecordTarget) error {
	return nil
}

func validateDomain_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDomainParameters(scope constructs.Construct, id *string, props *DomainProps) error {
	return nil
}

