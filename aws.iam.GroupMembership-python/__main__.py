import pulumi
import pulumi_aws as aws

group = aws.iam.Group("group")
user_one = aws.iam.User("userOne")
user_two = aws.iam.User("userTwo")
team = aws.iam.GroupMembership("team",
    group=group.name,
    users=[
        user_one.name,
        user_two.name,
    ])

