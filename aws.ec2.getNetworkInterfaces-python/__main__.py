import pulumi
import pulumi_aws as aws

example_network_interfaces = aws.ec2.get_network_interfaces()
pulumi.export("example", example_network_interfaces.ids)

