package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/inspector"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := inspector.NewAssessmentTemplate(ctx, "example", &inspector.AssessmentTemplateArgs{
			Duration: pulumi.Int(3600),
			RulesPackageArns: pulumi.StringArray{
				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p"),
				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc"),
				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ"),
				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD"),
			},
			TargetArn: pulumi.Any(aws_inspector_assessment_target.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

