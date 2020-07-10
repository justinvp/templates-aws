import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
test = aws.lb.TargetGroup("test",
    port=80,
    protocol="HTTP",
    vpc_id=main.id)

