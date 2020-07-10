package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewAmiCopy(ctx, "example", &ec2.AmiCopyArgs{
			Description:     pulumi.String("A copy of ami-xxxxxxxx"),
			SourceAmiId:     pulumi.String("ami-xxxxxxxx"),
			SourceAmiRegion: pulumi.String("us-west-1"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

