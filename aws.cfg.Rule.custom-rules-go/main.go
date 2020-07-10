package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cfg.NewRecorder(ctx, "exampleRecorder", nil)
		if err != nil {
			return err
		}
		exampleFunction, err := lambda.NewFunction(ctx, "exampleFunction", nil)
		if err != nil {
			return err
		}
		_, err = lambda.NewPermission(ctx, "examplePermission", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  exampleFunction.Arn,
			Principal: pulumi.String("config.amazonaws.com"),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewRule(ctx, "exampleRule", &cfg.RuleArgs{
			Source: &cfg.RuleSourceArgs{
				Owner:            pulumi.String("CUSTOM_LAMBDA"),
				SourceIdentifier: exampleFunction.Arn,
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_config_configuration_recorder.example",
			"aws_lambda_permission.example",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

