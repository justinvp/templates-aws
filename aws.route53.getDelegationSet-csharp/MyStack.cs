using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dset = Output.Create(Aws.Route53.GetDelegationSet.InvokeAsync(new Aws.Route53.GetDelegationSetArgs
        {
            Id = "MQWGHCBFAKEID",
        }));
    }

}

