import pulumi
import pulumi_aws as aws

example_table = aws.dynamodb.Table("exampleTable",
    attributes=[{
        "name": "UserId",
        "type": "S",
    }],
    hash_key="UserId",
    read_capacity=1,
    write_capacity=1)
example_role = aws.iam.Role("exampleRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "appsync.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}

""")
example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
    policy=example_table.arn.apply(lambda arn: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Action": [
        "dynamodb:*"
      ],
      "Effect": "Allow",
      "Resource": [
        "{arn}"
      ]
    }}
  ]
}}

"""),
    role=example_role.id)
example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
example_data_source = aws.appsync.DataSource("exampleDataSource",
    api_id=example_graph_ql_api.id,
    dynamodb_config={
        "table_name": example_table.name,
    },
    service_role_arn=example_role.arn,
    type="AMAZON_DYNAMODB")

