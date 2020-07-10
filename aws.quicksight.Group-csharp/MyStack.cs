using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Quicksight.Group("example", new Aws.Quicksight.GroupArgs
        {
            GroupName = "tf-example",
        });
    }

}

