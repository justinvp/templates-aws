package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/acm"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCertificate, err := acm.NewCertificate(ctx, "exampleCertificate", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
		if err != nil {
			return err
		}
		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewListenerCertificate(ctx, "exampleListenerCertificate", &lb.ListenerCertificateArgs{
			CertificateArn: exampleCertificate.Arn,
			ListenerArn:    frontEndListener.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

