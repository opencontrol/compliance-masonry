---
permalink: /NIST-800-53/RA-5/
title: RA-5 - Vulnerability Scanning
parent: NIST-800-53 Documentation
---

## AlienVault
AlienVault USM for AWS runs AWS friendly Authenticated vulnerability scans within the AWS infrastructure
### References

* [AlienVault Website](http://www.alienvault.com)

--------

## Amazon Machine Images
a. - 18F runs Nessus (Authenticated) scans of the Cloud Foundry environment weekly based on IP ranges in use. These scans include network discovery and vulnerability checks of operating systems, server software, and any supporting components or applications. Scans are automatically compared to previous scans to identify new vulnerabilities or changes which resolve previously identified vulnerabilities. Nessus reports are reviewed at least weekly and appropriate actions taken on discovery of vulnerabilities.
- Nessus is used to run (Authenticated) scans when a new host/AMI/Stemcell is built.  This scan determines baseline posture used to contribute to decision of Production acceptance.  Additionally, this tool is used to execute CIS benchmark compliance scans when actively working to address configuration and hardening requirements.
- OWASP Zap is used to conduct web Application scanning primarily for the OWASP Top 10. 18F uses it as an integrated security testing tool for finding vulnerabilities in web applications. 18F will provide more automated functionally of security tests using OWASP ZAP and Jenkins for its  software development lifecycle and continuous integration functions
- AlienVault USM for AWS runs AWS friendly Authenticated vulnerability scans within the 18F AWS infrastructure and does not require permission from AWS to run scan within its Virtual Private Cloud (VPC)
 
 
### References

* [Amazon Machine Images](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIs.html)

--------

## Nessus
a. - Nessus will be used conduct internal scanning of its VPC and private subnets within the 18F Virtual Private Cloud.
- 18F runs Nessus (Authenticated) scans of the Cloud Foundry environment weekly based on IP ranges in use. These scans include network discovery and vulnerability checks of operating systems, server software, and any supporting components or applications. Scans are automatically compared to previous scans to identify new vulnerabilities or changes which resolve previously identified vulnerabilities. Nessus reports are reviewed at least weekly and appropriate actions taken on discovery of vulnerabilities.
- Nessus is used to run (Authenticated) scans when a new host/AMI/Stemcell is built.  This scan determines baseline posture used to contribute to decision of Production acceptance.  Additionally, this tool is used to execute CIS benchmark compliance scans when actively working to address configuration and hardening requirements.
 
 b. - Nessus, and AlienVault USM for AWS utilize tools and techniques that promote interoperability such as Common Vulnerability Scoring System v2 (CVSS2), Common Platform Enumeration (CPE), and Common Vulnerability Enumeration (CVE). Tenable SecurityCenter is able to output reports in CyberScope format to meet NIST, DHS, and GSA reporting requirements.
 
 c. - AlienVault USM for AWS, OWASP Zap and Tenable Nessus reports are reviewed and analyzed at least weekly and appropriate actions taken on discovery of vulnerabilities within the 18F Cloud Infrastructure and applications and from security control assessments conducted on its information systems.
 
 
### References

* [Nessus Website](http://www.tenable.com/products/nessus-vulnerability-scanner)

--------

## OWASP Zap
a. OWASP Zap is used to conduct web Application scanning primarily 
for the OWASP Top 10. 18F uses it as an integrated security testing tool for finding vulnerabilities in web applications. 18F will provide more automated functionally of security tests using OWASP ZAP and Jenkins for its 
software development lifecycle and continuous integration functions.
 
 b. OWASP Zap is used for web application scanning of hosted websites 
and web based applications. It scans for the OWASP TOP 10 vulnerabilities 
and utilize tools and techniques that promote interoperability such 
as Common Vulnerability Scoring System v2 (CVSS2), Common Platform 
Enumeration (CPE), and Common Vulnerability Enumeration (CVE). 
 
 c. OWASP Zap reports are reviewed and analyzed at least weekly and 
appropriate actions taken on discovery of vulnerabilities within 
the 18F Cloud Infrastructure and applications and from security 
control assessments conducted on its information systems.
 
 
### References

* [OWASP Zap Site](https://www.owasp.org/index.php/OWASP_Zed_Attack_Proxy_Project)

--------
