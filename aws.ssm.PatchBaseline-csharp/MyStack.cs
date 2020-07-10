using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var production = new Aws.Ssm.PatchBaseline("production", new Aws.Ssm.PatchBaselineArgs
        {
            ApprovedPatches = 
            {
                "KB123456",
            },
        });
    }

}

