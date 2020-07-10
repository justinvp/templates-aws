package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providers.Newaws(ctx, "central", &providers.awsArgs{
			Region: pulumi.String("eu-central-1"),
		})
		if err != nil {
			return err
		}
		replicationRole, err := iam.NewRole(ctx, "replicationRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"s3.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		destination, err := s3.NewBucket(ctx, "destination", &s3.BucketArgs{
			Region: pulumi.String("eu-west-1"),
			Versioning: &s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Acl:    pulumi.String("private"),
			Region: pulumi.String("eu-central-1"),
			ReplicationConfiguration: &s3.BucketReplicationConfigurationArgs{
				Role: replicationRole.Arn,
				Rules: s3.BucketReplicationConfigurationRuleArray{
					&s3.BucketReplicationConfigurationRuleArgs{
						Destination: &s3.BucketReplicationConfigurationRuleDestinationArgs{
							Bucket:       destination.Arn,
							StorageClass: pulumi.String("STANDARD"),
						},
						Id:     pulumi.String("foobar"),
						Prefix: pulumi.String("foo"),
						Status: pulumi.String("Enabled"),
					},
				},
			},
			Versioning: &s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(true),
			},
		}, pulumi.Provider("aws.central"))
		if err != nil {
			return err
		}
		replicationPolicy, err := iam.NewPolicy(ctx, "replicationPolicy", &iam.PolicyArgs{
			Policy: pulumi.All(bucket.Arn, bucket.Arn, destination.Arn).ApplyT(func(_args []interface{}) (string, error) {
				bucketArn := _args[0].(string)
				bucketArn1 := _args[1].(string)
				destinationArn := _args[2].(string)
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"s3:GetReplicationConfiguration\",\n", "        \"s3:ListBucket\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": [\n", "        \"", bucketArn, "\"\n", "      ]\n", "    },\n", "    {\n", "      \"Action\": [\n", "        \"s3:GetObjectVersion\",\n", "        \"s3:GetObjectVersionAcl\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": [\n", "        \"", bucketArn1, "/*\"\n", "      ]\n", "    },\n", "    {\n", "      \"Action\": [\n", "        \"s3:ReplicateObject\",\n", "        \"s3:ReplicateDelete\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"", destinationArn, "/*\"\n", "    }\n", "  ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "replicationRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: replicationPolicy.Arn,
			Role:      replicationRole.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

