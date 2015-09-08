---
permalink: /AC-2/
title: AC-2 - Account Management
---
### AWS Account Management  
* AWS accounts are managed through AWS Identity and Access Management (IAM). Only users with a need to operate the AWS management console are provided individual AWS user accounts. There are no guest/anonymous, group, or temporary user accounts in AWS. Please see the Amazon Web Services US East/West and GovCloud System Security Plan for additional information.  
* 18F users must sign in to the account's signin URL by using their IAM user name and password. This signin URL is located in the Dashboard of the IAM console and must be communicated by the  18F AWS account's system administrator to the IAM user.  
* AWS Cloud Foundry security groups are documented in section 9.3 Types of Users.  
  
### Cloud Foundry PaaS Account Management  
* Cloud Foundry user and role accounts are managed and maintained through the  UAA Command Line Interface (UAAC). Cloud Foundry uses role-based access control with each role granting permissions in either a organization or an application space.  
* Cloud Foundry uses Application security groups (ASGs) to act as virtual firewalls to control outbound traffic from the applications in the deployment. A security group consists of a list of network egress access rules.  
* An administrator can assign one or more security groups to a Cloud Foundry deployment or to a specific space in an org within a deployment  
* Cloud Foundry user, organization, and application roles and security groups are documented in section 9.3 Types of Users.  
  
### Operating System Account (EC2) Management  
* Access to Amazon EC2 linux instances are managed by the use of EC2 key pairs and using SSH to access the local instance on the individual Linux, or appliance instance. Account types include individual user and system/application user accounts. Shared or group accounts are not permitted outside of default accounts such as local Administrators or root. There are no guest/anonymous or temporary user accounts.  
* Operating System user groups are documented in section 9.3 Types of Users.  
* Initial linux local root access is provided to AWS administrator account users only if they provide the key pair assigned to the linux EC2 instance and login using SSH.  
  
### 18F Account Management  
* 18F assigns account managers from the DevOps team for information system accounts  
  
### AWS Account Management  
* Conditions for group membership in AWS groups are based on an established need to manage the AWS environment. 18F doenst have No non-administrator accounts currently exist in AWS.  
  
### Cloud Foundry PaaS Account Management  
* Conditions for group membership for Cloud Foundry user accounts are based on requesting access within Github 18F tracking and ticketing system.  Once approved, the designated Cloud Foundry Administrators create the Cloud Foundry contact, add it to the appropriate organization.  After these steps are completed, the user can be assigned to projects and programs related to Cloud Foundry.  
* Account managers will be assigned for account groups in all environments, determined by the Cloud Foundry Project Manager and/or the Information System Owner.  
  
### Operating System Account (EC2) Management  
* Conditions for group membership for EC2  and local system user accounts are based on an established need to manage servers in the Cloud Foundry environment. User group membership is restricted to the least privilege necessary for the user to accomplish his or her assigned duties.  
  
### AWS Account Management  
* AWS user accounts are issued only to those with an established need to manage the AWS environment. No non-administrator accounts currently exist in AWS.  
* 18F and GSA identify authorized users of the information system and specify access rights/privileges. 18F grants access to the information system based on - (i) a valid need-to-know/need-to-share that is determined by assigned official duties and satisfying all personnel security criteria; (ii) Intended system usage. 18F and GSA require proper identification for requests to establish information system accounts and approves all such requests; and (iii) Organizational or mission/business function attributes.  
  
### Cloud Foundry PaaS Account Management  
* Cloud Foundry user accounts are issued only to those who have a business need and meet the conditions for group membership to Cloud Foundry user accounts. Accounts are created and removed based on requests sent within 18 F's Github tracking and ticketing system.  Once approved, the designated Cloud Foundry Administrators create/remove the Cloud Foundry contact, for the appropriate organization.  After these steps are completed, the Cloud Foundry user can be assigned to or removed  from projects and programs related to Cloud Foundry.  
* 18F and GSA identify authorized users of the information system and specify access rights/privileges. 18F grants access to the information system based on - (i) a valid need-to-know/need-to-share that is determined by assigned official duties and satisfying all personnel security criteria; (ii) Intended system usage. 18F and GSA requires proper identification for requests to establish information system accounts and approves all such requests; and (iii) Organizational or mission/business function attributes.  
  
### Operating System Account (EC2) Management  
* Linux and/or local system user accounts are issued only to those with an established need to manage servers in the 18F AWS environment. User group membership is restricted to the least privilege necessary for the user to accomplish his or her assigned duties.  
* 18F and GSA identify authorized users of the information system and specify access rights/privileges. 18F grants access to the information system based on - (i) a valid need-to-know/need-to-share that is determined by assigned official duties and satisfying all personnel security criteria; (ii) Intended system usage. 18F and GSA requires proper identification for requests to establish information system accounts and approves all such requests; and (iii) Organizational or mission/business function attributes.  
  
### AWS Account Management  
* AWS user account creation requires approval by the managing 18F project lead and Cloud Foundry Information System Technical Point of Contact (Operating Environment). Prior to account creation users must have a signed NDA and have at least begun the GSA background investigative process.  
  
