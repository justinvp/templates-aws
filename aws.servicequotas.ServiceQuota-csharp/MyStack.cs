using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ServiceQuotas.ServiceQuota("example", new Aws.ServiceQuotas.ServiceQuotaArgs
        {
            QuotaCode = "L-F678F1CE",
            ServiceCode = "vpc",
            Value = 75,
        });
    }

}

