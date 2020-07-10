using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var hoge = new Aws.DirectConnect.LinkAggregationGroup("hoge", new Aws.DirectConnect.LinkAggregationGroupArgs
        {
            ConnectionsBandwidth = "1Gbps",
            ForceDestroy = true,
            Location = "EqDC2",
        });
    }

}

