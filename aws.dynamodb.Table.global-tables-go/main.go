package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dynamodb.NewTable(ctx, "example", &dynamodb.TableArgs{
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("TestTableHashKey"),
					Type: pulumi.String("S"),
				},
			},
			BillingMode: pulumi.String("PAY_PER_REQUEST"),
			HashKey:     pulumi.String("TestTableHashKey"),
			Replicas: dynamodb.TableReplicaArray{
				&dynamodb.TableReplicaArgs{
					RegionName: pulumi.String("us-east-2"),
				},
				&dynamodb.TableReplicaArgs{
					RegionName: pulumi.String("us-west-2"),
				},
			},
			StreamEnabled:  pulumi.Bool(true),
			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

