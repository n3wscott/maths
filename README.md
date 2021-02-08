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


```
cat <<EOF | kubectl apply -f -
apiVersion: maths.tableflip.dev/v1alpha1
kind: Subtract
metadata:
  name: sub-few
spec:
  sub:
  - value: 10
  - value: 5
  - value: 1
  - value: 2
EOF
```


```
cat <<EOF | kubectl apply -f -
apiVersion: tgik.tgik.io/v1
kind: Square
metadata:
  name: demo
spec:
  base: 5
EOF
```

```
cat <<EOF | kubectl apply -f -
apiVersion: maths.tableflip.dev/v1alpha1
kind: Add
metadata:
  name: mixed
spec:
  add:
    - value: 1
    - ref:
        name: sub-few
        namespace: default
        kind: Subtract
        apiVersion: maths.tableflip.dev/v1alpha1
    - ref:
        name: add-few
        namespace: default
        kind: Add
        apiVersion: maths.tableflip.dev/v1alpha1
    - ref:
        name: square-sample
        namespace: default
        kind: Square
        apiVersion: tgik.tgik.io/v1
EOF
```

