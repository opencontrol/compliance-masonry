# Cloud.gov and NIST 800-53 Revision 4 Controls
Cloud.gov facilitates authorization (ATO) of government applications by fulfilling critical NIST 800-53 Revision 4 controls.

### Baseline assembly
Cloud.gov automatically deploys applications to an authorized hardened baseline of Ubuntu 14.04 LTS.

 * CM-2 Baseline configuration
 * CM-3 Configuration Change Control
 * CM-6 Configuration Settings

### Infrastructure as a service
Cloud.gov's infrastructure implements best practices and has been authorized.

* SC-7 Boundary protection
* AC-3 Access enforcement
* AC-6 Least privilege

### Continuous integration and testing
Security testing and infrastructure updates are all handled by Cloud.gov's team.

* CA-8 Penetration testing
* RA-5 Vulnerability Scanning
* SA-11
  * (1) Developer Security Testing and Evaluation | Static Code Analysis
* SI-2 Flaw Remediation
* SI-10 Information Input Validation

### HTTPS Everywhere
Applications deployed on Cloud.gov are automatically configured for https.

* SC-13 Cryptographic protection
* SC-28
  * (1) Protection of Information At Rest | Cryptographic Protection: applicable to systems with Sensitive Personally Identifiable Information Only

### Authorization and authentication
Cloud.gov uses [Cloudfoundry's authorization and authentication system](https://github.com/cloudfoundry/uaa) to manage access to application deployment and services.

* AC-2 Account Management
* IA-2 Identification and Authentication (Organizational Users)
  * (1) Identification and Authentication (Organizational Users) | Network Access to Privileged Accounts
  * (2) Identification and Authentication (Organizational Users) | Network Access to Non-Privileged Accounts
  * (12) Identification and Authentication | Acceptance of PIV Credentials: consult with DevOps/CyberSec for Applicability

### Monitoring
Cloud.gov uses New Relic to monitor application statuses. Additionally, user and event logs events are stored.

* AU-2 Audit Events
* AU-6 Audit Review, Analysis, and Reporting
* SI-4 Information System Monitoring

### Version control
Applications on Cloud.gov can only be deployed using GIT.

* CM-8 Information system component inventory
