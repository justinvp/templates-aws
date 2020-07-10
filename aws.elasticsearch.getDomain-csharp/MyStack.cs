using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myDomain = Output.Create(Aws.ElasticSearch.GetDomain.InvokeAsync(new Aws.ElasticSearch.GetDomainArgs
        {
            DomainName = "my-domain-name",
        }));
    }

}

