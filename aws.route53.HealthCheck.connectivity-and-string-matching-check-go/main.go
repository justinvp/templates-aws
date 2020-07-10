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
			Port:             pulumi.Int(443),
			RequestInterval:  pulumi.Int(30),
			ResourcePath:     pulumi.String("/"),
			SearchString:     pulumi.String("example"),
			Type:             pulumi.String("HTTPS_STR_MATCH"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

