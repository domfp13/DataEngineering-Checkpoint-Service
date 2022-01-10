# checkpoint-service
General purpose golang API to create a checkpoint service on AWS S3

This project contains source code and supporting files for a containerized
golang application. It includes the following files and folders.

- src (Directory) - Go source code directory.- 
    - constants.go
      - These are constants a variable to be used across the src package.
    - jsonutil.go
        - Works with JSON files (reads, writes).
    - s3util.go
        - Works with downloading/uploading file objest to S3.
- docker-compose.test.yml - docker-compose file to be used while testing.
- docker-compose.yml - docker-compose file to be used in Production.
- Dockerfile - Docker image manifest.
- main.go - Application main runner, creates a webserver for REST API.
- Makefile - Unix utility that is designed to start an execution.

## Requirements
* [Docker](https://hub.docker.com/search/?type=edition&offering=community):
    * This application has been packaged and requires Docker to be run. All require dependencies
      e.g (Oracle client, AWS boto3 libs, pandas, numpy, etc) have been already configured in the main Docker
      image, this guaranties a very easy to manage and portable code across the Windows/Linux ecosystems.
* [Git CLI](https://git-scm.com/)
    * This project is using GNU Makefile, to test locally using Windows the Git Bash
      (terminal emulator) is required, also set the environmental variables binary
      for git on top of the system environmental variables, this will allow command shell
      to use the Unix instead of Windows.
* [AWS Credentials]()
    * For local development this project requires the aws credentials to be set.
    * Create under your home directory an .aws directory
        * ``` mkdir ~/.aws && cd ~/.aws ```
        * ``` touch config credentials ```
        * Edit the config and credentials files as follows:
```
[default]
region = us-east-1
```
```
[default]
aws_access_key_id=<YOUR_AWS_ACCESS_KEY_ID>
aws_secret_access_key=<YOUR_AWS_SECRET_ACCESS_KEY>
```
<div style="font-size:140%;color:red"> 
Passing credentials or storing AWS credentials in the code is prohibited.
</div>

## Run application
This repo contains a `Makefile` with different target formulas to build and run (locally/production).
To check the different target formulas run:
```
$ make
```
This will display the different targets.

To run this application in local environment first add the credentials in the .env file of this repo 
after that run:
```
$ make run-local
```
To run this application in QC/Prod run
```
$ make run-prod
```

## How to contribute to the project
There are couple off rules in order to contribute to this repo.
* #### Git
    * DO NOT merge to the main master branch.
    * Use the following to create Pull Requests (fix/<your_fix>, feat/<your_feature>, etc.)
    * Every Pull requests will require a review.
* #### Go Standards
    * [The Go Programming Language Specification](https://docs.gitlab.com/ee/development/go_guide/) Go Code style.
        * functions should be lower camel case.
        * Use typing such as str, int, list, etc.
        * Functions should be comment using the standards.
        * Add test to your code.
        * Follow the directory standard.
        * ETC

### Class Diagram
* TBA ...

## Author
* **Luis Enrique Fuentes Plata** - *2022-01-11*
