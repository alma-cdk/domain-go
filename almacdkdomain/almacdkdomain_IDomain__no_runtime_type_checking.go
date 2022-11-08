//go:build no_runtime_type_checking

// Domain with certificate
package almacdkdomain

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDomain) validateAddTargetParameters(alias awsroute53.IAliasRecordTarget) error {
	return nil
}

