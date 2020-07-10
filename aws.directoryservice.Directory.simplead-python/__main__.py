import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
foo = aws.ec2.Subnet("foo",
    availability_zone="us-west-2a",
    cidr_block="10.0.1.0/24",
    vpc_id=main.id)
bar_subnet = aws.ec2.Subnet("barSubnet",
    availability_zone="us-west-2b",
    cidr_block="10.0.2.0/24",
    vpc_id=main.id)
bar_directory = aws.directoryservice.Directory("barDirectory",
    password="SuperSecretPassw0rd",
    size="Small",
    tags={
        "Project": "foo",
    },
    vpc_settings={
        "subnet_ids": [
            foo.id,
            bar_subnet.id,
        ],
        "vpc_id": main.id,
    })

