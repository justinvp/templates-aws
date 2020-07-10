import pulumi
import pulumi_aws as aws

yada = aws.cloudwatch.LogGroup("yada")
foo = aws.cloudwatch.LogStream("foo", log_group_name=yada.name)

