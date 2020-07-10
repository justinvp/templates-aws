using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Pricing.GetProduct.InvokeAsync(new Aws.Pricing.GetProductArgs
        {
            Filters = 
            {
                new Aws.Pricing.Inputs.GetProductFilterArgs
                {
                    Field = "instanceType",
                    Value = "c5.xlarge",
                },
                new Aws.Pricing.Inputs.GetProductFilterArgs
                {
                    Field = "operatingSystem",
                    Value = "Linux",
                },
                new Aws.Pricing.Inputs.GetProductFilterArgs
                {
                    Field = "location",
                    Value = "US East (N. Virginia)",
                },
                new Aws.Pricing.Inputs.GetProductFilterArgs
                {
                    Field = "preInstalledSw",
                    Value = "NA",
                },
                new Aws.Pricing.Inputs.GetProductFilterArgs
                {
                    Field = "licenseModel",
                    Value = "No License required",
                },
                new Aws.Pricing.Inputs.GetProductFilterArgs
                {
                    Field = "tenancy",
                    Value = "Shared",
                },
                new Aws.Pricing.Inputs.GetProductFilterArgs
                {
                    Field = "capacitystatus",
                    Value = "Used",
                },
            },
            ServiceCode = "AmazonEC2",
        }));
    }

}

