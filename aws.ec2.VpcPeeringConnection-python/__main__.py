import pulumi
import pulumi_aws as aws

foo = aws.ec2.VpcPeeringConnection("foo",
    peer_owner_id=var["peer_owner_id"],
    peer_vpc_id=aws_vpc["bar"]["id"],
    vpc_id=aws_vpc["foo"]["id"])

