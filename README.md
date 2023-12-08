# P2P

This repo is a simple poc around p2p.
It uses a tracker that keep track of peer's IP.

The tracker could be improved in few ways :
- it could remove inactive peer from the list.
    This could be achieved either by oftenly pinging known peer or by having peer maintaining tcp connection.
- it could have a peer limit and send back an error when max peer reached

The peer could be  improved in few ways :
- it could have a list of known trackers and try connecting to another peer if tracker has reached max peer.
- it could stop only when user decided to (interrupt signal - at the moment the second peer stop after stepping through every action and stop the tcp server)

## Run

Prerequesites :
- [Taskfile](https://taskfile.dev)
- Docker

You can run the poc by  :

1) cloning it 
2) run build tasks in `tracker` and `peer` folders
3) run `run` task in `tracker`
4) run `run-1` task in `peer`
5) run `run-2` task in `peer`
