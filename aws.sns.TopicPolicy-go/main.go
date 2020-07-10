package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := sns.NewTopic(ctx, "test", nil)
		if err != nil {
			return err
		}
		_, err = sns.NewTopicPolicy(ctx, "_default", &sns.TopicPolicyArgs{
			Arn: test.Arn,
			Policy: snsTopicPolicy.ApplyT(func(snsTopicPolicy iam.GetPolicyDocumentResult) (string, error) {
				return snsTopicPolicy.Json, nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

