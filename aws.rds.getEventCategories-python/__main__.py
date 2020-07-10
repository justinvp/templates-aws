import pulumi
import pulumi_aws as aws

example_event_categories = aws.rds.get_event_categories()
pulumi.export("example", example_event_categories.event_categories)

