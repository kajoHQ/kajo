# KAJO: Kubernetes Asynchronous Jobs Orchestrator

![KAJO](https://github.com/adhaamehab/KAJO/assets/13816742/af0ed9d2-14d7-42d0-bd02-e355cb4cb941)

Kubernetes-native platform bringing simplicity and efficiency to the orchestration and management of Jobs and Workers. This project is currently in the early stages of development and is looking to reshape the way we manage asynchronous tasks within Kubernetes environments.

## Why KAJO?

Built with Go, KAJO provides a scalable, highly available, and user-friendly solution designed to meet the needs of modern cloud-native applications. The key features of KAJO include:

### Job and CronJob Management
Easily create, update, and delete Kubernetes Jobs and CronJobs using a user-friendly interface or directly via the REST API.

### Monitoring and Logs
Monitor the status of your Jobs and CronJobs and directly access their logs from the user interface or via the REST API.

### Scalability and High Availability
Both the API server and the user interface are designed to be highly scalable and available. They are deployed as Kubernetes Deployments with multiple replicas for reliability and performance.

### Fault Tolerance
With system state stored in Kubernetes' etcd, KAJO offers resilience against server crashes or restarts.

### Security
KAJO operates entirely within your Kubernetes cluster, providing an inherent level of security and control.

### Idempotence
Tasks are designed to yield the same results, irrespective of the number of executions, to ensure consistency.

### Exactly Once Execution
KAJO guarantees that each task is executed only once, preventing duplicate executions and ensuring data integrity.

### Callbacks and Hooks
Define pre and post-execution hooks for tasks, allowing for custom logic to be executed before or after a job runs.

### Triggers
Users can define triggers that automatically initiate tasks based on specific events or conditions.

### Dependency Management
KAJO supports the specification of dependencies between tasks, ensuring they are executed in the correct order.

### Priority-Based Execution
Assign priority levels to tasks to influence the order in which they are executed.

### Patterns Templates
KAJO provides different patterns to choose from when creating tasks, simplifying the process of building distributed and large-scale systems.

## User Interfaces
KAJO provides multiple interfaces for interacting with the system:

**User Interface (UI)**: A user-friendly UI for managing tasks and monitoring their status.

**Command Line Interface (CLI)**: A CLI for managing tasks and monitoring their status.

**REST API**: A REST API for managing tasks and monitoring their status.


## Architecture

<img width="1457" alt="image" src="https://github.com/adhaamehab/KAJO/assets/13816742/eb09bad7-1d30-4ddf-b89f-b53f06b7b2ba">

