package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudformation"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		subnetId, err := cloudformation.GetExport(ctx, &cloudformation.GetExportArgs{
			Name: "mySubnetIdExportName",
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
			Ami:          pulumi.String("ami-abb07bcb"),
			InstanceType: pulumi.String("t1.micro"),
			SubnetId:     pulumi.String(subnetId.Value),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

