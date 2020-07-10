import pulumi
import pulumi_aws as aws

tftest = aws.elasticbeanstalk.Application("tftest", description="tf-test-desc")
tf_template = aws.elasticbeanstalk.ConfigurationTemplate("tfTemplate",
    application=tftest.name,
    solution_stack_name="64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4")

