using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var domainTest = new Aws.LightSail.Domain("domainTest", new Aws.LightSail.DomainArgs
        {
            DomainName = "mydomain.com",
        });
    }

}

