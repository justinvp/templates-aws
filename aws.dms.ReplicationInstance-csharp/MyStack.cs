using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dmsAssumeRole = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "sts:AssumeRole",
                    },
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "dms.amazonaws.com",
                            },
                            Type = "Service",
                        },
                    },
                },
            },
        }));
        var dms_access_for_endpoint = new Aws.Iam.Role("dms-access-for-endpoint", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = dmsAssumeRole.Apply(dmsAssumeRole => dmsAssumeRole.Json),
        });
        var dms_access_for_endpoint_AmazonDMSRedshiftS3Role = new Aws.Iam.RolePolicyAttachment("dms-access-for-endpoint-AmazonDMSRedshiftS3Role", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role",
            Role = dms_access_for_endpoint.Name,
        });
        var dms_cloudwatch_logs_role = new Aws.Iam.Role("dms-cloudwatch-logs-role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = dmsAssumeRole.Apply(dmsAssumeRole => dmsAssumeRole.Json),
        });
        var dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole = new Aws.Iam.RolePolicyAttachment("dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole",
            Role = dms_cloudwatch_logs_role.Name,
        });
        var dms_vpc_role = new Aws.Iam.Role("dms-vpc-role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = dmsAssumeRole.Apply(dmsAssumeRole => dmsAssumeRole.Json),
        });
        var dms_vpc_role_AmazonDMSVPCManagementRole = new Aws.Iam.RolePolicyAttachment("dms-vpc-role-AmazonDMSVPCManagementRole", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole",
            Role = dms_vpc_role.Name,
        });
        // Create a new replication instance
        var test = new Aws.Dms.ReplicationInstance("test", new Aws.Dms.ReplicationInstanceArgs
        {
            AllocatedStorage = 20,
            ApplyImmediately = true,
            AutoMinorVersionUpgrade = true,
            AvailabilityZone = "us-west-2c",
            EngineVersion = "3.1.4",
            KmsKeyArn = "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
            MultiAz = false,
            PreferredMaintenanceWindow = "sun:10:30-sun:14:30",
            PubliclyAccessible = true,
            ReplicationInstanceClass = "dms.t2.micro",
            ReplicationInstanceId = "test-dms-replication-instance-tf",
            ReplicationSubnetGroupId = aws_dms_replication_subnet_group.Test_dms_replication_subnet_group_tf.Id,
            Tags = 
            {
                { "Name", "test" },
            },
            VpcSecurityGroupIds = 
            {
                "sg-12345678",
            },
        });
    }

}

