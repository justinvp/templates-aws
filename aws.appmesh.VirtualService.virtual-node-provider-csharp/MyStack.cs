using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var servicea = new Aws.AppMesh.VirtualService("servicea", new Aws.AppMesh.VirtualServiceArgs
        {
            MeshName = aws_appmesh_mesh.Simple.Id,
            Spec = new Aws.AppMesh.Inputs.VirtualServiceSpecArgs
            {
                Provider = new Aws.AppMesh.Inputs.VirtualServiceSpecProviderArgs
                {
                    VirtualNode = new Aws.AppMesh.Inputs.VirtualServiceSpecProviderVirtualNodeArgs
                    {
                        VirtualNodeName = aws_appmesh_virtual_node.Serviceb1.Name,
                    },
                },
            },
        });
    }

}

