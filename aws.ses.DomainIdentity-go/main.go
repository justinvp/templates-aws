package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := ses.NewDomainIdentity(ctx, "example", &ses.DomainIdentityArgs{
			Domain: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleAmazonsesVerificationRecord", &route53.RecordArgs{
			Name: pulumi.String("_amazonses.example.com"),
			Records: pulumi.StringArray{
				example.VerificationToken,
			},
			Ttl:    pulumi.Int(600),
			Type:   pulumi.String("TXT"),
			ZoneId: pulumi.String("ABCDEFGHIJ123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

