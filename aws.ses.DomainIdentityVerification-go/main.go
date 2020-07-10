package main

import (
	"fmt"

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
			Name: example.ID().ApplyT(func(id string) (string, error) {
				return fmt.Sprintf("%v%v", "_amazonses.", id), nil
			}).(pulumi.StringOutput),
			Records: pulumi.StringArray{
				example.VerificationToken,
			},
			Ttl:    pulumi.Int(600),
			Type:   pulumi.String("TXT"),
			ZoneId: pulumi.Any(aws_route53_zone.Example.Zone_id),
		})
		if err != nil {
			return err
		}
		_, err = ses.NewDomainIdentityVerification(ctx, "exampleVerification", &ses.DomainIdentityVerificationArgs{
			Domain: example.ID(),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_route53_record.example_amazonses_verification_record",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

