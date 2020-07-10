using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var securityGroupId = config.RequireObject<dynamic>("securityGroupId");
        var selected = Output.Create(Aws.Ec2.GetSecurityGroup.InvokeAsync(new Aws.Ec2.GetSecurityGroupArgs
        {
            Id = securityGroupId,
        }));
        var subnet = new Aws.Ec2.Subnet("subnet", new Aws.Ec2.SubnetArgs
        {
            CidrBlock = "10.0.1.0/24",
            VpcId = selected.Apply(selected => selected.VpcId),
        });
    }

}

