---
title: "List"
Description: "LIST in Kubernetes"
type: docs
toc: false
weight: 20
---


### What is LIST?

In Kubernetes, **LIST** means asking the API Server:

> “Show me all the objects of this type that exist right now.”

It returns a **snapshot of the current state** at a specific point in time.

For example:
- All Pods in a namespace
- All Services in the cluster
- All Nodes currently registered

---

### Why LIST is important

Kubernetes is a dynamic system. Objects may already exist **before** a component starts running.

LIST answers a very basic but critical question:

> “What already exists in the cluster?”


---

### Simple example

When you run:

```bash
kubectl get pods
```
Internally, kubectl sends a LIST request to the API Server.<br>
- **The API Server responds with:**<br>
All existing Pods,
Their specifications,
Their current status,
A **resourceVersion** which is representing the state of the cluster at that moment<br>
This is a one-time response, not a live feed.<br>

- **LIST does not provide:**<br>
Real-time updates,
Notifications for future changes,
Information about objects created after the request.

### **LIST and race conditions**<br>
If a component relies only on LIST:<br>
*It reads the current state<br>
The cluster changes immediately after<br>
The component now has stale information<br>
This is why LIST alone is not enough for `controllers or automation.`*

### **LIST in the Kubernetes architecture**
```
Client / Controller
        ↓
     LIST request
        ↓
     API Server
        ↓
        etcd
```
The API Server always reads the current state from etcd and returns it to the client.

> LIST tells you what exists right now, but nothing about what will happen next.