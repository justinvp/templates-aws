import pulumi
import pulumi_aws as aws

example_connection = aws.directconnect.Connection("exampleConnection",
    bandwidth="1Gbps",
    location="EqSe2")
example_link_aggregation_group = aws.directconnect.LinkAggregationGroup("exampleLinkAggregationGroup",
    connections_bandwidth="1Gbps",
    location="EqSe2")
example_connection_association = aws.directconnect.ConnectionAssociation("exampleConnectionAssociation",
    connection_id=example_connection.id,
    lag_id=example_link_aggregation_group.id)

