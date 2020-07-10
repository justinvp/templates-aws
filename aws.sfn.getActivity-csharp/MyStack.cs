using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sfnActivity = Output.Create(Aws.Sfn.GetActivity.InvokeAsync(new Aws.Sfn.GetActivityArgs
        {
            Name = "my-activity",
        }));
    }

}

