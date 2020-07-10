package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glacier"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsSnsTopic, err := sns.NewTopic(ctx, "awsSnsTopic", nil)
		if err != nil {
			return err
		}
		_, err = glacier.NewVault(ctx, "myArchive", &glacier.VaultArgs{
			AccessPolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\":\"2012-10-17\",\n", "    \"Statement\":[\n", "       {\n", "          \"Sid\": \"add-read-only-perm\",\n", "          \"Principal\": \"*\",\n", "          \"Effect\": \"Allow\",\n", "          \"Action\": [\n", "             \"glacier:InitiateJob\",\n", "             \"glacier:GetJobOutput\"\n", "          ],\n", "          \"Resource\": \"arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive\"\n", "       }\n", "    ]\n", "}\n", "\n")),
			Notifications: glacier.VaultNotificationArray{
				&glacier.VaultNotificationArgs{
					Events: pulumi.StringArray{
						pulumi.String("ArchiveRetrievalCompleted"),
						pulumi.String("InventoryRetrievalCompleted"),
					},
					SnsTopic: awsSnsTopic.Arn,
				},
			},
			Tags: pulumi.StringMap{
				"Test": pulumi.String("MyArchive"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

