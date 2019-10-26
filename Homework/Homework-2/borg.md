# Borg

## Availability

### what problem to solve

Failures are the norm in large scale systems.

### how to solve

+ automatically reschedules evicted tasks, on a new machine if necessary
+ reduces correlated failures by spreading tasks of a job across failure domains such as machines, racks, and power domains
+ limits the allowed rate of task disruptions and the number of tasks from a job that can be simultaneously down during maintenance activities such as OS or machine upgrades
+ uses declarative desired-state representations and idempotent mutating operations, so that a failed client can harmlessly resubmit any forgotten requests
+ rate-limits ﬁnding new places for tasks from machines that become unreachable, because it cannot distinguish between large-scale machine failure and a network partition
+ avoids repeating task-machine pairings that cause task or machine crashes
+ recovers critical intermediate data written to local disk by repeatedly re-running a log saver task, even if the alloc it was attached to is terminated or moved to another machine

### how is the effect

Enabled to achieve 99.99% availability in practice.

## Utilization

### what problem to solve

Make efﬁcient use of Google’s ﬂeet of machines, which represents a signiﬁcant ﬁnancial investment: increasing utilization by a few percentage points can save millions of dollars.

### how to solve

+ evaluation methodology
+ cell sharing
+ large cells
+ fine-grained resource requests
+ resource reclamation

## Isolation

### what problem to solve

Although sharing machines between applications increases utilization, it also requires good mechanisms to prevent tasks from interfering with one another. This applies to both security and performance.

### how to solve

+ security isolation

  use a Linux chroot jail as the primary security isolation mechanism between multiple tasks on the same machine

+ performance isolation

  all Borg tasks run inside a Linux cgroup-based resource container and the Borglet manipulates the container settings, giving much improved control because the OS kernel is in the loop

*[more about the essay](<https://pdos.csail.mit.edu/6.824/papers/borg.pdf>)*