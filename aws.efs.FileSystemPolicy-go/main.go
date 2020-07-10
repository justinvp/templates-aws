package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fs, err := efs.NewFileSystem(ctx, "fs", nil)
		if err != nil {
			return err
		}
		_, err = efs.NewFileSystemPolicy(ctx, "policy", &efs.FileSystemPolicyArgs{
			FileSystemId: fs.ID(),
			Policy:       pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Id\": \"ExamplePolicy01\",\n", "    \"Statement\": [\n", "        {\n", "            \"Sid\": \"ExampleSatement01\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": {\n", "                \"AWS\": \"*\"\n", "            },\n", "            \"Resource\": \"", aws_efs_file_system.Test.Arn, "\",\n", "            \"Action\": [\n", "                \"elasticfilesystem:ClientMount\",\n", "                \"elasticfilesystem:ClientWrite\"\n", "            ],\n", "            \"Condition\": {\n", "                \"Bool\": {\n", "                    \"aws:SecureTransport\": \"true\"\n", "                }\n", "            }\n", "        }\n", "    ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

