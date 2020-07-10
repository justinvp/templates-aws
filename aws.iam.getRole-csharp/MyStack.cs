using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Iam.GetRole.InvokeAsync(new Aws.Iam.GetRoleArgs
        {
            Name = "an_example_role_name",
        }));
    }

}

