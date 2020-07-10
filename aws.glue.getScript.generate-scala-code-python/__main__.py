import pulumi
import pulumi_aws as aws

example = aws.glue.get_script(dag_edges=[
        {
            "source": "datasource0",
            "target": "applymapping1",
        },
        {
            "source": "applymapping1",
            "target": "selectfields2",
        },
        {
            "source": "selectfields2",
            "target": "resolvechoice3",
        },
        {
            "source": "resolvechoice3",
            "target": "datasink4",
        },
    ],
    dag_nodes=[
        {
            "args": [
                {
                    "name": "database",
                    "value": f"\"{aws_glue_catalog_database['source']['name']}\"",
                },
                {
                    "name": "table_name",
                    "value": f"\"{aws_glue_catalog_table['source']['name']}\"",
                },
            ],
            "id": "datasource0",
            "node_type": "DataSource",
        },
        {
            "args": [{
                "name": "mappings",
                "value": "[(\"column1\", \"string\", \"column1\", \"string\")]",
            }],
            "id": "applymapping1",
            "node_type": "ApplyMapping",
        },
        {
            "args": [{
                "name": "paths",
                "value": "[\"column1\"]",
            }],
            "id": "selectfields2",
            "node_type": "SelectFields",
        },
        {
            "args": [
                {
                    "name": "choice",
                    "value": "\"MATCH_CATALOG\"",
                },
                {
                    "name": "database",
                    "value": f"\"{aws_glue_catalog_database['destination']['name']}\"",
                },
                {
                    "name": "table_name",
                    "value": f"\"{aws_glue_catalog_table['destination']['name']}\"",
                },
            ],
            "id": "resolvechoice3",
            "node_type": "ResolveChoice",
        },
        {
            "args": [
                {
                    "name": "database",
                    "value": f"\"{aws_glue_catalog_database['destination']['name']}\"",
                },
                {
                    "name": "table_name",
                    "value": f"\"{aws_glue_catalog_table['destination']['name']}\"",
                },
            ],
            "id": "datasink4",
            "node_type": "DataSink",
        },
    ],
    language="SCALA")
pulumi.export("scalaCode", example.scala_code)

