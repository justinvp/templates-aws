import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
foo = aws.ec2.Subnet("foo",
    availability_zone="us-west-2a",
    cidr_block="10.0.1.0/24",
    vpc_id=main.id)
bar = aws.ec2.Subnet("bar",
    availability_zone="us-west-2b",
    cidr_block="10.0.2.0/24",
    vpc_id=main.id)
connector = aws.directoryservice.Directory("connector",
    connect_settings={
        "customerDnsIps": ["A.B.C.D"],
        "customerUsername": "Admin",
        "subnet_ids": [
            foo.id,
            bar.id,
        ],
        "vpc_id": main.id,
    },
    password="SuperSecretPassw0rd",
    size="Small",
    type="ADConnector")

