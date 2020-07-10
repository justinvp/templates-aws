import pulumi
import pulumi_aws as aws

users = aws.simpledb.Domain("users")

