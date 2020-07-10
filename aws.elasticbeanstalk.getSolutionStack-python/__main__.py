import pulumi
import pulumi_aws as aws

multi_docker = aws.elasticbeanstalk.get_solution_stack(most_recent=True,
    name_regex="^64bit Amazon Linux (.*) Multi-container Docker (.*)$")

