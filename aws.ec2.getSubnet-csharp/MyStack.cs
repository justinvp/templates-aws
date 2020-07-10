using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var subnetId = config.RequireObject<dynamic>("subnetId");
        var selected = Output.Create(Aws.Ec2.GetSubnet.InvokeAsync(new Aws.Ec2.GetSubnetArgs
        {
            Id = subnetId,
        }));
        var subnet = new Aws.Ec2.SecurityGroup("subnet", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    CidrBlocks = 
                    {
                        selected.Apply(selected => selected.CidrBlock),
                    },
                    FromPort = 80,
                    Protocol = "tcp",
                    ToPort = 80,
                },
            },
            VpcId = selected.Apply(selected => selected.VpcId),
        });
    }

}

