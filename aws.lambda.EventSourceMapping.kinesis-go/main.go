package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
			EventSourceArn:   pulumi.Any(aws_kinesis_stream.Example.Arn),
			FunctionName:     pulumi.Any(aws_lambda_function.Example.Arn),
			StartingPosition: pulumi.String("LATEST"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

