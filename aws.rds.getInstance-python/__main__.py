import pulumi
import pulumi_aws as aws

database = aws.rds.get_instance(db_instance_identifier="my-test-database")

