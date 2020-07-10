package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.NewWebAcl(ctx, "example", &wafregional.WebAclArgs{
			LoggingConfiguration: &wafregional.WebAclLoggingConfigurationArgs{
				LogDestination: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
				RedactedFields: &wafregional.WebAclLoggingConfigurationRedactedFieldsArgs{
					FieldToMatch: pulumi.Array{
						pulumi.StringMap{
							"type": pulumi.String("URI"),
						},
						pulumi.StringMap{
							"data": pulumi.String("referer"),
							"type": pulumi.String("HEADER"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

