package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testStaticIp, err := lightsail.NewStaticIp(ctx, "testStaticIp", nil)
		if err != nil {
			return err
		}
		testInstance, err := lightsail.NewInstance(ctx, "testInstance", &lightsail.InstanceArgs{
			AvailabilityZone: pulumi.String("us-east-1b"),
			BlueprintId:      pulumi.String("string"),
			BundleId:         pulumi.String("string"),
			KeyPairName:      pulumi.String("some_key_name"),
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewStaticIpAttachment(ctx, "testStaticIpAttachment", &lightsail.StaticIpAttachmentArgs{
			InstanceName: testInstance.ID(),
			StaticIpName: testStaticIp.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

