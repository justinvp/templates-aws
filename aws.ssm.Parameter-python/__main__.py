import pulumi
import pulumi_aws as aws

foo = aws.ssm.Parameter("foo",
    type="String",
    value="bar")

