# goquery

goquery is a remote investigation client that uses your existing osquery deployment to provide remote shell level functionality with fewer risks than SSH.

Using osquery's distributed API, hosts can be targeted for single queries to pull back specific information without having to modify the osquery schedule. goquery uses this API and abstracts it into a shell like experience that investigators are more used to while retaining all the power of osquery tables, extensions, and features like Auto Table Construction.

With goquery you can connect to hosts via UUID, and remotely interact with a host's osquery instance in an interactive session. This works over osquery's asynchronous distributed API. The concept of node keys, discovery queries, and the osquery schedule is abstracted away (or not used) to provide a clean, efficient way to remotely interrogate hosts for abuse, compromise investigation, or fleet management.

![goquery](https://user-images.githubusercontent.com/3303787/67834837-a8cbff00-faa5-11e9-8723-7526be0f63a5.png "goquery")

To get up and running, view the [building and running section](#running-goquery)

### Features:

- Interactive prompt with typeahead and help text (courtesy of [go-prompt](https://github.com/c-bata/go-prompt))
- [osctrl](https://github.com/jmpsec/osctrl) integration
- Command aliasing
- Print modes
- Both interactive and non interactive scheduling modes

# Commands

The following is a list of all goquery commands and their calling requirements.

### .connect \<UUID\>
This opens a session with a remote host. It will ask the backend if a host with that UUID is registered and if not return to the user saying it doesn't exist. If the backend returns that the host exists then a session is opened and that machine is set as the active host. All future commands will interact with this host until it's disconnected from or the user changes to another host. Supports suggestions.

### .disconnect \<UUID\>
Close a session with a remote host. Fails if you're not connected to a host with that UUID. Supports suggestions.

### .exit
Exit goquery. Shell state will not be saved but command history is.

### .help
Show goquery help formatted with the currently selected printing mode.

### .clear
Clear the terminal screen

### .history
Show all past queries in the current session for the current host.

### .hosts
Show all hosts you are connected to with their osquery version, hostname, UUID, and platform

### .mode \<print_mode\>
Change the printing mode. goquery supports multiple printing modes to help you make sense of data at a glance. We currently support: Line, JSON, and Pretty (default).

### .query \<query\>
Runs a query on a remote host and waits for the result before returning control to the REPL. Equivalent to running .schedule and .resume together.
![query_table_suggestion](https://user-images.githubusercontent.com/2386877/67360345-79077f00-f51a-11e9-8d12-c897818f992a.png "Query Table Suggestions")

### .resume \<query_name\>
This will either wait for a query to complete or fetch the results and display them if the query has already posted results. This is used in conjunction with .schedule to pull the results of queries that are running asynchronously. This can also be used to display the results of any previously run query.

### .schedule \<query\>
Run a query asynchronously on the remote host. The query will be tracked in the session for that host so results can be fetched at any point in time, but this allows the investigator to kick off a bunch of things without waiting for each one to complete first.

### .alias \<alias_name\> \<command\> \<interpolated_args\>
List current aliases when called with no arguments or flags. To create a new alias, call with `--add` flag and provide arguments as follows: `.alias --add ALIAS_NAME command_string`

Positional arguments with $# placeholders are interpolated when the command is run, for example the following alias `.all` with command `.query select * from $#` will evaluate to `.query select * from processes` when called with `.all processes`.

Command name must not contain any spaces in order to preserve the space delimited arguments

To remove an alias, use `.alias --remove ALIAS_NAME`

### cd \<dir\>
Change directories on a remote host. This affects other pseudo-commands like `ls`.

### ls
List the files in the current directory. The current directory is set by using the `cd` command and starts at `/`.

# Integration

To support the various features of goquery, your backend will need to support a number of APIs to interact with your fleet. The core APIs are required for basic functionality but future APIs may focus on more fringe features such as ATC, file pulling, etc. goquery can work without these APIs and that functionality will be disabled.

## Core API

The following endpoints are required to enable goquery to talk to a host's osquery instance. See `goserver/mock_osquery_server.go` for a reference implementation.

### checkHost
**Description:** Verify a host exists in the fleet.

**goquery Provides:** UUID

**goquery Expects:** Information on that UUID: if it exists or not, osquery version, hostname, operating system version.

---

### scheduleQuery
**Description:** Schedule a query on a remote machine.

**goquery Provides:** UUID, query

**goquery Expects:** A unique identifier for that query will be passed to fetchResults for updates on the query.

---

### fetchResults
**Description:** Pull the results of a query by the name returned from scheduleQuery

**goquery Provides:** queryName

**goquery Expects:** The query results if they are available

## Config

Goquery can be configured via a configuration json file. Debug mode, defaults, and aliases can be set in the structure of the provided `config.template.json`. Valid print modes are as follows "json", "line", and "pretty".

By default, goquery will check for a config file at the following path: `~/.goquery/config.json`. This can be overidden when calling the binary or running with the following flags: `--config ./path_to_file.json`

# Building and Running

### Docker Testing Infra
Hopefully one day goquery will be plug'n'play with the most popular osquery backends, but for now it'll take a little work to integrate. To get up and running playing with goquery as quickly as possible, you can use the docker test infra.

Running `make docker` will build a set of nodes used to create a simulated osquery deployment with two Ubuntu hosts, a central osquery server, along with a SAML IdP. goquery's docker infra contains its own osquery server written in Go which is designed to be lightweight and easy to understand to help you learn how to integrate goquery into your enterprise.

Deploy it locally with `make deploy` (which uses docker swarm) and then you're ready to start testing by running goquery.

### Running goquery

Use `go run cmd/main.go --config ./config.template.json` to simply run from the root of the directory, or build a binary if you wish with `go build -o goquery ./cmd/main.go `

For a quick demo, try the following commands:

- `.connect uuid`
- `ls`
- `.mode line`
- `.query select * from system_info`

## Slack
[![Slack Status](https://osquery-slack.herokuapp.com/badge.svg)](https://osquery-slack.herokuapp.com)

There is a goquery channel in the osquery slack for discussion and help. If you have questions about deployment, integration, or usage, you can ask there.

## Bugs?
File an issue!

## License

goquery is licensed under the [MIT License](https://raw.github.com/AbGuthrie/goquery/master/LICENSE).
