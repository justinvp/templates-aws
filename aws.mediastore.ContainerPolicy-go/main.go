package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/mediastore"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleContainer, err := mediastore.NewContainer(ctx, "exampleContainer", nil)
		if err != nil {
			return err
		}
		_, err = mediastore.NewContainerPolicy(ctx, "exampleContainerPolicy", &mediastore.ContainerPolicyArgs{
			ContainerName: exampleContainer.Name,
			Policy: exampleContainer.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"Version\": \"2012-10-17\",\n", "	\"Statement\": [{\n", "		\"Sid\": \"MediaStoreFullAccess\",\n", "		\"Action\": [ \"mediastore:*\" ],\n", "		\"Principal\": {\"AWS\" : \"arn:aws:iam::", currentCallerIdentity.AccountId, ":root\"},\n", "		\"Effect\": \"Allow\",\n", "		\"Resource\": \"arn:aws:mediastore:", currentCallerIdentity.AccountId, ":", currentRegion.Name, ":container/", name, "/*\",\n", "		\"Condition\": {\n", "			\"Bool\": { \"aws:SecureTransport\": \"true\" }\n", "		}\n", "	}]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

