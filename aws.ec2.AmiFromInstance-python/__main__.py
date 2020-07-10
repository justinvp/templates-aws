import pulumi
import pulumi_aws as aws

example = aws.ec2.AmiFromInstance("example", source_instance_id="i-xxxxxxxx")

