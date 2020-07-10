import pulumi
import pulumi_aws as aws

table_name = aws.dynamodb.get_table(name="tableName")

