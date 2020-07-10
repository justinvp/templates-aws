package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := organizations.LookupOrganization(ctx, nil, nil)
		if err != nil {
			return err
		}
		snsTopic, err := sns.NewTopic(ctx, "snsTopic", nil)
		if err != nil {
			return err
		}
		_, err = sns.NewTopicPolicy(ctx, "snsTopicPolicyTopicPolicy", &sns.TopicPolicyArgs{
			Arn:    snsTopic.Arn,
			Policy: pulumi.String(snsTopicPolicyPolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

