using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleNetworkAcls = Output.Create(Aws.Ec2.GetNetworkAcls.InvokeAsync(new Aws.Ec2.GetNetworkAclsArgs
        {
            VpcId = @var.Vpc_id,
        }));
        this.Example = exampleNetworkAcls.Apply(exampleNetworkAcls => exampleNetworkAcls.Ids);
    }

    [Output("example")]
    public Output<string> Example { get; set; }
}

