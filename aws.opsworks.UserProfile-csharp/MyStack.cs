using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myProfile = new Aws.OpsWorks.UserProfile("myProfile", new Aws.OpsWorks.UserProfileArgs
        {
            SshUsername = "my_user",
            UserArn = aws_iam_user.User.Arn,
        });
    }

}

