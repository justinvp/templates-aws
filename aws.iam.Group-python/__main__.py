import pulumi
import pulumi_aws as aws

developers = aws.iam.Group("developers", path="/users/")

