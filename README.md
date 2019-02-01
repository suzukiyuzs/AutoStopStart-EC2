# AutoStopStart EC2 - AWS Serverless Application Model Golang & holiday_jp-go

## Architecture
- Lambda
  - AutoStopStart-EC2
- Cloudwatch Schedule
  - AutoStart
  - AutoStot
- IAM
  - AutoStopStart-EC2-IamRole
  - AutoStopStart-EC2-IamPolicy (InlinePolicy)

## Compiling & Deploying
```
sam package --template-file template.yaml --output-template-file output.yaml --s3-bucket <<S3 Bucket>>
sam deploy --template-file output.yaml --stack-name <<Stack Name>> --capabilities CAPABILITY_NAMED_IAM
```
