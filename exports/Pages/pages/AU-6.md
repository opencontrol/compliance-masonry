---
permalink: /NIST-800-53/AU-6/
title: AU-6 - Audit Review, Analysis, and Reporting
parent: NIST-800-53 Documentation
---

## Loggregator
a. - GSA Order CIO P 2100.1 states audit records must be reviewed frequently for signs of unauthorized activity or other security events.
- By default, Loggregator streams logs to a terminal. 18F will drain logs to a third-party log management service such as AlienVault USM and AWS CloudTrail 
- Cloud Foundry logs are captured in multiple tables and log files.  These will be reviewed weekly and if discovery of anomalous audit log content which appears to indicate a breach are handled according to the GSA Security Incident Handling Guide: CIO IT Security 01-02 Revision 7 (August 18, 2009) requirements.  
 
### References

* [Loggregator code](https://github.com/cloudfoundry/loggregator)

* [Cloud Foundry Logging](https://docs.cloudfoundry.org/running/managing-cf/logging.html)

### Governors

* [Loggregator Diagram](https://raw.githubusercontent.com/cloudfoundry/loggregator/develop/docs/loggregator.png)

--------
