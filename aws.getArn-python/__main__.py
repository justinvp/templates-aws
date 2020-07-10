import pulumi
import pulumi_aws as aws

db_instance = aws.get_arn(arn="arn:aws:rds:eu-west-1:123456789012:db:mysql-db")

