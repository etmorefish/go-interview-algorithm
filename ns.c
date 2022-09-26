struct task_struct { 
  ...
/* namespaces */
struct nsproxy *nsproxy; 
    ...
}

struct nsproxy {
atomic_t count;
struct uts_namespace *uts_ns; struct ipc_namespace *ipc_ns; struct mnt_namespace *mnt_ns; struct pid_namespace
*pid_ns_for_children; struct net *net_ns;
}