### Cloud Foundry PaaS Account Management  
* Cloud Foundry user account creation requires approval by the Approving Authority. Prior to account creation users must have a signed NDA and have at least begun the GSA background investigative process.  
  
### Operating System Account (EC2) Management  
* Local system user account creation requires approval by the managing 18F project lead and Cloud Foundry Information System Technical Point of Contact (Operating Environment).  Prior to account creation users must have a signed NDA and have at least begun the GSA background investigative process.  
  
### 18F Account Management  
* All accounts will be documented within their respective information systems, detailing their group and role membership, and access authorizations. This documentation will be exported by the System Engineers and/or DevOps and archived for up to a year from the date of account creation by the managing 18F project lead and Cloud Foundry Information System Technical Point of Contact (Operating Environment) in accordance with best business and security practices.  
  
### AWS Account Management  
* AWS user account establishment, activation, modification, disablement or removal requires approval by the managing 18F project lead and Cloud Foundry Information System Technical Point of Contact (Operating Environment).  
  
### Cloud Foundry PaaS Account Management  
* Cloud Foundry user account establishment, activation, modification, disablement or removal  requires approval by their respective account group managers, Cloud Foundry Project Manager, the Information System Owner or Cloud Foundry Information System Technical Point of Contact.  
  
### Operating System Account (EC2) Management  
* Local system user account establishment, activation, modification, disablement or removal requires approval by the managing 18F project lead and Cloud Foundry Information System Technical Point of Contact (Operating Environment).  
  
### 18F Account Management  
* 18F requires approvals by the 18F project lead and 18F system owners for requests to create information system accounts.  
  
### AWS Account Management  
* AWS Accounts will be Created, Enabled, Modified, Disabled, and Removed from AWS in accordance with policies, guidelines, requirements in NIST SP 800-37, and established by the 18F  Project Lead, and or  Cloud Foundry Information System Owner.  
  
### Cloud Foundry PaaS Account Management  
* Cloud Foundry User Accounts will be Created, Enabled, Modified, Disabled, and Removed from Cloud Foundry in accordance with policies, guidelines, requirements in NIST SP 800-37, and established by the 18F Project Lead, and or Cloud Foundry Information System Owner.  
* All accounts for all environments will be Created, Enabled, Modified, Disabled, and Removed by their respective account group managers, 18F Project Manager, and/or the Information System Owner.  
  
### Operating System Account (EC2) Management  
* Operating System Accounts will be Created, Enabled, Modified, Disabled, and Removed from Operating Systems in accordance with policies, guidelines, requirements in NIST SP 800-37, and established by the 18F Project Lead, and or Cloud Foundry Information System Owner.  
  
### 18F Account Management  
* 18F creates, enables, modifies, disables, and removes information system accounts in accordance with the 18F access control policy,  account access procedures and approvals from 18F DevOps.  
  
### AWS Account Monitoring  
* 18F has implemented AWS CloudWatch for its account monitoring. It allows 18F to monitor AWS resources in near real-time, including Amazon EC2 instances, Amazon EBS volumes, Elastic Load Balancers, and Amazon RDS DB instances. Metrics such as CPU utilization, latency, and request counts are provided automatically for these AWS resources. It allows 18F to supply logs or custom application and system metrics, such as memory usage, transaction volumes, or error rates.  
* Amazon CloudWatch functionality is accessible via API, command-line tools, the AWS SDK, and the AWS Management Console.  
  
### Cloud Foundry PaaS Account Management  
* 18F uses Steno for Cloud Foundry account logging. Steno is a shared logging service between the DEA, Cloud Controller The Cloud Foundry components share a common interface for configuring logs.  
* Loggregator/doppler, the Cloud Foundry component responsible for logging, provides a stream of log output from 18F applications and from Cloud Foundry system components that interact with  applications during updates and execution.  
  
### Operating System and network (EC2) (EBS) (ELB) Monitoring  
* 18F has implemented AWS CloudWatch  for basic monitoring of  Amazon EC2 instances. Basic Monitoring for Amazon EC2 instances- Seven pre-selected metrics at five-minute frequency and three status check metrics at one-minute frequency.  
* [object Object]  
* 18F has also implemented the use Auto Scaling and Elastic Load Balancing where Amazon CloudWatch provides Amazon EC2 instance metrics aggregated by Auto Scaling group and by Elastic Load Balancers.  
* Monitoring data is retained for two weeks, even if AWS resources have been terminated. This enables 18F to quickly look back at the metrics preceding an event of interest.  
* Basic Monitoring is already enabled automatically for all Amazon EC2 instances, and these metrics are accessed in either the Amazon EC2 tab or the Amazon CloudWatch tab of the AWS Management Console, or by using the Amazon CloudWatch API.  
  
### 18F Account Management  
* 18F Monitors the use of all information system accounts  
  
### AWS Account Management  
* AWS user accounts will be deactivated immediately on receipt of notification of an email from the managing 18F project lead or at a future date as directed.  
* User accounts will be monitored automatically on a continuous basis: manually on a monthly basis  and accounts will be disabled after 90 days of inactivity.  
* 18F will use AWS CloudTrail, OS Query and AlienVault USM to automatically monitor and notify when users account on AWS have been inactive for 90 days.  
  
