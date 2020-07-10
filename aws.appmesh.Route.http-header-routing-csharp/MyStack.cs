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
                                { "virtualNode", aws_appmesh_virtual_node.Serviceb.Name },
                                { "weight", 100 },
                            },
                        },
                    },
                    Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchArgs
                    {
                        Header = 
                        {
                            
                            {
                                { "match", 
                                {
                                    { "prefix", "123" },
                                } },
                                { "name", "clientRequestId" },
                            },
                        },
                        Method = "POST",
                        Prefix = "/",
                        Scheme = "https",
                    },
                },
            },
            VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
        });
    }

}

