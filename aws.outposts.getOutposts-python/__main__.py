import pulumi
import pulumi_aws as aws

example = aws.outposts.get_outposts(site_id=data["aws_outposts_site"]["id"])

