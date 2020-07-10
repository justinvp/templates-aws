package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.Any(aws_glue_job.Example1.Name),
				},
			},
			Predicate: &glue.TriggerPredicateArgs{
				Conditions: glue.TriggerPredicateConditionArray{
					&glue.TriggerPredicateConditionArgs{
						CrawlState:  pulumi.String("SUCCEEDED"),
						CrawlerName: pulumi.Any(aws_glue_crawler.Example2.Name),
					},
				},
			},
			Type: pulumi.String("CONDITIONAL"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

