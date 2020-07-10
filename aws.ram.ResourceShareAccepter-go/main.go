package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providers.Newaws(ctx, "alternate", &providers.awsArgs{
			Profile: pulumi.String("profile1"),
		})
		if err != nil {
			return err
		}
		senderShare, err := ram.NewResourceShare(ctx, "senderShare", &ram.ResourceShareArgs{
			AllowExternalPrincipals: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-test-resource-share"),
			},
		}, pulumi.Provider("aws.alternate"))
		if err != nil {
			return err
		}
		receiver, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		senderInvite, err := ram.NewPrincipalAssociation(ctx, "senderInvite", &ram.PrincipalAssociationArgs{
			Principal:        pulumi.String(receiver.AccountId),
			ResourceShareArn: senderShare.Arn,
		}, pulumi.Provider("aws.alternate"))
		if err != nil {
			return err
		}
		_, err = ram.NewResourceShareAccepter(ctx, "receiverAccept", &ram.ResourceShareAccepterArgs{
			ShareArn: senderInvite.ResourceShareArn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

