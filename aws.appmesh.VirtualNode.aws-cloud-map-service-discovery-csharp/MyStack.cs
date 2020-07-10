using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ServiceDiscovery.HttpNamespace("example", new Aws.ServiceDiscovery.HttpNamespaceArgs
        {
        });
        var serviceb1 = new Aws.AppMesh.VirtualNode("serviceb1", new Aws.AppMesh.VirtualNodeArgs
        {
            MeshName = aws_appmesh_mesh.Simple.Id,
            Spec = new Aws.AppMesh.Inputs.VirtualNodeSpecArgs
            {
                Backend = 
                {
                    
                    {
                        { "virtualService", 
                        {
                            { "virtualServiceName", "servicea.simpleapp.local" },
                        } },
                    },
                },
                Listener = new Aws.AppMesh.Inputs.VirtualNodeSpecListenerArgs
                {
                    PortMapping = new Aws.AppMesh.Inputs.VirtualNodeSpecListenerPortMappingArgs
                    {
                        Port = 8080,
                        Protocol = "http",
                    },
                },
                ServiceDiscovery = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryArgs
                {
                    AwsCloudMap = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryAwsCloudMapArgs
                    {
                        Attributes = 
                        {
                            { "stack", "blue" },
                        },
                        NamespaceName = example.Name,
                        ServiceName = "serviceb1",
                    },
                },
            },
        });
    }

}

