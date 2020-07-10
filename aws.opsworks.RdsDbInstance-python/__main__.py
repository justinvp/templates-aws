import pulumi
import pulumi_aws as aws

my_instance = aws.opsworks.RdsDbInstance("myInstance",
    db_password="somePass",
    db_user="someUser",
    rds_db_instance_arn=aws_db_instance["my_instance"]["arn"],
    stack_id=aws_opsworks_stack["my_stack"]["id"])

