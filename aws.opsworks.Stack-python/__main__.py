import pulumi
import pulumi_aws as aws

main = aws.opsworks.Stack("main",
    custom_json="""{
 "foobar": {
    "version": "1.0.0"
  }
}

""",
    default_instance_profile_arn=aws_iam_instance_profile["opsworks"]["arn"],
    region="us-west-1",
    service_role_arn=aws_iam_role["opsworks"]["arn"],
    tags={
        "Name": "foobar-stack",
    })

