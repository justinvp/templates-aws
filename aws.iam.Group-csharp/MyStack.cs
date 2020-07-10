using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var developers = new Aws.Iam.Group("developers", new Aws.Iam.GroupArgs
        {
            Path = "/users/",
        });
    }

}

