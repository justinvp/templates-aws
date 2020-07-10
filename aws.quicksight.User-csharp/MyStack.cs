using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Quicksight.User("example", new Aws.Quicksight.UserArgs
        {
            Email = "author@example.com",
            IdentityType = "IAM",
            UserName = "an-author",
            UserRole = "AUTHOR",
        });
    }

}

