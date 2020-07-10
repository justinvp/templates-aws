using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.AmiFromInstance("example", new Aws.Ec2.AmiFromInstanceArgs
        {
            SourceInstanceId = "i-xxxxxxxx",
        });
    }

}

