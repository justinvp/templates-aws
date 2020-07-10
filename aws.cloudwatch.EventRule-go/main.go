package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		console, err := cloudwatch.NewEventRule(ctx, "console", &cloudwatch.EventRuleArgs{
			Description:  pulumi.String("Capture each AWS Console Sign In"),
			EventPattern: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v", "{\n", "  \"detail-type\": [\n", "    \"AWS Console Sign In via CloudTrail\"\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		awsLogins, err := sns.NewTopic(ctx, "awsLogins", nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "sns", &cloudwatch.EventTargetArgs{
			Arn:  awsLogins.Arn,
			Rule: console.Name,
		})
		if err != nil {
			return err
		}
		_, err = sns.NewTopicPolicy(ctx, "_default", &sns.TopicPolicyArgs{
			Arn: awsLogins.Arn,
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

