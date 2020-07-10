using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var endpoint = new Aws.Sagemaker.Endpoint("endpoint", new Aws.Sagemaker.EndpointArgs
        {
            EndpointConfigName = aws_sagemaker_endpoint_configuration.Ec.Name,
            Tags = 
            {
                { "Name", "foo" },
            },
        });
    }

}

