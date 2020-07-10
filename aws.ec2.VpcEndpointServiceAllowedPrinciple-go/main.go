package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcEndpointServiceAllowedPrinciple(ctx, "allowMeToFoo", &ec2.VpcEndpointServiceAllowedPrincipleArgs{
			PrincipalArn:         pulumi.String(current.Arn),
			VpcEndpointServiceId: pulumi.Any(aws_vpc_endpoint_service.Foo.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

