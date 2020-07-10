package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appsync"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "        \"Effect\": \"Allow\",\n", "        \"Principal\": {\n", "            \"Service\": \"appsync.amazonaws.com\"\n", "        },\n", "        \"Action\": \"sts:AssumeRole\"\n", "        }\n", "    ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSAppSyncPushToCloudWatchLogs"),
			Role:      exampleRole.Name,
		})
		if err != nil {
			return err
		}
		_, err = appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
			LogConfig: &appsync.GraphQLApiLogConfigArgs{
				CloudwatchLogsRoleArn: exampleRole.Arn,
				FieldLogLevel:         pulumi.String("ERROR"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

