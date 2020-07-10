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
        var patchgroup = new Aws.Ssm.PatchGroup("patchgroup", new Aws.Ssm.PatchGroupArgs
        {
            BaselineId = production.Id,
            PatchGroup = "patch-group-name",
        });
    }

}

