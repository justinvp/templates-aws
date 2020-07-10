import pulumi
import pulumi_aws as aws

yada = aws.cloudwatch.LogGroup("yada", tags={
    "Application": "serviceA",
    "Environment": "production",
})

