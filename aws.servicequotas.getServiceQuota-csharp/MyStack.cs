using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var byQuotaCode = Output.Create(Aws.ServiceQuotas.GetServiceQuota.InvokeAsync(new Aws.ServiceQuotas.GetServiceQuotaArgs
        {
            QuotaCode = "L-F678F1CE",
            ServiceCode = "vpc",
        }));
        var byQuotaName = Output.Create(Aws.ServiceQuotas.GetServiceQuota.InvokeAsync(new Aws.ServiceQuotas.GetServiceQuotaArgs
        {
            QuotaName = "VPCs per Region",
            ServiceCode = "vpc",
        }));
    }

}

