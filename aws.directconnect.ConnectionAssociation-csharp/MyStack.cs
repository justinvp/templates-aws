using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleConnection = new Aws.DirectConnect.Connection("exampleConnection", new Aws.DirectConnect.ConnectionArgs
        {
            Bandwidth = "1Gbps",
            Location = "EqSe2",
        });
        var exampleLinkAggregationGroup = new Aws.DirectConnect.LinkAggregationGroup("exampleLinkAggregationGroup", new Aws.DirectConnect.LinkAggregationGroupArgs
        {
            ConnectionsBandwidth = "1Gbps",
            Location = "EqSe2",
        });
        var exampleConnectionAssociation = new Aws.DirectConnect.ConnectionAssociation("exampleConnectionAssociation", new Aws.DirectConnect.ConnectionAssociationArgs
        {
            ConnectionId = exampleConnection.Id,
            LagId = exampleLinkAggregationGroup.Id,
        });
    }

}

