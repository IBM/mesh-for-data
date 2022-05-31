# Cleanup

When you're finished experimenting with a sample, you may clean up as follows:

1. Stop `kubectl port-forward` processes (e.g., using `pkill kubectl`)
2. Delete the namespace created for this sample:
    ```bash
    kubectl delete namespace fybrik-notebook-sample
    ```
3. Delete the policy created in the fybrik-system namespace:
    ```bash
    NS="fybrik-system"; kubectl -n $NS get configmap | awk '/sample/{print $1}' | xargs  kubectl delete -n $NS configmap
    ```
