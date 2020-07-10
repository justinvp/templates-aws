using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var awsGlueCatalogTable = new Aws.Glue.CatalogTable("awsGlueCatalogTable", new Aws.Glue.CatalogTableArgs
        {
            DatabaseName = "MyCatalogDatabase",
            Name = "MyCatalogTable",
        });
    }

}

