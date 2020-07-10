using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ecsInstanceRoleRole = new Aws.Iam.Role("ecsInstanceRoleRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
    ""Version"": ""2012-10-17"",
    ""Statement"": [
	{
	    ""Action"": ""sts:AssumeRole"",
	    ""Effect"": ""Allow"",
	    ""Principal"": {
		""Service"": ""ec2.amazonaws.com""
	    }
	}
    ]
}

",
        });
        var ecsInstanceRoleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("ecsInstanceRoleRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role",
            Role = ecsInstanceRoleRole.Name,
        });
        var ecsInstanceRoleInstanceProfile = new Aws.Iam.InstanceProfile("ecsInstanceRoleInstanceProfile", new Aws.Iam.InstanceProfileArgs
        {
            Role = ecsInstanceRoleRole.Name,
        });
        var awsBatchServiceRoleRole = new Aws.Iam.Role("awsBatchServiceRoleRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
    ""Version"": ""2012-10-17"",
    ""Statement"": [
	{
	    ""Action"": ""sts:AssumeRole"",
	    ""Effect"": ""Allow"",
	    ""Principal"": {
		""Service"": ""batch.amazonaws.com""
	    }
	}
    ]
}

",
        });
        var awsBatchServiceRoleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("awsBatchServiceRoleRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole",
            Role = awsBatchServiceRoleRole.Name,
        });
        var sampleSecurityGroup = new Aws.Ec2.SecurityGroup("sampleSecurityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Egress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupEgressArgs
                {
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                    FromPort = 0,
                    Protocol = "-1",
                    ToPort = 0,
                },
            },
        });
        var sampleVpc = new Aws.Ec2.Vpc("sampleVpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.1.0.0/16",
        });
        var sampleSubnet = new Aws.Ec2.Subnet("sampleSubnet", new Aws.Ec2.SubnetArgs
        {
            CidrBlock = "10.1.1.0/24",
            VpcId = sampleVpc.Id,
        });
        var sampleComputeEnvironment = new Aws.Batch.ComputeEnvironment("sampleComputeEnvironment", new Aws.Batch.ComputeEnvironmentArgs
        {
            ComputeEnvironmentName = "sample",
            ComputeResources = new Aws.Batch.Inputs.ComputeEnvironmentComputeResourcesArgs
            {
                InstanceRole = ecsInstanceRoleInstanceProfile.Arn,
                InstanceType = 
                {
                    "c4.large",
                },
                MaxVcpus = 16,
                MinVcpus = 0,
                SecurityGroupIds = 
                {
                    sampleSecurityGroup.Id,
                },
                Subnets = 
                {
                    sampleSubnet.Id,
                },
                Type = "EC2",
            },
            ServiceRole = awsBatchServiceRoleRole.Arn,
            Type = "MANAGED",
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_iam_role_policy_attachment.aws_batch_service_role",
            },
        });
    }

}

