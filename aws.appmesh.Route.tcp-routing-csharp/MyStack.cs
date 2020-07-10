using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var serviceb = new Aws.AppMesh.Route("serviceb", new Aws.AppMesh.RouteArgs
        {
            MeshName = aws_appmesh_mesh.Simple.Id,
            Spec = new Aws.AppMesh.Inputs.RouteSpecArgs
            {
                TcpRoute = new Aws.AppMesh.Inputs.RouteSpecTcpRouteArgs
                {
                    Action = new Aws.AppMesh.Inputs.RouteSpecTcpRouteActionArgs
                    {
                        WeightedTarget = 
                        {
                            
                            {
                                { "virtualNode", aws_appmesh_virtual_node.Serviceb1.Name },
                                { "weight", 100 },
                            },
                        },
                    },
                },
            },
            VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
        });
    }

}

