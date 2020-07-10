package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := ebs.LookupDefaultKmsKey(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Encrypted:        pulumi.Bool(true),
			KmsKeyId:         pulumi.String(current.KeyArn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

