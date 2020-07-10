using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var network = Output.Create(Aws.CloudFormation.GetStack.InvokeAsync(new Aws.CloudFormation.GetStackArgs
        {
            Name = "my-network-stack",
        }));
        var web = new Aws.Ec2.Instance("web", new Aws.Ec2.InstanceArgs
        {
            Ami = "ami-abb07bcb",
            InstanceType = "t1.micro",
            SubnetId = network.Apply(network => network.Outputs.SubnetId),
            Tags = 
            {
                { "Name", "HelloWorld" },
            },
        });
    }

}

