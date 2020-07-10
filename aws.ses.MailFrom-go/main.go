package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomainIdentity, err := ses.NewDomainIdentity(ctx, "exampleDomainIdentity", &ses.DomainIdentityArgs{
			Domain: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		exampleMailFrom, err := ses.NewMailFrom(ctx, "exampleMailFrom", &ses.MailFromArgs{
			Domain: exampleDomainIdentity.Domain,
			MailFromDomain: exampleDomainIdentity.Domain.ApplyT(func(domain string) (string, error) {
				return fmt.Sprintf("%v%v", "bounce.", domain), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleSesDomainMailFromMx", &route53.RecordArgs{
			Name: exampleMailFrom.MailFromDomain,
			Records: pulumi.StringArray{
				pulumi.String("10 feedback-smtp.us-east-1.amazonses.com"),
			},
			Ttl:    pulumi.Int(600),
			Type:   pulumi.String("MX"),
			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleSesDomainMailFromTxt", &route53.RecordArgs{
			Name: exampleMailFrom.MailFromDomain,
			Records: pulumi.StringArray{
				pulumi.String("v=spf1 include:amazonses.com -all"),
			},
			Ttl:    pulumi.Int(600),
			Type:   pulumi.String("TXT"),
			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

