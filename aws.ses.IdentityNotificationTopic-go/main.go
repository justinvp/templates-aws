package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewIdentityNotificationTopic(ctx, "test", &ses.IdentityNotificationTopicArgs{
			Identity:               pulumi.Any(aws_ses_domain_identity.Example.Domain),
			IncludeOriginalHeaders: pulumi.Bool(true),
			NotificationType:       pulumi.String("Bounce"),
			TopicArn:               pulumi.Any(aws_sns_topic.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

