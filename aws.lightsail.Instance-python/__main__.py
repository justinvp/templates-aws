import pulumi
import pulumi_aws as aws

# Create a new GitLab Lightsail Instance
gitlab_test = aws.lightsail.Instance("gitlabTest",
    availability_zone="us-east-1b",
    blueprint_id="string",
    bundle_id="string",
    key_pair_name="some_key_name",
    tags={
        "foo": "bar",
    })

