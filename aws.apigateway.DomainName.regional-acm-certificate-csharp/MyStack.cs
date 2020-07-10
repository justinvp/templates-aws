using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleDomainName = new Aws.ApiGateway.DomainName("exampleDomainName", new Aws.ApiGateway.DomainNameArgs
        {
            DomainName = "api.example.com",
            EndpointConfiguration = new Aws.ApiGateway.Inputs.DomainNameEndpointConfigurationArgs
            {
                Types = "REGIONAL",
            },
            RegionalCertificateArn = aws_acm_certificate_validation.Example.Certificate_arn,
        });
        // Example DNS record using Route53.
        // Route53 is not specifically required; any DNS host can be used.
        var exampleRecord = new Aws.Route53.Record("exampleRecord", new Aws.Route53.RecordArgs
        {
            Aliases = 
            {
                new Aws.Route53.Inputs.RecordAliasArgs
                {
                    EvaluateTargetHealth = true,
                    Name = exampleDomainName.RegionalDomainName,
                    ZoneId = exampleDomainName.RegionalZoneId,
                },
            },
            Name = exampleDomainName.Domain,
            Type = "A",
            ZoneId = aws_route53_zone.Example.Id,
        });
    }

}

