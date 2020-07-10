import pulumi
import pulumi_aws as aws

example = aws.msk.Configuration("example",
    kafka_versions=["2.1.0"],
    server_properties="""auto.create.topics.enable = true
delete.topic.enable = true

""")

