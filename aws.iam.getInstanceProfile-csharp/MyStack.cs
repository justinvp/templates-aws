using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Iam.GetInstanceProfile.InvokeAsync(new Aws.Iam.GetInstanceProfileArgs
        {
            Name = "an_example_instance_profile_name",
        }));
    }

}

