package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dms"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dmsAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				iam.GetPolicyDocumentStatement{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						iam.GetPolicyDocumentStatementPrincipal{
							Identifiers: []string{
								"dms.amazonaws.com",
							},
							Type: "Service",
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRole(ctx, "dms_access_for_endpoint", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(dmsAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "dms_access_for_endpoint_AmazonDMSRedshiftS3Role", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role"),
			Role:      dms_access_for_endpoint.Name,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRole(ctx, "dms_cloudwatch_logs_role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(dmsAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole"),
			Role:      dms_cloudwatch_logs_role.Name,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRole(ctx, "dms_vpc_role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(dmsAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "dms_vpc_role_AmazonDMSVPCManagementRole", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole"),
			Role:      dms_vpc_role.Name,
		})
		if err != nil {
			return err
		}
		_, err = dms.NewReplicationInstance(ctx, "test", &dms.ReplicationInstanceArgs{
			AllocatedStorage:           pulumi.Int(20),
			ApplyImmediately:           pulumi.Bool(true),
			AutoMinorVersionUpgrade:    pulumi.Bool(true),
			AvailabilityZone:           pulumi.String("us-west-2c"),
			EngineVersion:              pulumi.String("3.1.4"),
			KmsKeyArn:                  pulumi.String("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012"),
			MultiAz:                    pulumi.Bool(false),
			PreferredMaintenanceWindow: pulumi.String("sun:10:30-sun:14:30"),
			PubliclyAccessible:         pulumi.Bool(true),
			ReplicationInstanceClass:   pulumi.String("dms.t2.micro"),
			ReplicationInstanceId:      pulumi.String("test-dms-replication-instance-tf"),
			ReplicationSubnetGroupId:   pulumi.Any(aws_dms_replication_subnet_group.Test - dms - replication - subnet - group - tf.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test"),
			},
			VpcSecurityGroupIds: pulumi.StringArray{
				pulumi.String("sg-12345678"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

