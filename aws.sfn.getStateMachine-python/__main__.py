import pulumi
import pulumi_aws as aws

example = aws.sfn.get_state_machine(name="an_example_sfn_name")

