package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bootstrapScript, err := s3.LookupBucketObject(ctx, &s3.LookupBucketObjectArgs{
			Bucket: "ourcorp-deploy-config",
			Key:    "ec2-bootstrap-script.sh",
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewInstance(ctx, "example", &ec2.InstanceArgs{
			Ami:          pulumi.String("ami-2757f631"),
			InstanceType: pulumi.String("t2.micro"),
			UserData:     pulumi.String(bootstrapScript.Body),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

