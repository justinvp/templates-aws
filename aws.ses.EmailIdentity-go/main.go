package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewEmailIdentity(ctx, "example", &ses.EmailIdentityArgs{
			Email: pulumi.String("email@example.com"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

