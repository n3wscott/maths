## TODO



```
cat <<EOF | kubectl apply -f -
apiVersion: maths.tableflip.dev/v1alpha1
kind: Add
metadata:
  name: add-one
spec:
  add:
  - value: 1
EOF
```

```
cat <<EOF | kubectl apply -f -
apiVersion: maths.tableflip.dev/v1alpha1
kind: Add
metadata:
  name: add-few
spec:
  add:
  - value: 1
  - value: 2
  - value: 3
  - value: 4
EOF
```

```
cat <<EOF | kubectl apply -f -
apiVersion: maths.tableflip.dev/v1alpha1
kind: Add
metadata:
  name: add-refs
spec:
  add:
    - value: 1
    - ref:
        name: add-one
        namespace: default
        kind: Add
        apiVersion: maths.tableflip.dev/v1alpha1
    - ref:
        name: add-few
        namespace: default
        kind: Add
        apiVersion: maths.tableflip.dev/v1alpha1
EOF
```

```
kubectl delete add add-one
kubectl delete add add-few
kubectl delete add add-refs
```