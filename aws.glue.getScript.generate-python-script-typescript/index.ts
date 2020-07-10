import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.all([aws_glue_catalog_database_source.name, aws_glue_catalog_table_source.name, aws_glue_catalog_database_destination.name, aws_glue_catalog_table_destination.name, aws_glue_catalog_database_destination.name, aws_glue_catalog_table_destination.name]).apply(([aws_glue_catalog_database_sourceName, aws_glue_catalog_table_sourceName, aws_glue_catalog_database_destinationName, aws_glue_catalog_table_destinationName, aws_glue_catalog_database_destinationName1, aws_glue_catalog_table_destinationName1]) => aws.glue.getScript({
    dagEdges: [
        {
            source: "datasource0",
            target: "applymapping1",
        },
        {
            source: "applymapping1",
            target: "selectfields2",
        },
        {
            source: "selectfields2",
            target: "resolvechoice3",
        },
        {
            source: "resolvechoice3",
            target: "datasink4",
        },
    ],
    dagNodes: [
        {
            args: [
                {
                    name: "database",
                    value: `"${aws_glue_catalog_database_sourceName}"`,
                },
                {
                    name: "table_name",
                    value: `"${aws_glue_catalog_table_sourceName}"`,
                },
            ],
            id: "datasource0",
            nodeType: "DataSource",
        },
        {
            args: [{
                name: "mapping",
                value: "[(\"column1\", \"string\", \"column1\", \"string\")]",
            }],
            id: "applymapping1",
            nodeType: "ApplyMapping",
        },
        {
            args: [{
                name: "paths",
                value: "[\"column1\"]",
            }],
            id: "selectfields2",
            nodeType: "SelectFields",
        },
        {
            args: [
                {
                    name: "choice",
                    value: "\"MATCH_CATALOG\"",
                },
                {
                    name: "database",
                    value: `"${aws_glue_catalog_database_destinationName}"`,
                },
                {
                    name: "table_name",
                    value: `"${aws_glue_catalog_table_destinationName}"`,
                },
            ],
            id: "resolvechoice3",
            nodeType: "ResolveChoice",
        },
        {
            args: [
                {
                    name: "database",
                    value: `"${aws_glue_catalog_database_destinationName1}"`,
                },
                {
                    name: "table_name",
                    value: `"${aws_glue_catalog_table_destinationName1}"`,
                },
            ],
            id: "datasink4",
            nodeType: "DataSink",
        },
    ],
    language: "PYTHON",
}, { async: true }));

export const pythonScript = example.pythonScript;

