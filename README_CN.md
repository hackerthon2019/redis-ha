---
typora-root-url: .
typora-copy-images-to: ./images
---

# 问题

Redis哨兵是Redis官方给出的高可用主从方案，使用该方案可以部署一个容错的高可用Redis集群，故障转移过程是自动化地进行的，而不需要人为干涉。

> Redis Sentinel provides high availability (HA) for Redis. In practical terms this means that using Sentinel you can create a Redis deployment that tolerates certain kinds of failures without human intervention. For more information about Redis Sentinel refer to: <https://redis.io/topics/sentinel>.

典型的哨兵架构图如下所示：

![image-20190419165707181](images/image-20190419165707181.png)

它由两部分组成，哨兵节点和数据节点：

- 哨兵节点：哨兵系统由一个或多个哨兵节点组成，哨兵节点是特殊的 Redis 节点，不存储数据。
- 数据节点：主节点和从节点都是数据节点。
- 客户端节点：Redis客户端通过sentinel协议感知到当前的主从节点信息，再连接后端的数据节点存取数据。

上面的架构我们能看到一个问题，这个方案对于客户端是不透明的，需要客户端支持sentinel协议以感知主从节点信息。这个对于有些场景来说意味着要修改客户端的Redis的驱动程序，因此整个方案在实施时有一些困难。

> Connecting an application to a Sentinel-managed Redis deployment is usually done with a Sentinel-aware Redis client. While most Redis clients do support Sentinel, the application needs to call a specialized connection management interface of the client to use it. When one wishes to migrate to a Sentinel-enabled Redis deployment, she/he must modify the application to use Sentinel-based connection management. Moreover, when the application uses a Redis client that does not provide support for Sentinel, the migration becomes that much more complex because it also requires replacing the entire client library.

现在社区中的一些helm chart，如[redis-ha](<https://github.com/helm/charts/tree/master/stable/redis-ha>)，部署的redis集群就是上面那种方案，因此存在的问题是类似的。

# 解决方案

为了解决上述问题，我们这里采用Redis官方给的[sentinel_tunnel](<https://github.com/RedisLabs/sentinel_tunnel>)作为Redis SmartProxy，以屏蔽下层的Redis集群状态细节，让客户端以普通Redis协议直接连接过来，架构图如下：

```
+-------------------------------------------+                                                           _,-'*'-,_
| +---------------------------------------+ |                                               _,-._      (_ o v # _)
| |                           +--------+  | |  +----------+       +----------+          _,-'  *  `-._  (_'-,_,-'_)
| |Application code           | Redis  |  | |  | Sentinel | +     |  Redis   | +       (_  O     #  _) (_'|,_,|'_)
| |(uses regular connections) | client +<------>+  Tunnel  +<----->+ Sentinel +<--+---->(_`-._ ^ _,-'_)   '-,_,-'
| |                           +--------+  | |  +----------+ | |   +----------+ | |     (_`|._`|'_,|'_)
| +---------------------------------------+ |    +----------+ |     +----------+ |     (_`|._`|'_,|'_)
| Application node                          |      +----------+       +----------+       `-._`|'_,-'
+-------------------------------------------+                                               `-'

```

# 安装说明

在kubernetes里部署一个对客户端透明的高可用Redis集群变成了一个很简单的工作, 只需要执行`make`命令就可以了：

```bash
cd redis-ha-operator
make
```

[Makefile] (<https://github.com/hackerthon2019/redis-ha/blob/master/redis-ha-operator/Makefile>)做以下工作:

1. 编译 [sentinel_tunnel](<https://github.com/RedisLabs/sentinel_tunnel>)的docker镜像
2. 部署redis-ha的Custom Resource Definitions进Kubernetes集群
3. 编译redis-ha-operator的docker镜像, 部署redis-ha-operator进Kubernetes集群
4. 部署示例的redis-ha Custom Resource进Kubernetes集群

我们还写了一个示例应用，该应用使用了上面部署出的Redis高可用集群, 运行该示例应用的办法可参考[这里] (<https://github.com/hackerthon2019/redis-ha/tree/master/redis-ha-demo>)

# 如何实现

1. 将[sentinel_tunnel](<https://github.com/RedisLabs/sentinel_tunnel>)封装为[docker镜像](<https://github.com/hackerthon2019/redis-ha/tree/master/redis-ha-operator/helm-charts/redis-ha-st/docker/redis-st>)，并[提供helm chart](<https://github.com/hackerthon2019/redis-ha/tree/master/redis-ha-operator/helm-charts/redis-ha-st/charts/redis-st>)以快速安装它。
2. 组织[redis-ha](<https://github.com/helm/charts/tree/master/stable/redis-ha>)及[上述helm chart](<https://github.com/hackerthon2019/redis-ha/tree/master/redis-ha-operator/helm-charts/redis-ha-st/charts/redis-st>)，最终形成一个[大chart](<https://github.com/hackerthon2019/redis-ha/tree/master/redis-ha-operator/helm-charts/redis-ha-st>)，用以快速在kubernetes中将上述架构部署出来。
3. 参考[operator-sdk的helm例子](<https://github.com/operator-framework/operator-sdk/blob/master/doc/helm/user-guide.md>)，将上述解决方案的chart封装成一个operator，用户只需要按照规范创建cr，即可在kubernetes集群中快速部署客户端透明的高可用redis集群。
4. 我们还写一个[简单的应用](<https://github.com/hackerthon2019/redis-ha/tree/master/redis-ha-demo>)，用以验证本架构解决的问题。

# 后续计划

1. 支持Redis分片集群

# 参考

1. <https://github.com/RedisLabs/sentinel_tunnel>
2. <https://helm.sh/>
3. <https://github.com/helm/charts/tree/master/stable/redis-ha>
4. <https://redis.io/topics/sentinel>
5. <https://github.com/operator-framework/operator-sdk/blob/master/doc/helm/user-guide.md>
6. <https://github.com/gin-gonic/gin>
7. <https://github.com/spf13/viper>
8. <https://vuejs.org/>
9. <https://element.eleme.io/>
10. <https://github.com/axios/axios>

