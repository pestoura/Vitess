## Known Issues

* A critical vulnerability [CVE-2021-44228](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-44228) in the Apache Log4j logging library was disclosed on Dec 9 2021.
  The project provided release `2.15.0` with a patch that mitigates the impact of this CVE. It was quickly found that the initial patch was insufficient, and additional CVEs
  [CVE-2021-45046](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-45046) and [CVE-2021-44832](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-44832) followed.
  These have been fixed in release `2.17.1`. This release of Vitess, `v10.0.2`, uses a version of Log4j below `2.17.1`, for this reason, we encourage you to use version `v10.0.5` instead, to benefit from the vulnerability patches.

* An issue where the value of the `-force` flag is used instead of `-keep_data` flag's value in v2 vreplication workflows (#9174) is known to be present in this release. A workaround is available in the description of issue #9174.

## Bug fixes
### Query Serving
 * Fixes encoding of sql strings #8029
 * Fix for issue with information_schema queries with both table name and schema name predicates #8099
 * PRIMARY in index hint list for release 10.0 #8159
### VReplication
 * VReplication: Pad binlog values for binary() columns to match the value returned by mysql selects #8137
## CI/Build
### Build/CI
 * update release notes with known issue #8081
## Documentation
### Other
 * Post v10.0.1 updates #8045
## Enhancement
### Build/CI
 * Added release script to the makefile #8030
### Other
 * Add optional TLS feature to gRPC servers #8176
## Other
### Build/CI
 * Release 10.0.1 #8031

The release includes 14 commits (excluding merges)
Thanks to all our contributors: @GuptaManan100, @askdba, @deepthi, @harshit-gangal, @noxiouz, @rohit-nayak-ps, @systay
