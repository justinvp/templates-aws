import pulumi
import pulumi_aws as aws

config = pulumi.Config()
file_system_id = config.get("fileSystemId")
if file_system_id is None:
    file_system_id = ""
by_id = aws.efs.get_file_system(file_system_id=file_system_id)

