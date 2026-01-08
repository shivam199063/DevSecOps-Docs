---
title: "Code"
description: "GitOps-managed DevOps, Security, and Development codebases"
type: docs
toc: false
---

This section represents the **source of truth** for all implementations in this repository.

## Principles

- Declarative infrastructure
- GitOps-first deployments
- Reproducible via Minikube
- Production-style layouts

## Structure

- **DevOps** → Platform & Kubernetes tooling
- **Security** → Wazuh, Nuclei, DefectDojo, hardening
- **Development** → Go & automation services

> All deployments are reconciled via Argo CD from Git.
