using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Iam.GetPolicy.InvokeAsync(new Aws.Iam.GetPolicyArgs
        {
            Arn = "arn:aws:iam::123456789012:policy/UsersManageOwnCredentials",
        }));
    }

}

