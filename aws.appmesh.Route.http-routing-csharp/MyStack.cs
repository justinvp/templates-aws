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
                HttpRoute = new Aws.AppMesh.Inputs.RouteSpecHttpRouteArgs
                {
                    Action = new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionArgs
                    {
                        WeightedTarget = 
                        {
                            
                            {
                                { "virtualNode", aws_appmesh_virtual_node.Serviceb1.Name },
                                { "weight", 90 },
                            },
                            
                            {
                                { "virtualNode", aws_appmesh_virtual_node.Serviceb2.Name },
                                { "weight", 10 },
                            },
                        },
                    },
                    Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchArgs
                    {
                        Prefix = "/",
                    },
                },
            },
            VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
        });
    }

}

