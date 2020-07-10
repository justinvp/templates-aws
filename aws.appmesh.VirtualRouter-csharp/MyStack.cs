using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var serviceb = new Aws.AppMesh.VirtualRouter("serviceb", new Aws.AppMesh.VirtualRouterArgs
        {
            MeshName = aws_appmesh_mesh.Simple.Id,
            Spec = new Aws.AppMesh.Inputs.VirtualRouterSpecArgs
            {
                Listener = new Aws.AppMesh.Inputs.VirtualRouterSpecListenerArgs
                {
                    PortMapping = new Aws.AppMesh.Inputs.VirtualRouterSpecListenerPortMappingArgs
                    {
                        Port = 8080,
                        Protocol = "http",
                    },
                },
            },
        });
    }

}

