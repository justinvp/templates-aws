import pulumi
import pulumi_aws as aws

lambda_example = aws.lb.TargetGroup("lambda-example", target_type="lambda")

