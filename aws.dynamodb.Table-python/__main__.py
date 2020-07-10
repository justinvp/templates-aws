import pulumi
import pulumi_aws as aws

basic_dynamodb_table = aws.dynamodb.Table("basic-dynamodb-table",
    attributes=[
        {
            "name": "UserId",
            "type": "S",
        },
        {
            "name": "GameTitle",
            "type": "S",
        },
        {
            "name": "TopScore",
            "type": "N",
        },
    ],
    billing_mode="PROVISIONED",
    global_secondary_indexes=[{
        "hash_key": "GameTitle",
        "name": "GameTitleIndex",
        "nonKeyAttributes": ["UserId"],
        "projectionType": "INCLUDE",
        "range_key": "TopScore",
        "read_capacity": 10,
        "write_capacity": 10,
    }],
    hash_key="UserId",
    range_key="GameTitle",
    read_capacity=20,
    tags={
        "Environment": "production",
        "Name": "dynamodb-table-1",
    },
    ttl={
        "attributeName": "TimeToExist",
        "enabled": False,
    },
    write_capacity=20)

