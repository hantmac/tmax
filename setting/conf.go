package setting


var TmaxDefaultConf = `
---
k8sDemo:
- get: kubectl get -o template pod/web-pod-13je7 --template={{.status.phase}} */* kubectl get pod test-pod -o custom-columns=CONTAINER:.spec.containers[0].name,IMAGE:.spec.containers[0].image */* kubectl get -o json pod web-pod-13je7
- create: kubectl create -f ./pod.json */* cat pod.json | kubectl create -f - */* kubectl create -f docker-registry.yaml --edit -o json 
- label: kubectl label pods foo unhealthy=true */* kubectl label --overwrite pods foo status=unhealthy  */* kubectl label pods foo bar- 
- run: kubectl run nginx --image=nginx */* kubectl run hazelcast --image=hazelcast/hazelcast --port=5701 */* kubectl run -i -t busybox --image=busybox --restart=Never
- cp: kubectl cp /tmp/foo_dir <some-pod>:/tmp/bar_dir */* kubectl cp /tmp/foo <some-pod>:/tmp/bar -c <specific-container> */* kubectl cp /tmp/foo <some-namespace>/<some-pod>:/tmp/bar
- taint: kubectl taint nodes foo dedicated=special-user:NoSchedule */* kubectl taint nodes foo dedicated:NoSchedule- */* kubectl taint nodes foo dedicated-
k8s:
- filternodecpu: kubectl get nodes -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.status.capacity.cpu}{'\t'}{.status.capacity.memory}{'\n'}{end}"
- filternodetaint: kubectl get nodes -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.spec.taints[*].key}{'\n'}{end}"
- corednsedit: kubectl edit cm coredns -nkube-system
- allnode: kubectl get no
- alldeploy: kubectl get deploy
- allpod: kubectl get pod -A
- getns: kubectl get ns
- busyboxrun: kubectl run busybox --rm -ti --image=busybox /bin/sh
- allnodeip: kubectl get node -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.status.addresses[0].address}{'\n'}{end}"
- podResource: kubectl get pod -o custom-columns=NAME:metadata.name,podIP:status.podIP,hostIp:spec.containers[0].resources | grep 8Gi
custom:
- check: curl 127.0.0.1:8080/@in/api/health
unix:
- 'tar': tar -xvjf test.tar.bz2
- 'netutils': yum install net-tools -y
`