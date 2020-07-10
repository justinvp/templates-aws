package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "com.amazonaws.us-west-2.s3"
		opt1 := aws_vpc.Foo.Id
		s3, err := ec2.LookupVpcEndpoint(ctx, &ec2.LookupVpcEndpointArgs{
			ServiceName: &opt0,
			VpcId:       &opt1,
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcEndpointRouteTableAssociation(ctx, "privateS3", &ec2.VpcEndpointRouteTableAssociationArgs{
			RouteTableId:  pulumi.Any(aws_route_table.Private.Id),
			VpcEndpointId: pulumi.String(s3.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

