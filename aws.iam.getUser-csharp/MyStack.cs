using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Iam.GetUser.InvokeAsync(new Aws.Iam.GetUserArgs
        {
            UserName = "an_example_user_name",
        }));
    }

}

