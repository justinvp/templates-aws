package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
			EventSourceArn: pulumi.Any(aws_sqs_queue.Sqs_queue_test.Arn),
			FunctionName:   pulumi.Any(aws_lambda_function.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

