# KubeGazer

Kubegazer is a lightweight, high-performance Kubernetes cluster observability dashboard built with Go and `client-go`. 

It provides real-time node resource tracking, pod status aggregation, and workload diagnostics without the overhead of heavy metrics infrastructure or third-party agent dependencies.

---

## Features (v0.1 MVP - In Active Development)

- **Node Resource Normalization:** Server-side aggregation of CPU millicores (`MilliValue`) and RAM capacity (`Value` in bytes) to offload parsing overhead from the client UI.
- **Accurate Workload Status Evaluation:** Handles complex pod states, prioritizing `DeletionTimestamp` (Terminating), container `Waiting` states (`CrashLoopBackOff`, `ImagePullBackOff`), and top-level `Phase` fallbacks.
- **Pure `corev1` Dependency:** Operates directly against standard Kubernetes API servers—no mandatory reliance on `metrics-server` or `metrics.k8s.io` for core cluster state inspection.
- **Idiomatic Go API Server:** Built using standard library `http.ServeMux` routing, contextual lifecycle management (`r.Context()` request cancellation), and `kubernetes.Interface` abstractions for testability.

---

## Architecture

```text
[ Browser / UI ] 
       │ (REST / JSON)
       ▼
[ pkg/api ] ── HTTP Server & Request Handlers
       │
[ pkg/k8s ] ── Domain DTOs & Resource Translators (Nodes / Pods)
       │
[ k8s.io/client-go ]
       │
       ▼
[ Kubernetes API Server ]

much more to come!    