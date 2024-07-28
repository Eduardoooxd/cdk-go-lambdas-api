
# CDK Go Lambdas API

## Overview
This project leverages the AWS Cloud Development Kit (CDK) with Go to deploy Lambda functions.

## Structure
- `cdk.json`: Configures CDK execution.
- `lambda/`: Contains Lambda function code.
- `aws-go-cdk-course.go`: Infrastructure for lambdas with CDK.

## Commands
- `cdk deploy`: Deploys stack to AWS.
- `cdk diff`: Compares deployed stack with local state.
- `cdk synth`: Generates CloudFormation template.
- `go test`: Runs unit tests.

## Requirements
- AWS CLI
- AWS CDK
- Go

## Usage
1. **Install dependencies**:
   ```sh
   go mod tidy
   ```
2. **Deploy stack**:
   ```sh
   cdk deploy
   ```

## License
MIT
