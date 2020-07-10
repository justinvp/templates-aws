package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticsearch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogResourcePolicy(ctx, "exampleLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
			PolicyDocument: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"es.amazonaws.com\"\n", "      },\n", "      \"Action\": [\n", "        \"logs:PutLogEvents\",\n", "        \"logs:PutLogEventsBatch\",\n", "        \"logs:CreateLogStream\"\n", "      ],\n", "      \"Resource\": \"arn:aws:logs:*\"\n", "    }\n", "  ]\n", "}\n", "\n")),
			PolicyName:     pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = elasticsearch.NewDomain(ctx, "exampleDomain", &elasticsearch.DomainArgs{
			LogPublishingOptions: elasticsearch.DomainLogPublishingOptionArray{
				&elasticsearch.DomainLogPublishingOptionArgs{
					CloudwatchLogGroupArn: exampleLogGroup.Arn,
					LogType:               pulumi.String("INDEX_SLOW_LOGS"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

