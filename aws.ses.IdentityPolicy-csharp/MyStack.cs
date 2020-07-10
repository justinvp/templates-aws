using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleDomainIdentity = new Aws.Ses.DomainIdentity("exampleDomainIdentity", new Aws.Ses.DomainIdentityArgs
        {
            Domain = "example.com",
        });
        var examplePolicyDocument = exampleDomainIdentity.Arn.Apply(arn => Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "SES:SendEmail",
                        "SES:SendRawEmail",
                    },
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "*",
                            },
                            Type = "AWS",
                        },
                    },
                    Resources = 
                    {
                        arn,
                    },
                },
            },
        }));
        var exampleIdentityPolicy = new Aws.Ses.IdentityPolicy("exampleIdentityPolicy", new Aws.Ses.IdentityPolicyArgs
        {
            Identity = exampleDomainIdentity.Arn,
            Policy = examplePolicyDocument.Apply(examplePolicyDocument => examplePolicyDocument.Json),
        });
    }

}

