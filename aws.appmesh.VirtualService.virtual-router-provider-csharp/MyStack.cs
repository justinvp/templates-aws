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
                    VirtualRouter = new Aws.AppMesh.Inputs.VirtualServiceSpecProviderVirtualRouterArgs
                    {
                        VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
                    },
                },
            },
        });
    }

}

