import pulumi
import pulumi_aws as aws

user_updates = aws.sns.Topic("userUpdates")

