using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Sfn.GetStateMachine.InvokeAsync(new Aws.Sfn.GetStateMachineArgs
        {
            Name = "an_example_sfn_name",
        }));
    }

}

