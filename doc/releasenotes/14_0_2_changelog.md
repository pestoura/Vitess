# Changelog of Vitess v14.0.2

### Bug fixes
#### Backup and Restore
* BugFix: Using backups from v13 should work with v14 [#10919](https://github.com/vitessio/vitess/pull/10919@)
#### Build/CI
* Fix Upgrade Downgrade manual backup dependencies install [#11084](https://github.com/vitessio/vitess/pull/11084)
#### Cluster management
* [vtctld] Fix nil-ness in healthcheck (#11067) [#11068](https://github.com/vitessio/vitess/pull/11068)
#### Query Serving
* Online DDL: deal with autogenerated CHECK constraint names [#10638](https://github.com/vitessio/vitess/pull/10638)
* Fix AST copying of basic types in release-14.0 [#11051](https://github.com/vitessio/vitess/pull/11051)
* fix: return when instructions are nil in checkThatPlanIsValid [#11070](https://github.com/vitessio/vitess/pull/11070)
* On demand heartbeats: fix race condition closing the writer (#11157) [#11158](https://github.com/vitessio/vitess/pull/11158)
### CI/Build
#### Build/CI
* Remove MariaDB 10.2 Unit Test in v15 [#10700](https://github.com/vitessio/vitess/pull/10700)
* Add more robust go version handling [#11001](https://github.com/vitessio/vitess/pull/11001)
* Fix mariadb103 ci [#11015](https://github.com/vitessio/vitess/pull/11015)
* CI: change upgrade/downgrade tests to use vitessio fork of go-junit-report (#11023) [#11024](https://github.com/vitessio/vitess/pull/11024)
* Add workflow file to the filter rules [#11032](https://github.com/vitessio/vitess/pull/11032)
* Add upgrade-downgrade tests for next releases (release-14.0) [#11034](https://github.com/vitessio/vitess/pull/11034)
* Remove deprecated flag `enable_queries` from end to end tests in v14  [#11049](https://github.com/vitessio/vitess/pull/11049)
* Add SKIP check to workflows removed in v15 [#11083](https://github.com/vitessio/vitess/pull/11083)
#### Cluster management
* Fix flakiness in TestReplicationStatus [#10927](https://github.com/vitessio/vitess/pull/10927)
#### Query Serving
* CI Fix: Collation tests (#10839) [#10842](https://github.com/vitessio/vitess/pull/10842)
#### VReplication
* v14 backport of two PRs: vreplication throttling info & vreplication endtoend flakiness fix [#10928](https://github.com/vitessio/vitess/pull/10928)
### Enhancement
#### Backup and Restore
* Backport: Backups: Support InnoDB Redo Log Location With 8.0.30+ [#10895](https://github.com/vitessio/vitess/pull/10895)
### Release
#### Documentation
* Addition of the release summary for v14.0.2 [#11142](https://github.com/vitessio/vitess/pull/11142)
#### General
* Post release `v14.0.1` [#10850](https://github.com/vitessio/vitess/pull/10850)
* Include the compose examples in the `do_release` script [#11130](https://github.com/vitessio/vitess/pull/11130)
* Upgrade go version to `1.18.5` on `release-14.0` [#11131](https://github.com/vitessio/vitess/pull/11131)

