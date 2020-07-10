package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomainName, err := apigateway.NewDomainName(ctx, "exampleDomainName", &apigateway.DomainNameArgs{
			CertificateArn: pulumi.Any(aws_acm_certificate_validation.Example.Certificate_arn),
			DomainName:     pulumi.String("api.example.com"),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					EvaluateTargetHealth: pulumi.Bool(true),
					Name:                 exampleDomainName.CloudfrontDomainName,
					ZoneId:               exampleDomainName.CloudfrontZoneId,
				},
			},
			Name:   exampleDomainName.DomainName,
			Type:   pulumi.String("A"),
			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

