using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var web = new Aws.Ec2.PlacementGroup("web", new Aws.Ec2.PlacementGroupArgs
        {
            Strategy = "cluster",
        });
    }

}

