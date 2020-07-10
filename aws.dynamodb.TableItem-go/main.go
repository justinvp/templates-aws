package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleTable, err := dynamodb.NewTable(ctx, "exampleTable", &dynamodb.TableArgs{
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("exampleHashKey"),
					Type: pulumi.String("S"),
				},
			},
			HashKey:       pulumi.String("exampleHashKey"),
			ReadCapacity:  pulumi.Int(10),
			WriteCapacity: pulumi.Int(10),
		})
		if err != nil {
			return err
		}
		_, err = dynamodb.NewTableItem(ctx, "exampleTableItem", &dynamodb.TableItemArgs{
			HashKey:   exampleTable.HashKey,
			Item:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v", "{\n", "  \"exampleHashKey\": {\"S\": \"something\"},\n", "  \"one\": {\"N\": \"11111\"},\n", "  \"two\": {\"N\": \"22222\"},\n", "  \"three\": {\"N\": \"33333\"},\n", "  \"four\": {\"N\": \"44444\"}\n", "}\n", "\n")),
			TableName: exampleTable.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

