import pulumi
import pulumi_aws as aws

example = aws.outposts.get_outpost_instance_types(arn=data["aws_outposts_outpost"]["example"]["arn"])

