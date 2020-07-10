import pulumi
import pulumi_aws as aws

tftest = aws.elasticbeanstalk.Application("tftest", description="tf-test-desc")
tfenvtest = aws.elasticbeanstalk.Environment("tfenvtest",
    application=tftest.name,
    solution_stack_name="64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4")

