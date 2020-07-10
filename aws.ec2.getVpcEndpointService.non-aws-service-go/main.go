package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "com.amazonaws.vpce.us-west-2.vpce-svc-0e87519c997c63cd8"
		_, err := ec2.LookupVpcEndpointService(ctx, &ec2.LookupVpcEndpointServiceArgs{
			ServiceName: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

