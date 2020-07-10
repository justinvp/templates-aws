using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var subnetId = Output.Create(Aws.CloudFormation.GetExport.InvokeAsync(new Aws.CloudFormation.GetExportArgs
        {
            Name = "mySubnetIdExportName",
        }));
        var web = new Aws.Ec2.Instance("web", new Aws.Ec2.InstanceArgs
        {
            Ami = "ami-abb07bcb",
            InstanceType = "t1.micro",
            SubnetId = subnetId.Apply(subnetId => subnetId.Value),
        });
    }

}

