using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Request a spot instance at $0.03
        var cheapWorker = new Aws.Ec2.SpotInstanceRequest("cheapWorker", new Aws.Ec2.SpotInstanceRequestArgs
        {
            Ami = "ami-1234",
            InstanceType = "c4.xlarge",
            SpotPrice = "0.03",
            Tags = 
            {
                { "Name", "CheapWorker" },
            },
        });
    }

}

