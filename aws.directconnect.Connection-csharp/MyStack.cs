using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var hoge = new Aws.DirectConnect.Connection("hoge", new Aws.DirectConnect.ConnectionArgs
        {
            Bandwidth = "1Gbps",
            Location = "EqDC2",
        });
    }

}

