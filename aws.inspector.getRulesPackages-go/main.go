package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/inspector"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		rules, err := inspector.GetRulesPackages(ctx, nil, nil)
		if err != nil {
			return err
		}
		group, err := inspector.NewResourceGroup(ctx, "group", &inspector.ResourceGroupArgs{
			Tags: pulumi.StringMap{
				"test": pulumi.String("test"),
			},
		})
		if err != nil {
			return err
		}
		assessmentAssessmentTarget, err := inspector.NewAssessmentTarget(ctx, "assessmentAssessmentTarget", &inspector.AssessmentTargetArgs{
			ResourceGroupArn: group.Arn,
		})
		if err != nil {
			return err
		}
		_, err = inspector.NewAssessmentTemplate(ctx, "assessmentAssessmentTemplate", &inspector.AssessmentTemplateArgs{
			Duration:         pulumi.Int(60),
			RulesPackageArns: toPulumiStringArray(rules.Arns),
			TargetArn:        assessmentAssessmentTarget.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
func toPulumiStringArray(arr []string) pulumi.StringArray {
	var pulumiArr pulumi.StringArray
	for _, v := range arr {
		pulumiArr = append(pulumiArr, pulumi.String(v))
	}
	return pulumiArr
}