### Cloud Foundry PaaS Account Management  
* Cloud Foundry user accounts will be deactivated immediately on receipt of notification of an email or through Github from the managing 18 project lead/system owner or at a future date as directed.  
  
### Operating System Account (EC2) Management  
* Local system user accounts will be deactivated immediately on receipt of notification of an email from the managing 18F project lead or at a future date as directed.  
* User accounts will be monitored monthly and accounts will be disabled after 90 days of inactivity.  
  
### 18F Account Management  
* For all environments, Account Managers, the 18F Project Manager, and/or Cloud Foundry Information System Owner will be notified when accounts are no longer required, accounts are terminated or transferred, and when individual system usage or need to know changes. Notification will be achieved via electronic notification or other official means.  
  
### AWS Account Management  
* System access to the AWS system components is provided only to those with an established need to manage the AWS environment. No non-administrator accounts currently exist in AWS.  
* 18F and GSA identify authorized users of the information system and specify access rights/privileges. 18F grants access to the information system based on - (i) a valid need-to-know/need-to-share that is determined by assigned official duties and satisfying all personnel security criteria; (ii) Intended system usage. 18F and GSA requires proper identification for requests to establish information system accounts and approves all such requests; and (iii) Organizational or mission/business function attributes.  
  
### Cloud Foundry PaaS Account Management  
* Cloud Foundry user accounts are issued only to those who submitted a request through Github and gained approval by 18F DevOps . Once approved, the Administrator creates the Cloud Foundry account access and adds it to the appropriate organization, role  and issue a license.  After these steps are completed, the Cloud Foundry user can be assigned to projects and programs.  
* 18F and GSA identify authorized users of the information system and specify access rights/privileges.  18F grants access to the information system based on: a valid need-to-know/need-to-share that is determined by assigned official duties and satisfying all personnel security criteria; Intended system usage. 18F and GSA requires proper identification for requests to establish information system accounts and approves all such requests; and Organizational or mission/business function attributes.  
  
### Operating System Account (EC2) Management  
* System access to local instance accounts is provided only to those with an established need to manage servers in the Cloud Foundry environment. User group membership is restricted to the least privilege necessary for the user to accomplish their assigned duties.  
* 18F and GSA identify authorized users of the information system and specify access rights/privileges. 18F grants access to the information system based on - (i) a valid need-to-know/need-to-share that is determined by assigned official duties and satisfying all personnel security criteria; (ii) Intended system usage. 18F and GSA requires proper identification for requests to establish information system accounts and approves all such requests; and (iii) Organizational or mission/business function attributes.  
* 18F and GSA identify authorized users of the information system and specify access rights/privileges. 18F grants access to the information system based on (i) a valid need-to-know/need-to-share that is determined by assigned official duties and satisfying all personnel security criteria; (ii) Intended system usage. 18F and GSA requires proper identification for requests to establish information system accounts and approves all such requests; and (iii) Organizational or mission/business function attributes.  
  
### 18F Account Management  
* 18F authorizes access to its information systems based on a valid access authorization from system owners and DevOps, intended system usage within the 18F network environment, and other attributes as required by the organization or associated missions/business functions  
  
### AWS Account Management  
* User accounts will be monitored monthly and accounts will be disabled after 90 days of inactivity; this will be a manual review process every 30 days. 18F is in the process of automating this thought the use of implementing AWS OSQuery with AlienVault USM for AWS  
* A review of all user accounts will be conducted on an annual basis  
  
### Cloud Foundry Account Management  
* User accounts will be monitored monthly and accounts will be disabled after 90 days of inactivity; this will be a manual review process every 30 days, but the disablement will be automatic.  
  
### Operating System (EC2) Account Management  
* User accounts will be monitored monthly and accounts will be disabled after 90 days of inactivity.  
* Linux accounts will be monitored via scripts which query the last logon date/time of each user account and provide the results in the form of a CSV file which an authorized administrator will use for the basis of disablement.  
* Appliance accounts will be monitored monthly and accounts will be disabled after 90 days of inactivity; this will be a manual review process, but the disablement will be automatic.  
  
### 18F Account Management  
* 18F reviews user and system accounts for compliance with account management requirements at least on an annual basis.  Currently, accounts are being monitored manually on a monthly basis and programmatically within an ongoing basis.  
  
### AWS Account Management  
* The root administrator AWS account password will be reset when a member of the group is removed or annually, whichever comes first.  
* 18F does not allow shared/group account credentials within its environment  
  
### Cloud Foundry Account Management  
* Cloud Foundry application administrator account passwords will be reset when a member of the group is removed or annually, whichever comes first.  
  
### Operating System (EC2) Account Management  
* Root Linux and local administrator account passwords will be reset when a member of the group is removed or annually, whichever comes first.  
  
### 18F Account Management  
* 18F establishes a process for reissuing shared/group account credentials (if deployed) when individuals are removed from the group.  
  
