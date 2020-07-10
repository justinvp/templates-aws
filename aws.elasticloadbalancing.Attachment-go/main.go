package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elb.NewAttachment(ctx, "baz", &elb.AttachmentArgs{
			Elb:      pulumi.Any(aws_elb.Bar.Id),
			Instance: pulumi.Any(aws_instance.Foo.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

