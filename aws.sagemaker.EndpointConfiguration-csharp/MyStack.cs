using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ec = new Aws.Sagemaker.EndpointConfiguration("ec", new Aws.Sagemaker.EndpointConfigurationArgs
        {
            ProductionVariants = 
            {
                new Aws.Sagemaker.Inputs.EndpointConfigurationProductionVariantArgs
                {
                    InitialInstanceCount = 1,
                    InstanceType = "ml.t2.medium",
                    ModelName = aws_sagemaker_model.M.Name,
                    VariantName = "variant-1",
                },
            },
            Tags = 
            {
                { "Name", "foo" },
            },
        });
    }

}

