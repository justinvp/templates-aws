package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewRdsDbInstance(ctx, "myInstance", &opsworks.RdsDbInstanceArgs{
			DbPassword:       pulumi.String("somePass"),
			DbUser:           pulumi.String("someUser"),
			RdsDbInstanceArn: pulumi.Any(aws_db_instance.My_instance.Arn),
			StackId:          pulumi.Any(aws_opsworks_stack.My_stack.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

