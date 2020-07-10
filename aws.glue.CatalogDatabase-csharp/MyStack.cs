using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var awsGlueCatalogDatabase = new Aws.Glue.CatalogDatabase("awsGlueCatalogDatabase", new Aws.Glue.CatalogDatabaseArgs
        {
            Name = "MyCatalogDatabase",
        });
    }

}

