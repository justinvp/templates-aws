import pulumi
import pulumi_aws as aws

ecs_instance_role_role = aws.iam.Role("ecsInstanceRoleRole", assume_role_policy="""{
    "Version": "2012-10-17",
    "Statement": [
	{
	    "Action": "sts:AssumeRole",
	    "Effect": "Allow",
	    "Principal": {
		"Service": "ec2.amazonaws.com"
	    }
	}
    ]
}

""")
ecs_instance_role_role_policy_attachment = aws.iam.RolePolicyAttachment("ecsInstanceRoleRolePolicyAttachment",
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role",
    role=ecs_instance_role_role.name)
ecs_instance_role_instance_profile = aws.iam.InstanceProfile("ecsInstanceRoleInstanceProfile", role=ecs_instance_role_role.name)
aws_batch_service_role_role = aws.iam.Role("awsBatchServiceRoleRole", assume_role_policy="""{
    "Version": "2012-10-17",
    "Statement": [
	{
	    "Action": "sts:AssumeRole",
	    "Effect": "Allow",
	    "Principal": {
		"Service": "batch.amazonaws.com"
	    }
	}
    ]
}

""")
aws_batch_service_role_role_policy_attachment = aws.iam.RolePolicyAttachment("awsBatchServiceRoleRolePolicyAttachment",
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole",
    role=aws_batch_service_role_role.name)
sample_security_group = aws.ec2.SecurityGroup("sampleSecurityGroup", egress=[{
    "cidr_blocks": ["0.0.0.0/0"],
    "from_port": 0,
    "protocol": "-1",
    "to_port": 0,
}])
sample_vpc = aws.ec2.Vpc("sampleVpc", cidr_block="10.1.0.0/16")
sample_subnet = aws.ec2.Subnet("sampleSubnet",
    cidr_block="10.1.1.0/24",
    vpc_id=sample_vpc.id)
sample_compute_environment = aws.batch.ComputeEnvironment("sampleComputeEnvironment",
    compute_environment_name="sample",
    compute_resources={
        "instanceRole": ecs_instance_role_instance_profile.arn,
        "instance_type": ["c4.large"],
        "maxVcpus": 16,
        "minVcpus": 0,
        "security_group_ids": [sample_security_group.id],
        "subnets": [sample_subnet.id],
        "type": "EC2",
    },
    service_role=aws_batch_service_role_role.arn,
    type="MANAGED",
    opts=ResourceOptions(depends_on=["aws_iam_role_policy_attachment.aws_batch_service_role"]))

