import pulumi
import pulumi_aws as aws

ni = aws.sagemaker.NotebookInstance("ni",
    instance_type="ml.t2.medium",
    role_arn=aws_iam_role["role"]["arn"],
    tags={
        "Name": "foo",
    })

