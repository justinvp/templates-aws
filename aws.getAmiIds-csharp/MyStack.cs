using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ubuntu = Output.Create(Aws.GetAmiIds.InvokeAsync(new Aws.GetAmiIdsArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiIdsFilterArgs
                {
                    Name = "name",
                    Values = 
                    {
                        "ubuntu/images/ubuntu-*-*-amd64-server-*",
                    },
                },
            },
            Owners = 
            {
                "099720109477",
            },
        }));
    }

}

