using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = Output.Create(Aws.AutoScaling.GetGroup.InvokeAsync(new Aws.AutoScaling.GetGroupArgs
        {
            Name = "foo",
        }));
    }

}

