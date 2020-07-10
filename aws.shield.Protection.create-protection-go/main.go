package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/shield"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetAvailabilityZones(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		fooEip, err := ec2.NewEip(ctx, "fooEip", &ec2.EipArgs{
			Vpc: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = shield.NewProtection(ctx, "fooProtection", &shield.ProtectionArgs{
			ResourceArn: fooEip.ID().ApplyT(func(id string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v", "arn:aws:ec2:", currentRegion.Name, ":", currentCallerIdentity.AccountId, ":eip-allocation/", id), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

