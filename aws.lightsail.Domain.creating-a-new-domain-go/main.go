package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewDomain(ctx, "domainTest", &lightsail.DomainArgs{
			DomainName: pulumi.String("mydomain.com"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

