import pulumi
import pulumi_aws as aws

foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.1.0.0/16")
foo_subnet = aws.ec2.Subnet("fooSubnet",
    availability_zone="us-west-2a",
    cidr_block="10.1.1.0/24",
    tags={
        "Name": "tf-dbsubnet-test-1",
    },
    vpc_id=foo_vpc.id)
bar = aws.ec2.Subnet("bar",
    availability_zone="us-west-2b",
    cidr_block="10.1.2.0/24",
    tags={
        "Name": "tf-dbsubnet-test-2",
    },
    vpc_id=foo_vpc.id)
foo_subnet_group = aws.redshift.SubnetGroup("fooSubnetGroup",
    subnet_ids=[
        foo_subnet.id,
        bar.id,
    ],
    tags={
        "environment": "Production",
    })

