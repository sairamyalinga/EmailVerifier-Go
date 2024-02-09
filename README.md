# Email Verifier

## Overview
This is an email verifier tool implemented in Go. It verifies the validity of a domain name by checking its MX records, SPF records, and DMARC records. 

## Features
- Domain validation: Checks if the domain exists by verifying its MX records.
- SPF validation: Verifies the presence of SPF (Sender Policy Framework) records for the domain.
- DMARC validation: Checks for DMARC (Domain-based Message Authentication, Reporting, and Conformance) records associated with the domain.

## How to Use
1. Clone or download this repository to your local machine.
2. Ensure you have Go installed on your system.
3. Open a terminal and navigate to the directory where the `EmailVerifier-Go` file is located.
4. Run the program by executing the command: `go run main.go`.
5. Enter the domain name you want to verify when prompted.
6. The program will display the verification results, including the presence of MX, SPF, and DMARC records for the domain.

## References
- [Cloudflare - Understanding DNS MX Records](https://www.cloudflare.com/learning/dns/dns-records/dns-mx-record/)
- [Cloudflare - Understanding DNS SPF Records](https://www.cloudflare.com/learning/dns/dns-records/dns-spf-record/)
- [Cloudflare - Understanding DNS DMARC Records](https://www.cloudflare.com/learning/dns/dns-records/dns-dmarc-record/)
- [Youtube](https://www.youtube.com/watch?v=9E4UEsWpYvM&list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9&index=7&t=143s)

