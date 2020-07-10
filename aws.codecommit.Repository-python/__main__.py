import pulumi
import pulumi_aws as aws

test = aws.codecommit.Repository("test",
    description="This is the Sample App Repository",
    repository_name="MyTestRepository")

