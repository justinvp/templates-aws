using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ElasticSearch.Domain("example", new Aws.ElasticSearch.DomainArgs
        {
            ClusterConfig = new Aws.ElasticSearch.Inputs.DomainClusterConfigArgs
            {
                ClusterConfig = "r4.large.elasticsearch",
            },
            ElasticsearchVersion = "1.5",
            SnapshotOptions = new Aws.ElasticSearch.Inputs.DomainSnapshotOptionsArgs
            {
                SnapshotOptions = 23,
            },
            Tags = 
            {
                { "Domain", "TestDomain" },
            },
        });
    }

}

