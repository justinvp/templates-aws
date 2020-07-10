package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/batch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ecsInstanceRoleRole, err := iam.NewRole(ctx, "ecsInstanceRoleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "	{\n", "	    \"Action\": \"sts:AssumeRole\",\n", "	    \"Effect\": \"Allow\",\n", "	    \"Principal\": {\n", "		\"Service\": \"ec2.amazonaws.com\"\n", "	    }\n", "	}\n", "    ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "ecsInstanceRoleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"),
			Role:      ecsInstanceRoleRole.Name,
		})
		if err != nil {
			return err
		}
		ecsInstanceRoleInstanceProfile, err := iam.NewInstanceProfile(ctx, "ecsInstanceRoleInstanceProfile", &iam.InstanceProfileArgs{
			Role: ecsInstanceRoleRole.Name,
		})
		if err != nil {
			return err
		}
		awsBatchServiceRoleRole, err := iam.NewRole(ctx, "awsBatchServiceRoleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "	{\n", "	    \"Action\": \"sts:AssumeRole\",\n", "	    \"Effect\": \"Allow\",\n", "	    \"Principal\": {\n", "		\"Service\": \"batch.amazonaws.com\"\n", "	    }\n", "	}\n", "    ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "awsBatchServiceRoleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole"),
			Role:      awsBatchServiceRoleRole.Name,
		})
		if err != nil {
			return err
		}
		sampleSecurityGroup, err := ec2.NewSecurityGroup(ctx, "sampleSecurityGroup", &ec2.SecurityGroupArgs{
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
					FromPort: pulumi.Int(0),
					Protocol: pulumi.String("-1"),
					ToPort:   pulumi.Int(0),
				},
			},
		})
		if err != nil {
			return err
		}
		sampleVpc, err := ec2.NewVpc(ctx, "sampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		sampleSubnet, err := ec2.NewSubnet(ctx, "sampleSubnet", &ec2.SubnetArgs{
			CidrBlock: pulumi.String("10.1.1.0/24"),
			VpcId:     sampleVpc.ID(),
		})
		if err != nil {
			return err
		}
		_, err = batch.NewComputeEnvironment(ctx, "sampleComputeEnvironment", &batch.ComputeEnvironmentArgs{
			ComputeEnvironmentName: pulumi.String("sample"),
			ComputeResources: &batch.ComputeEnvironmentComputeResourcesArgs{
				InstanceRole: ecsInstanceRoleInstanceProfile.Arn,
				InstanceType: pulumi.StringArray{
					pulumi.String("c4.large"),
				},
				MaxVcpus: pulumi.Int(16),
				MinVcpus: pulumi.Int(0),
				SecurityGroupIds: pulumi.StringArray{
					sampleSecurityGroup.ID(),
				},
				Subnets: pulumi.StringArray{
					sampleSubnet.ID(),
				},
				Type: pulumi.String("EC2"),
			},
			ServiceRole: awsBatchServiceRoleRole.Arn,
			Type:        pulumi.String("MANAGED"),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_iam_role_policy_attachment.aws_batch_service_role",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

