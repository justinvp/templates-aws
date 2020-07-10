import pulumi
import pulumi_aws as aws

example = aws.outposts.get_outpost(name="example")

