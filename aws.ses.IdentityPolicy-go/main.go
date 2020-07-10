package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
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
		_, err = ses.NewIdentityPolicy(ctx, "exampleIdentityPolicy", &ses.IdentityPolicyArgs{
			Identity: exampleDomainIdentity.Arn,
			Policy: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (string, error) {
				return examplePolicyDocument.Json, nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

