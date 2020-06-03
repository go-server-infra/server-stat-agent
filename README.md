# server-stat-agent

An agent that collects the server information and sends it back to master

This agent is written in beautiful GO, and the image size is ~7mb. It's super fast and lightweight

This deamon sends json to the target url, Here is a sample JSON

```json
{
  "host": "hostname",
  "stat": {
    "mem_total": 3142250496,
    "mem_used": 584495104,
    "mem_cached": 1593884672,
    "mem_free": 1005064192,
    "cpu_user": 0.12706480304955528,
    "cpu_system": 0.6353240152477764,
    "cpu_idle": 99.23761118170266
  }
}
```

There are plans to send more information (disk size) e.t.c

## Launch


Launch is via docker.

```bash
 docker run -v /proc/stat:/proc/stat -e HOST_NAME=hostname -e REPORT_URL=http://thehost.com/report nchanged/stat-agent
```

If you're working with linux, don't forget to mount the volume


