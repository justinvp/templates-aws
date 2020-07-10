package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dynamodb.NewTable(ctx, "basic_dynamodb_table", &dynamodb.TableArgs{
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("UserId"),
					Type: pulumi.String("S"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("GameTitle"),
					Type: pulumi.String("S"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("TopScore"),
					Type: pulumi.String("N"),
				},
			},
			BillingMode: pulumi.String("PROVISIONED"),
			GlobalSecondaryIndexes: dynamodb.TableGlobalSecondaryIndexArray{
				&dynamodb.TableGlobalSecondaryIndexArgs{
					HashKey: pulumi.String("GameTitle"),
					Name:    pulumi.String("GameTitleIndex"),
					NonKeyAttributes: pulumi.StringArray{
						pulumi.String("UserId"),
					},
					ProjectionType: pulumi.String("INCLUDE"),
					RangeKey:       pulumi.String("TopScore"),
					ReadCapacity:   pulumi.Int(10),
					WriteCapacity:  pulumi.Int(10),
				},
			},
			HashKey:      pulumi.String("UserId"),
			RangeKey:     pulumi.String("GameTitle"),
			ReadCapacity: pulumi.Int(20),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
				"Name":        pulumi.String("dynamodb-table-1"),
			},
			Ttl: &dynamodb.TableTtlArgs{
				AttributeName: pulumi.String("TimeToExist"),
				Enabled:       pulumi.Bool(false),
			},
			WriteCapacity: pulumi.Int(20),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

