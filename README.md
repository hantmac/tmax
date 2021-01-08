                 _
                | |_ _ __ ___   __ ___  __
                | __| '_ ` _ \ / _` \ \/ /
                | |_| | | | | | (_| |>  <
                 \__|_| |_| |_|\__,_/_/\_\

### What is `tmax`?

The positioning of `tmax` is a command line tool with a little artificial intelligence. 
If you frequently deal with the terminal daily, tmax will greatly improve your work efficiency.

### The design idea of tmax

- Have a local storage as a knowledge base
- Efficient search algorithm and accurate feedback
- Full command line interaction
- Make your very long cmd short

### How tmax works?
Firstly, you need to `tmax generate` to generate a config file in `$HOME/.tmax.yaml` and the file look like as follows:
```yaml
custom:
- check: curl 127.0.0.1:8080/health
k8s:
- filternodecpu: kubectl get nodes -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.status.capacity.cpu}{'\t'}{.status.capacity.memory}{'\n'}{end}"
- filternodetaint: kubectl get nodes -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.spec.taints[*].key}{'\n'}{end}"
- corednsedit: kubectl edit cm coredns -nkube-system
- allnode: kubectl get no
- alldeploy: kubectl get deploy
- allpod: kubectl get pod -A
- busyboxrun: kubectl run busybox --rm -ti --image=busybox /bin/sh
- allnodeip: kubectl get node -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.status.addresses[0].address}{'\n'}{end}"
- podResource: kubectl get pod -o custom-columns=NAME:metadata.name,podIP:status.podIP,hostIp:spec.containers[0].resources
- getns: kubectl get ns
- createdemo: kubectl create deployment nginx --image=nginx
- exposedemo: kubectl expose deployment nginx --port=80
- pronacos: nacos.zmlearn.com
- getnetwork: kubectl get networkpolicy
- runbox: kubectl run busybox --rm -ti --image=busybox /bin/sh
unix:
- "tar": "tar -xjvf test.tar.bz2"

```
As you can see, there are many long commands those hard to keep in your mind. 
If you want to quickly get a long command, even if you have memorized it, it takes a long time to type it into the console, 
not to mention that sometimes you can't remember such a long command.

At this moment, `tmax` appeared, it will solve the problem just mentioned.


### What will tmax bring me?

`tmax` has 3 mode: directly mode, search mode and interactive mode. And `tmax` will make your very long terminal cmd short, improve your operation efficiency.

#### directly mode
If you clearly know the key you want to execute the command, you can use directly mode.
Use 'tmxa somekey' , example: `tmax check` will execute `curl 127.0.0.1:8080/health`

![](https://tva1.sinaimg.cn/large/008eGmZEgy1gmgj2gk1n7g31u40u0wr0.gif)


#### search mdoe

If you know the general content of the command you want to execute, 
you can use the search mode to find and execute it.

Use `tmax s CONTENTOFCMD`, example: `tmax s pod` or `tmax search pod`.

![](https://tva1.sinaimg.cn/large/008eGmZEgy1gmgjtnunfzg31qp0u0x3w.gif)

#### interactive search mode
If you don't want to search, tmax has interactive mode.
Just type `tmax` and `enter` to interactive mode.

![](https://tva1.sinaimg.cn/large/008eGmZEgy1gmgk55zgudg31tp0u0x2o.gif)