using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ptfeServiceVpcEndpoint = new Aws.Ec2.VpcEndpoint("ptfeServiceVpcEndpoint", new Aws.Ec2.VpcEndpointArgs
        {
            PrivateDnsEnabled = false,
            SecurityGroupIds = 
            {
                aws_security_group.Ptfe_service.Id,
            },
            ServiceName = @var.Ptfe_service,
            SubnetIds = 
            {
                local.Subnet_ids,
            },
            VpcEndpointType = "Interface",
            VpcId = @var.Vpc_id,
        });
        var @internal = Output.Create(Aws.Route53.GetZone.InvokeAsync(new Aws.Route53.GetZoneArgs
        {
            Name = "vpc.internal.",
            PrivateZone = true,
            VpcId = @var.Vpc_id,
        }));
        var ptfeServiceRecord = new Aws.Route53.Record("ptfeServiceRecord", new Aws.Route53.RecordArgs
        {
            Name = @internal.Apply(@internal => $"ptfe.{@internal.Name}"),
            Records = 
            {
                ptfeServiceVpcEndpoint.DnsEntries.Apply(dnsEntries => dnsEntries[0])["dns_name"],
            },
            Ttl = 300,
            Type = "CNAME",
            ZoneId = @internal.Apply(@internal => @internal.ZoneId),
        });
    }

}

