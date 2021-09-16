# sniff
A Simple Golang Tool That Automates OSINT For Threat Intelligence And Mapping Your Attack Surface.

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

## Installation:
```
$ git clone https://github.com/krishpranav/sniff
$ cd sniff 
$ go run sniff.go
```

## Usage:
```
  -t string
        Domain or IP address (Required)
  -p string
        Print results to console
```

## Output:
```
$ sniff -t google.com -p

[?] ASN: "15169" ORG: "GOOGLE, US"
8.8.4.0/24
--- sniff ---
[:] Writing 616 CIDRs to file...
[:] Converting to IPs...
8.8.8.1
--- sniff ---
[:] Writing 14725936 IPs to file...
[!] Done.
```
