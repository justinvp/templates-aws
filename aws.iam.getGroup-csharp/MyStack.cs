using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Iam.GetGroup.InvokeAsync(new Aws.Iam.GetGroupArgs
        {
            GroupName = "an_example_group_name",
        }));
    }

}

