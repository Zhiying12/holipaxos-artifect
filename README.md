# holipaxos-artifect

## Directory
1. [holipaxos](https://github.com/Zhiying12/holipaxos): The HoliPaxos implementation. Its instruction is inside the folder.
2. multipaxos: The MultiPaxos implementation.
3. raft-kv-store: The etcd's Raft implementation.
4. Copilot: The Copilot Paxos implementation.
5. omnipaxos-kv-store: The OmniPaxos KV store implementation.
6. ycsb: The YCSB benchmark close-loop client.
7. ycsb-openloop: The updated open-loop client on top of original YCSB benchmark.

## Installation
Please skip this part if you are using the provided AWS AMI instance, as environment and scriptes are set on the instance.

### Prerequisites
This artifect needs [rocksdb](https://github.com/facebook/rocksdb/blob/master/INSTALL.md), [maven](https://maven.apache.org/install.html), [docker](https://docs.docker.com/engine/install/), [go](https://go.dev/doc/install), and [rust](https://www.rust-lang.org/tools/install).
### Install
For HoliPaxos:
```
pushd holipaxos
go build -o bin/replicant main/main.go
docker build -t holipaxos .
popd
```
For MultiPaxos:
```
pushd multipaxos
go build -o bin/replicant main/main.go
docker build -t multipaxos .
popd
```
For raft and raft+:
```
pushd raft-kv-store
go build -o bin/replicant main.go
popd
```
For Copilot:
```
pushd copilot
go build -o bin/master master/master.go
go build -o bin/server master/server.go
docker build -t copilot .
popd
```
For OmniPaxos:
```
pushd omnipaxos-kv-store
cargo build --release
popd
```


## Run
The entire experiment requires 7 instances. 6 of them must be m5.2xlarge size (8 CPU and 32 GB RAM), and the rest one serves as a experiment controller which can be a smaller instance.

Replace the IP address for each machine in the 3-node.json and 5-nodee.json repectively. Then run the script, which will broadcast the json file to all machines and start experiments.