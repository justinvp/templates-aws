using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
            EnableDnsHostnames = true,
            EnableDnsSupport = true,
        });
        var examplePrivateDnsNamespace = new Aws.ServiceDiscovery.PrivateDnsNamespace("examplePrivateDnsNamespace", new Aws.ServiceDiscovery.PrivateDnsNamespaceArgs
        {
            Description = "example",
            Vpc = exampleVpc.Id,
        });
        var exampleService = new Aws.ServiceDiscovery.Service("exampleService", new Aws.ServiceDiscovery.ServiceArgs
        {
            DnsConfig = new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigArgs
            {
                DnsRecords = 
                {
                    new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigDnsRecordArgs
                    {
                        Ttl = 10,
                        Type = "A",
                    },
                },
                NamespaceId = examplePrivateDnsNamespace.Id,
                RoutingPolicy = "MULTIVALUE",
            },
            HealthCheckCustomConfig = new Aws.ServiceDiscovery.Inputs.ServiceHealthCheckCustomConfigArgs
            {
                FailureThreshold = 1,
            },
        });
    }

}

