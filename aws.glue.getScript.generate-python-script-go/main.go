package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "PYTHON"
		example, err := glue.GetScript(ctx, &glue.GetScriptArgs{
			DagEdges: []glue.GetScriptDagEdge{
				glue.GetScriptDagEdge{
					Source: "datasource0",
					Target: "applymapping1",
				},
				glue.GetScriptDagEdge{
					Source: "applymapping1",
					Target: "selectfields2",
				},
				glue.GetScriptDagEdge{
					Source: "selectfields2",
					Target: "resolvechoice3",
				},
				glue.GetScriptDagEdge{
					Source: "resolvechoice3",
					Target: "datasink4",
				},
			},
			DagNodes: []glue.GetScriptDagNode{
				glue.GetScriptDagNode{
					Args: []glue.GetScriptDagNodeArg{
						glue.GetScriptDagNodeArg{
							Name:  "database",
							Value: fmt.Sprintf("%v%v%v", "\"", aws_glue_catalog_database.Source.Name, "\""),
						},
						glue.GetScriptDagNodeArg{
							Name:  "table_name",
							Value: fmt.Sprintf("%v%v%v", "\"", aws_glue_catalog_table.Source.Name, "\""),
						},
					},
					Id:       "datasource0",
					NodeType: "DataSource",
				},
				glue.GetScriptDagNode{
					Args: []glue.GetScriptDagNodeArg{
						glue.GetScriptDagNodeArg{
							Name:  "mapping",
							Value: "[(\"column1\", \"string\", \"column1\", \"string\")]",
						},
					},
					Id:       "applymapping1",
					NodeType: "ApplyMapping",
				},
				glue.GetScriptDagNode{
					Args: []glue.GetScriptDagNodeArg{
						glue.GetScriptDagNodeArg{
							Name:  "paths",
							Value: "[\"column1\"]",
						},
					},
					Id:       "selectfields2",
					NodeType: "SelectFields",
				},
				glue.GetScriptDagNode{
					Args: []glue.GetScriptDagNodeArg{
						glue.GetScriptDagNodeArg{
							Name:  "choice",
							Value: "\"MATCH_CATALOG\"",
						},
						glue.GetScriptDagNodeArg{
							Name:  "database",
							Value: fmt.Sprintf("%v%v%v", "\"", aws_glue_catalog_database.Destination.Name, "\""),
						},
						glue.GetScriptDagNodeArg{
							Name:  "table_name",
							Value: fmt.Sprintf("%v%v%v", "\"", aws_glue_catalog_table.Destination.Name, "\""),
						},
					},
					Id:       "resolvechoice3",
					NodeType: "ResolveChoice",
				},
				glue.GetScriptDagNode{
					Args: []glue.GetScriptDagNodeArg{
						glue.GetScriptDagNodeArg{
							Name:  "database",
							Value: fmt.Sprintf("%v%v%v", "\"", aws_glue_catalog_database.Destination.Name, "\""),
						},
						glue.GetScriptDagNodeArg{
							Name:  "table_name",
							Value: fmt.Sprintf("%v%v%v", "\"", aws_glue_catalog_table.Destination.Name, "\""),
						},
					},
					Id:       "datasink4",
					NodeType: "DataSink",
				},
			},
			Language: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("pythonScript", example.PythonScript)
		return nil
	})
}

