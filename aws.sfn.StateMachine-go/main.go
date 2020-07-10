package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sfn.NewStateMachine(ctx, "sfnStateMachine", &sfn.StateMachineArgs{
			Definition: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Comment\": \"A Hello World example of the Amazon States Language using an AWS Lambda Function\",\n", "  \"StartAt\": \"HelloWorld\",\n", "  \"States\": {\n", "    \"HelloWorld\": {\n", "      \"Type\": \"Task\",\n", "      \"Resource\": \"", aws_lambda_function.Lambda.Arn, "\",\n", "      \"End\": true\n", "    }\n", "  }\n", "}\n", "\n")),
			RoleArn:    pulumi.Any(aws_iam_role.Iam_for_sfn.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

