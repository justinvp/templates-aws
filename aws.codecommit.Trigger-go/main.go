package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codecommit"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testRepository, err := codecommit.NewRepository(ctx, "testRepository", &codecommit.RepositoryArgs{
			RepositoryName: pulumi.String("test"),
		})
		if err != nil {
			return err
		}
		_, err = codecommit.NewTrigger(ctx, "testTrigger", &codecommit.TriggerArgs{
			RepositoryName: testRepository.RepositoryName,
			Triggers: codecommit.TriggerTriggerArray{
				&codecommit.TriggerTriggerArgs{
					DestinationArn: pulumi.Any(aws_sns_topic.Test.Arn),
					Events: pulumi.StringArray{
						pulumi.String("all"),
					},
					Name: pulumi.String("all"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

