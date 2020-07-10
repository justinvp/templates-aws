import pulumi
import pulumi_aws as aws

example = aws.transfer.get_server(server_id="s-1234567")

