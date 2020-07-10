package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "s3"
		s3, err := ec2.LookupVpcEndpointService(ctx, &ec2.LookupVpcEndpointServiceArgs{
			Service: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		foo, err := ec2.NewVpc(ctx, "foo", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcEndpoint(ctx, "ep", &ec2.VpcEndpointArgs{
			ServiceName: pulumi.String(s3.ServiceName),
			VpcId:       foo.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

