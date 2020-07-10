import pulumi
import pulumi_aws as aws

foo = aws.ec2.Vpc("foo", cidr_block="10.0.0.0/16")
alpha_subnet = aws.ec2.Subnet("alphaSubnet",
    availability_zone="us-west-2a",
    cidr_block="10.0.1.0/24",
    vpc_id=foo.id)
alpha_mount_target = aws.efs.MountTarget("alphaMountTarget",
    file_system_id=aws_efs_file_system["foo"]["id"],
    subnet_id=alpha_subnet.id)

