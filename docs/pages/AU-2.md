---
permalink: /AU-2/
title: AU-2 - Audit Events
---
## a  
* * *   
### AWS Auditable Events  
* 18F has implemented AWS CloudTrail for its account monitoring. It provides visibility into user activity by recording API calls made on an AWS account. CloudTrail records important information about each API call, including the name of the API, the identity of the caller, the time of the API call, the request parameters, and the response elements returned by the AWS service. This information helps 18F track changes made to its AWS resources and to troubleshoot operational issues.  
* CloudTrail delivers API call information by depositing log files in an Amazon S3 bucket.  Each log file can contain multiple events, and each event represents an API call.  
  
### Cloud Foundry Auditable Events  
* Security Event Logging and Auditing  
* For operators, Cloud Foundry provides an audit trail through the bosh tasks command. This command shows all actions that and operator has taken with the platform. Additionally, operators can redirect Cloud Foundry component logs to a standard syslog server using the syslog_daemon_config property in the metron_agent job of cf-release.  
* For users, Cloud Foundry records an audit trail of all relevant API invocations of and app. The CLI command cf events returns this information.  
* Loggregator/doppler, the Cloud Foundry component responsible for logging, provides a stream of log output from hosted applications and from Cloud Foundry system components that interact with applications during updates and execution.  
  
### Operating System (EC2) (EBS) (VPC) Auditable Events  
* 18F has implemented AWS CloudTrail for monitoring Amazon EC2, Amazon EBS, and Amazon VPC  that captures API calls and delivers the log files to an Amazon S3 bucket. By default, all CloudTrail log files are encrypted by using Amazon S3 server-side encryption (SSE)  
* When CloudTrail logging is enabled, API calls made to Amazon EC2, Amazon EBS, and Amazon VPC actions are tracked in log files, along with any other AWS service records. CloudTrail determines when to create and write to a new file based on a specified time period and file size.  
* Every log entry contains information about who generated the request. The user identity information in the log helps you determine whether the request was made with root or IAM user credentials, with temporary security credentials for a role or federated user, or by another AWS service.  
* CloudTrail publish Amazon SNS notifications when new log files are delivered for 18F staff to review upon log file delivery.  
* 18F implements AlienVault USM for AWS to correlate all AWS audit logs for analysis and review.  
  
## b  
* * *   
### AWS Auditable Events  
* All 18F Event audit logs will be made available as needed to other organizational entities for mutual support and selection of events to be audited.  
  
### Cloud Foundry Auditable Events  
* Event audit logs will be made available as needed to other organizational entities for mutual support and selection of events to be audited.  
* This will make audit log review more effective and ensure that user activities are captured with the data that is available.  
  
## c  
* * *   
## d  
* * *   
### AWS Auditable Events  
* 18F has implemented AWS CloudTrail for monitoring Amazon EC2, Amazon EBS, and Amazon VPC  that captures API calls and delivers the log files to an Amazon S3 bucket. By default, all CloudTrail  log files are encrypted by using Amazon S3 server-side encryption (SSE)  
* When CloudTrail logging is enabled, API calls made to Amazon EC2, Amazon EBS, and Amazon VPC actions are tracked in log files, along with any other AWS service records. CloudTrail determines when to create and write to a new file based on a specified time period and file size.  
* Every log entry contains information about who generated the request. The user identity information in the log helps  determine whether the request was made with root or IAM user credentials, with temporary security credentials for a role or federated user, or by another AWS service.  
* CloudTrail publish Amazon SNS notifications when new log files are delivered for 18F staff to review upon log file delivery.  
  
### Cloud Foundry Auditable Events  
* For operators, Cloud Foundry provides an audit trail through the bosh tasks command. This command shows all actions that and operator has taken with the platform. Additionally, operators can redirect Cloud Foundry component logs to a standard syslog server using the syslog_daemon_config property in the metron_agent job of cf-release.  
* For users, Cloud Foundry records and audit trail of all relevant API invocations of and app. The CLI command cf events returns this information.  
* Loggregator, the Cloud Foundry component responsible for logging, provides a stream of log output from 18F applications and from Cloud Foundry system components that interact with your app during updates and execution.  
* By default, Loggregator streams logs to a terminal. 18F will drain logs to a third-party log management service such as AlienVault USM and AWS Cloudtrail.  
* Cloud Foundry gathers and stores logs in a best-effort manner. If a client is unable to consume log lines quickly enough, the Loggregator buffer may need to overwrite some lines before the client has consumed them. a syslog drain or a CLI tail can usually keep up with the flow of application logs.  
  
