package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewWebAcl(ctx, "example", &waf.WebAclArgs{
			LoggingConfiguration: &waf.WebAclLoggingConfigurationArgs{
				LogDestination: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
				RedactedFields: &waf.WebAclLoggingConfigurationRedactedFieldsArgs{
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

