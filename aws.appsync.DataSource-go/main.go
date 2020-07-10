package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appsync"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleTable, err := dynamodb.NewTable(ctx, "exampleTable", &dynamodb.TableArgs{
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("UserId"),
					Type: pulumi.String("S"),
				},
			},
			HashKey:       pulumi.String("UserId"),
			ReadCapacity:  pulumi.Int(1),
			WriteCapacity: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"appsync.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "exampleRolePolicy", &iam.RolePolicyArgs{
			Policy: exampleTable.Arn.ApplyT(func(arn string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"dynamodb:*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": [\n", "        \"", arn, "\"\n", "      ]\n", "    }\n", "  ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
			Role: exampleRole.ID(),
		})
		if err != nil {
			return err
		}
		exampleGraphQLApi, err := appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
			AuthenticationType: pulumi.String("API_KEY"),
		})
		if err != nil {
			return err
		}
		_, err = appsync.NewDataSource(ctx, "exampleDataSource", &appsync.DataSourceArgs{
			ApiId: exampleGraphQLApi.ID(),
			DynamodbConfig: &appsync.DataSourceDynamodbConfigArgs{
				TableName: exampleTable.Name,
			},
			ServiceRoleArn: exampleRole.Arn,
			Type:           pulumi.String("AMAZON_DYNAMODB"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

