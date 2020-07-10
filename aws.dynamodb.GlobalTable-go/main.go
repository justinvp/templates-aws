package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providers.Newaws(ctx, "us_east_1", &providers.awsArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		_, err = providers.Newaws(ctx, "us_west_2", &providers.awsArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		_, err = dynamodb.NewTable(ctx, "us_east_1Table", &dynamodb.TableArgs{
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("myAttribute"),
					Type: pulumi.String("S"),
				},
			},
			HashKey:        pulumi.String("myAttribute"),
			ReadCapacity:   pulumi.Int(1),
			StreamEnabled:  pulumi.Bool(true),
			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
			WriteCapacity:  pulumi.Int(1),
		}, pulumi.Provider("aws.us-east-1"))
		if err != nil {
			return err
		}
		_, err = dynamodb.NewTable(ctx, "us_west_2Table", &dynamodb.TableArgs{
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("myAttribute"),
					Type: pulumi.String("S"),
				},
			},
			HashKey:        pulumi.String("myAttribute"),
			ReadCapacity:   pulumi.Int(1),
			StreamEnabled:  pulumi.Bool(true),
			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
			WriteCapacity:  pulumi.Int(1),
		}, pulumi.Provider("aws.us-west-2"))
		if err != nil {
			return err
		}
		_, err = dynamodb.NewGlobalTable(ctx, "myTable", &dynamodb.GlobalTableArgs{
			Replicas: dynamodb.GlobalTableReplicaArray{
				&dynamodb.GlobalTableReplicaArgs{
					RegionName: pulumi.String("us-east-1"),
				},
				&dynamodb.GlobalTableReplicaArgs{
					RegionName: pulumi.String("us-west-2"),
				},
			},
		}, pulumi.Provider("aws.us-east-1"), pulumi.DependsOn([]pulumi.Resource{
			"aws_dynamodb_table.us-east-1",
			"aws_dynamodb_table.us-west-2",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

