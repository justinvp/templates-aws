using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2ClientVpn.Endpoint("example", new Aws.Ec2ClientVpn.EndpointArgs
        {
            AuthenticationOptions = 
            {
                new Aws.Ec2ClientVpn.Inputs.EndpointAuthenticationOptionArgs
                {
                    RootCertificateChainArn = aws_acm_certificate.Root_cert.Arn,
                    Type = "certificate-authentication",
                },
            },
            ClientCidrBlock = "10.0.0.0/16",
            ConnectionLogOptions = new Aws.Ec2ClientVpn.Inputs.EndpointConnectionLogOptionsArgs
            {
                CloudwatchLogGroup = aws_cloudwatch_log_group.Lg.Name,
                CloudwatchLogStream = aws_cloudwatch_log_stream.Ls.Name,
                Enabled = true,
            },
            Description = "clientvpn-example",
            ServerCertificateArn = aws_acm_certificate.Cert.Arn,
        });
    }

}

