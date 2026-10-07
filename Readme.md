

### why do we need compute resources isolation within a container?
it's possible that a process has memory leak and it starts consuming so much memory/cpu.
And if there are no guardrails kernel can choose any other process for e.g. DB or ssh daemon
which is problematic.
cgroup is also used to limit number of process it can spinup in other words to prevent fork bomb.

### why do we need network isolation?
```bash
# Terminal 1
docker run --rm --network=host python:3.12 python3 -m http.server 8080

# Terminal 2
docker run --rm --network=host python:3.12 python3 -m http.server 8080
# OSError: [Errno 98] Address already in use
```
Both containers are fighting over the host's single port 8080. Now with isolation:

### why do we need UTS isolation ?
To isolate hostname

### why do we need IPC isolation ?
If not isolated one process could intervene with other process.
