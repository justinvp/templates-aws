using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ec2.TrafficMirrorFilter("foo", new Aws.Ec2.TrafficMirrorFilterArgs
        {
            Description = "traffic mirror filter - example",
            NetworkServices = 
            {
                "amazon-dns",
            },
        });
    }

}

