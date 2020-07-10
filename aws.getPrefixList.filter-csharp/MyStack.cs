using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = Output.Create(Aws.GetPrefixList.InvokeAsync(new Aws.GetPrefixListArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetPrefixListFilterArgs
                {
                    Name = "prefix-list-id",
                    Values = 
                    {
                        "pl-68a54001",
                    },
                },
            },
        }));
    }

}

