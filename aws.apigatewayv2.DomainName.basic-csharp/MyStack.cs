using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.DomainName("example", new Aws.ApiGatewayV2.DomainNameArgs
        {
            DomainName = "ws-api.example.com",
            DomainNameConfiguration = new Aws.ApiGatewayV2.Inputs.DomainNameDomainNameConfigurationArgs
            {
                CertificateArn = aws_acm_certificate.Example.Arn,
                EndpointType = "REGIONAL",
                SecurityPolicy = "TLS_1_2",
            },
        });
    }

}

