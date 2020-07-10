using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Glue.GetScript.InvokeAsync(new Aws.Glue.GetScriptArgs
        {
            DagEdges = 
            {
                new Aws.Glue.Inputs.GetScriptDagEdgeArgs
                {
                    Source = "datasource0",
                    Target = "applymapping1",
                },
                new Aws.Glue.Inputs.GetScriptDagEdgeArgs
                {
                    Source = "applymapping1",
                    Target = "selectfields2",
                },
                new Aws.Glue.Inputs.GetScriptDagEdgeArgs
                {
                    Source = "selectfields2",
                    Target = "resolvechoice3",
                },
                new Aws.Glue.Inputs.GetScriptDagEdgeArgs
                {
                    Source = "resolvechoice3",
                    Target = "datasink4",
                },
            },
            DagNodes = 
            {
                new Aws.Glue.Inputs.GetScriptDagNodeArgs
                {
                    Args = 
                    {
                        new Aws.Glue.Inputs.GetScriptDagNodeArgArgs
                        {
                            Name = "database",
                            Value = $"\"{aws_glue_catalog_database.Source.Name}\"",
                        },
                        new Aws.Glue.Inputs.GetScriptDagNodeArgArgs
                        {
                            Name = "table_name",
                            Value = $"\"{aws_glue_catalog_table.Source.Name}\"",
                        },
                    },
                    Id = "datasource0",
                    NodeType = "DataSource",
                },
                new Aws.Glue.Inputs.GetScriptDagNodeArgs
                {
                    Args = 
                    {
                        new Aws.Glue.Inputs.GetScriptDagNodeArgArgs
                        {
                            Name = "mapping",
                            Value = "[(\"column1\", \"string\", \"column1\", \"string\")]",
                        },
                    },
                    Id = "applymapping1",
                    NodeType = "ApplyMapping",
                },
                new Aws.Glue.Inputs.GetScriptDagNodeArgs
                {
                    Args = 
                    {
                        new Aws.Glue.Inputs.GetScriptDagNodeArgArgs
                        {
                            Name = "paths",
                            Value = "[\"column1\"]",
                        },
                    },
                    Id = "selectfields2",
                    NodeType = "SelectFields",
                },
                new Aws.Glue.Inputs.GetScriptDagNodeArgs
                {
                    Args = 
                    {
                        new Aws.Glue.Inputs.GetScriptDagNodeArgArgs
                        {
                            Name = "choice",
                            Value = "\"MATCH_CATALOG\"",
                        },
                        new Aws.Glue.Inputs.GetScriptDagNodeArgArgs
                        {
                            Name = "database",
                            Value = $"\"{aws_glue_catalog_database.Destination.Name}\"",
                        },
                        new Aws.Glue.Inputs.GetScriptDagNodeArgArgs
                        {
                            Name = "table_name",
                            Value = $"\"{aws_glue_catalog_table.Destination.Name}\"",
                        },
                    },
                    Id = "resolvechoice3",
                    NodeType = "ResolveChoice",
                },
                new Aws.Glue.Inputs.GetScriptDagNodeArgs
                {
                    Args = 
                    {
                        new Aws.Glue.Inputs.GetScriptDagNodeArgArgs
                        {
                            Name = "database",
                            Value = $"\"{aws_glue_catalog_database.Destination.Name}\"",
                        },
                        new Aws.Glue.Inputs.GetScriptDagNodeArgArgs
                        {
                            Name = "table_name",
                            Value = $"\"{aws_glue_catalog_table.Destination.Name}\"",
                        },
                    },
                    Id = "datasink4",
                    NodeType = "DataSink",
                },
            },
            Language = "PYTHON",
        }));
        this.PythonScript = example.Apply(example => example.PythonScript);
    }

    [Output("pythonScript")]
    public Output<string> PythonScript { get; set; }
}

