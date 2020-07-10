package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewHealthCheck(ctx, "example", &route53.HealthCheckArgs{
			FailureThreshold: pulumi.Int(5),
			Fqdn:             pulumi.String("example.com"),
			Port:             pulumi.Int(80),
			RequestInterval:  pulumi.Int(30),
			ResourcePath:     pulumi.String("/"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-test-health-check"),
			},
			Type: pulumi.String("HTTP"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

