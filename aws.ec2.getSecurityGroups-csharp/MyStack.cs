using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = Output.Create(Aws.Ec2.GetSecurityGroups.InvokeAsync(new Aws.Ec2.GetSecurityGroupsArgs
        {
            Tags = 
            {
                { "Application", "k8s" },
                { "Environment", "dev" },
            },
        }));
    }

}

