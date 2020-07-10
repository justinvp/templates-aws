import pulumi
import pulumi_aws as aws

example = aws.elasticbeanstalk.get_application(name="example")
pulumi.export("arn", example.arn)
pulumi.export("description", example.description)

