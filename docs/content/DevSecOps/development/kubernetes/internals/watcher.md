---
title: "Watcher"
Description: "WATCH in Kubernetes"
type: docs
toc: false
weight: 30
---


### What is WATCH?

In Kubernetes, **WATCH** means asking the API Server:

> “Keep listening and tell me whenever something changes.”

Instead of repeatedly querying the API Server, WATCH opens a **long-lived connection** and receives events in real time.

These events describe **changes**, not full history.

---

### Why WATCH exists

Kubernetes clusters are highly dynamic:
- Pods are created and deleted frequently
- Configurations change
- Nodes join and leave

Polling the API Server again and again would:
- Waste resources
- Add latency
- Miss fast changes

WATCH solves this by making Kubernetes **event-driven**.

---

### Simple example

When you run:

```bash
kubectl get pods -w
```
- **kubectl starts a WATCH request.**<br>
Now, whenever a Pod:<br>
    - is created
    - is modified
    - is deleted

    kubectl immediately prints the update.

### WATCH events
A WATCH stream sends events of these types:<br>
- ADDED – a new object was created
- MODIFIED – an existing object changed
- DELETED – an object was removed
- ERROR – something went wrong<br>

Each event contains:
- The event type
- The latest version of the object

### What WATCH does NOT give you

- Objects that already existed before the watch started
- A full snapshot of the current state
- Historical data (If an object already exists and does not change, WATCH will never send it.)


### WATCH and resourceVersion
WATCH always starts from a specific `resourceVersion`.<br>
This tells Kubernetes:-
`Send me all changes that happen after this point.`<br><br>
Using resourceVersion ensures:
- No missed events
- Correct ordering of updates

### WATCH in Kubernetes architecture

```
Client / Controller
        ↓
     WATCH request
        ↓
     API Server
        ↓
   Event stream
```
The API Server pushes events as changes occur.

> WATCH tells you what changes in the future, but nothing about what already exists.