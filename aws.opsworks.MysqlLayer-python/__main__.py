import pulumi
import pulumi_aws as aws

db = aws.opsworks.MysqlLayer("db", stack_id=aws_opsworks_stack["main"]["id"])

