import pulumi
import pulumi_aws as aws

foo = aws.ec2.TrafficMirrorFilter("foo",
    description="traffic mirror filter - example",
    network_services=["amazon-dns"])

