import pulumi
import pulumi_aws as aws

foo_vpc = aws.ec2.Vpc("fooVpc",
    cidr_block="10.0.0.0/16",
    tags={
        "Name": "tf-test",
    })
foo_subnet = aws.ec2.Subnet("fooSubnet",
    availability_zone="us-west-2a",
    cidr_block="10.0.0.0/24",
    tags={
        "Name": "tf-test",
    },
    vpc_id=foo_vpc.id)
bar = aws.elasticache.SubnetGroup("bar", subnet_ids=[foo_subnet.id])

