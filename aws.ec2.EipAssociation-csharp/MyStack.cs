using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var web = new Aws.Ec2.Instance("web", new Aws.Ec2.InstanceArgs
        {
            Ami = "ami-21f78e11",
            AvailabilityZone = "us-west-2a",
            InstanceType = "t1.micro",
            Tags = 
            {
                { "Name", "HelloWorld" },
            },
        });
        var example = new Aws.Ec2.Eip("example", new Aws.Ec2.EipArgs
        {
            Vpc = true,
        });
        var eipAssoc = new Aws.Ec2.EipAssociation("eipAssoc", new Aws.Ec2.EipAssociationArgs
        {
            AllocationId = example.Id,
            InstanceId = web.Id,
        });
    }

}

