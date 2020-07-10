import pulumi
import pulumi_aws as aws

current = aws.elasticbeanstalk.get_hosted_zone()

