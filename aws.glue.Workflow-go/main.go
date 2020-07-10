package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := glue.NewWorkflow(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = glue.NewTrigger(ctx, "example_start", &glue.TriggerArgs{
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.String("example-job"),
				},
			},
			Type:         pulumi.String("ON_DEMAND"),
			WorkflowName: example.Name,
		})
		if err != nil {
			return err
		}
		_, err = glue.NewTrigger(ctx, "example_inner", &glue.TriggerArgs{
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.String("another-example-job"),
				},
			},
			Predicate: &glue.TriggerPredicateArgs{
				Conditions: glue.TriggerPredicateConditionArray{
					&glue.TriggerPredicateConditionArgs{
						JobName: pulumi.String("example-job"),
						State:   pulumi.String("SUCCEEDED"),
					},
				},
			},
			Type:         pulumi.String("CONDITIONAL"),
			WorkflowName: example.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

