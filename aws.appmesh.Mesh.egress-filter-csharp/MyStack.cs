using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var simple = new Aws.AppMesh.Mesh("simple", new Aws.AppMesh.MeshArgs
        {
            Spec = new Aws.AppMesh.Inputs.MeshSpecArgs
            {
                EgressFilter = new Aws.AppMesh.Inputs.MeshSpecEgressFilterArgs
                {
                    Type = "ALLOW_ALL",
                },
            },
        });
    }

}

