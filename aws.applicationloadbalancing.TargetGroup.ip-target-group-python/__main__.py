import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
ip_example = aws.lb.TargetGroup("ip-example",
    port=80,
    protocol="HTTP",
    target_type="ip",
    vpc_id=main.id)

