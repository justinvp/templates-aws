package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewServiceLinkedRole(ctx, "elasticbeanstalk", &iam.ServiceLinkedRoleArgs{
			AwsServiceName: pulumi.String("elasticbeanstalk.amazonaws.com"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

