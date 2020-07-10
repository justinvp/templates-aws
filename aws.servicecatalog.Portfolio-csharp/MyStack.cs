using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var portfolio = new Aws.ServiceCatalog.Portfolio("portfolio", new Aws.ServiceCatalog.PortfolioArgs
        {
            Description = "List of my organizations apps",
            ProviderName = "Brett",
        });
    }

}

