name: {{.Name}}
on: [push, pull_request]
concurrency:
  group: format('{0}-{1}', ${{"{{"}} github.ref {{"}}"}}, '{{.Name}}')
  cancel-in-progress: true

jobs:
  build:
    name: Run endtoend tests on {{.Name}}
    {{if .Ubuntu20}}runs-on: ubuntu-20.04{{else}}runs-on: ubuntu-18.04{{end}}

    steps:
    - name: Check if workflow needs to be skipped
      id: skip-workflow
      run: |
        skip='false'
        if [[ "{{"${{github.event.pull_request}}"}}" ==  "" ]] && [[ "{{"${{github.ref}}"}}" != "refs/heads/main" ]] && [[ ! "{{"${{github.ref}}"}}" =~ ^refs/heads/release-[0-9]+\.[0-9]$ ]] && [[ ! "{{"${{github.ref}}"}}" =~ "refs/tags/.*" ]]; then
          skip='true'
        fi
        echo Skip ${skip}
        echo "::set-output name=skip-workflow::${skip}"

    - name: Set up Go
      if: steps.skip-workflow.outputs.skip-workflow == 'false'
      uses: actions/setup-go@v2
      with:
        go-version: 1.17.13

    - name: Set up python
      if: steps.skip-workflow.outputs.skip-workflow == 'false'
      uses: actions/setup-python@v2

    - name: Tune the OS
      if: steps.skip-workflow.outputs.skip-workflow == 'false'
      run: |
        echo '1024 65535' | sudo tee -a /proc/sys/net/ipv4/ip_local_port_range

    - name: Check out code
      if: steps.skip-workflow.outputs.skip-workflow == 'false'
      uses: actions/checkout@v2

    - name: Get dependencies
      if: steps.skip-workflow.outputs.skip-workflow == 'false'
      run: |
        sudo apt-get update
        sudo apt-get install -y mysql-server mysql-client make unzip g++ etcd curl git wget eatmydata
        sudo service mysql stop
        sudo service etcd stop
        sudo ln -s /etc/apparmor.d/usr.sbin.mysqld /etc/apparmor.d/disable/
        sudo apparmor_parser -R /etc/apparmor.d/usr.sbin.mysqld
        go mod download

        # install JUnit report formatter
        go get -u github.com/vitessio/go-junit-report@HEAD

        {{if .InstallXtraBackup}}

        wget https://repo.percona.com/apt/percona-release_latest.$(lsb_release -sc)_all.deb
        sudo apt-get install -y gnupg2
        sudo dpkg -i percona-release_latest.$(lsb_release -sc)_all.deb
        sudo apt-get update
        sudo apt-get install percona-xtrabackup-24

        {{end}}

    {{if .MakeTools}}

    - name: Installing zookeeper and consul
      if: steps.skip-workflow.outputs.skip-workflow == 'false'
      run: |
          make tools

    {{end}}

    - name: Run cluster endtoend test
      if: steps.skip-workflow.outputs.skip-workflow == 'false'
      timeout-minutes: 30
      run: |
        source build.env

        set -x

        # run the tests however you normally do, then produce a JUnit XML file
        eatmydata -- go run test.go -docker=false -follow -shard {{.Shard}}
