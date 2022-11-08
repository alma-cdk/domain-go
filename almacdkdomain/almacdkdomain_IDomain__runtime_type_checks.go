//go:build !no_runtime_type_checking

// Domain with certificate
package almacdkdomain

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

func (i *jsiiProxy_IDomain) validateAddTargetParameters(alias awsroute53.IAliasRecordTarget) error {
	if alias == nil {
		return fmt.Errorf("parameter alias is required, but nil was provided")
	}

	return nil
}

