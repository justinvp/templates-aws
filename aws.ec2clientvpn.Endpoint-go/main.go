package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2clientvpn"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2clientvpn.NewEndpoint(ctx, "example", &ec2clientvpn.EndpointArgs{
			AuthenticationOptions: ec2clientvpn.EndpointAuthenticationOptionArray{
				&ec2clientvpn.EndpointAuthenticationOptionArgs{
					RootCertificateChainArn: pulumi.Any(aws_acm_certificate.Root_cert.Arn),
					Type:                    pulumi.String("certificate-authentication"),
				},
			},
			ClientCidrBlock: pulumi.String("10.0.0.0/16"),
			ConnectionLogOptions: &ec2clientvpn.EndpointConnectionLogOptionsArgs{
				CloudwatchLogGroup:  pulumi.Any(aws_cloudwatch_log_group.Lg.Name),
				CloudwatchLogStream: pulumi.Any(aws_cloudwatch_log_stream.Ls.Name),
				Enabled:             pulumi.Bool(true),
			},
			Description:          pulumi.String("clientvpn-example"),
			ServerCertificateArn: pulumi.Any(aws_acm_certificate.Cert.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

