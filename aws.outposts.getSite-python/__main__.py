import pulumi
import pulumi_aws as aws

example = aws.outposts.get_site(name="example")

