import pulumi
import pulumi_aws as aws

example_network_acls = aws.ec2.get_network_acls(vpc_id=var["vpc_id"])
pulumi.export("example", example_network_acls.ids)

