using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Organizations.GetOrganization.InvokeAsync());
        this.AccountIds = example.Apply(example => example.Accounts.Select(__item => __item.Id).ToList());
    }

    [Output("accountIds")]
    public Output<string> AccountIds { get; set; }
}

