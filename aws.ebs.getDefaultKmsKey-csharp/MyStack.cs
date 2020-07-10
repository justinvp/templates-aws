using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var current = Output.Create(Aws.Ebs.GetDefaultKmsKey.InvokeAsync());
        var example = new Aws.Ebs.Volume("example", new Aws.Ebs.VolumeArgs
        {
            AvailabilityZone = "us-west-2a",
            Encrypted = true,
            KmsKeyId = current.Apply(current => current.KeyArn),
        });
    }

}

