using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var selected = Output.Create(Aws.Ec2.GetVpnGateway.InvokeAsync(new Aws.Ec2.GetVpnGatewayArgs
        {
            Filters = 
            {
                new Aws.Ec2.Inputs.GetVpnGatewayFilterArgs
                {
                    Name = "tag:Name",
                    Values = 
                    {
                        "vpn-gw",
                    },
                },
            },
        }));
        this.VpnGatewayId = selected.Apply(selected => selected.Id);
    }

    [Output("vpnGatewayId")]
    public Output<string> VpnGatewayId { get; set; }
}

