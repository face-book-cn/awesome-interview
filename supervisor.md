## supervisor

- 设置开机启动

```bash
wget -O /usr/lib/systemd/system/supervisord.service https://raw.githubusercontent.com/Supervisor/initscripts/master/centos-systemd-etcs
```

### 重新加载Systemd配置，使得Supervisord配置生效：

```shell
systemctl daemon-reload
```

### 设置开机启动

```shell
systemctl enable supervisord.service
```

### 检测是否开机启动

```shell
systemctl is-enabled supervisord.service
```

