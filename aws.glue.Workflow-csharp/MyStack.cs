using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Workflow("example", new Aws.Glue.WorkflowArgs
        {
        });
        var example_start = new Aws.Glue.Trigger("example-start", new Aws.Glue.TriggerArgs
        {
            Actions = 
            {
                new Aws.Glue.Inputs.TriggerActionArgs
                {
                    JobName = "example-job",
                },
            },
            Type = "ON_DEMAND",
            WorkflowName = example.Name,
        });
        var example_inner = new Aws.Glue.Trigger("example-inner", new Aws.Glue.TriggerArgs
        {
            Actions = 
            {
                new Aws.Glue.Inputs.TriggerActionArgs
                {
                    JobName = "another-example-job",
                },
            },
            Predicate = new Aws.Glue.Inputs.TriggerPredicateArgs
            {
                Conditions = 
                {
                    new Aws.Glue.Inputs.TriggerPredicateConditionArgs
                    {
                        JobName = "example-job",
                        State = "SUCCEEDED",
                    },
                },
            },
            Type = "CONDITIONAL",
            WorkflowName = example.Name,
        });
    }

}

