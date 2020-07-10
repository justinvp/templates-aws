import pulumi
import pulumi_aws as aws

user1 = aws.iam.User("user1")
group1 = aws.iam.Group("group1")
group2 = aws.iam.Group("group2")
example1 = aws.iam.UserGroupMembership("example1",
    groups=[
        group1.name,
        group2.name,
    ],
    user=user1.name)
group3 = aws.iam.Group("group3")
example2 = aws.iam.UserGroupMembership("example2",
    groups=[group3.name],
    user=user1.name)

