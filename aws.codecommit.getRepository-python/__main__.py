import pulumi
import pulumi_aws as aws

test = aws.codecommit.get_repository(repository_name="MyTestRepository")

