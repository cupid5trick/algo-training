#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>
#define LOCAL_IP_ADDR (0x7F000001) // IP 127.0.0.1
#define LOCAL_TCP_PORT (34567)     // 端口
int main(void)
{
    struct sockaddr_in local, peer;
    int ret;
    char buf[128];
    int sock = socket(AF_INET, SOCK_STREAM, 0);
    memset(&local, 0, sizeof(local));
    memset(&peer, 0, sizeof(peer));
    local.sin_family = AF_INET;
    local.sin_port = htons(LOCAL_TCP_PORT);
    local.sin_addr.s_addr = htonl(LOCAL_IP_ADDR);
    peer = local;
    int flag = 1;
    ret = setsockopt(sock, SOL_SOCKET, SO_REUSEADDR, &flag, sizeof(flag));
    if (ret == -1)
    {
        printf("Fail to setsocket SO_REUSEADDR: %s\n", strerror(errno));
        exit(1);
    }
    ret = bind(sock, (const struct sockaddr *)&local, sizeof(local));
    if (ret)
    {
        printf("Fail to bind: %s\n", strerror(errno));
        exit(1);
    }
    ret = connect(sock, (const struct sockaddr *)&peer, sizeof(peer));
    if (ret)
    {
        printf("Fail to connect myself: %s\n", strerror(errno));
        exit(1);
    }
    printf("Connect to myself successfully\n"); // 发送数据
    strcpy(buf, "Hello, myself~");
    send(sock, buf, strlen(buf), 0);
    memset(buf, 0, sizeof(buf)); // 接收数据
    recv(sock, buf, sizeof(buf), 0);
    printf("Recv the msg: %s\n", buf);
    sleep(1000);
    close(sock);
    return 0;
}

/*
这段代码的主要目的是通过创建一个TCP套接字并连接到自己（即本地回环地址`127.0.0.1`），来演示如何在没有调用`listen()`的情况下，建立自连接（loopback connection）。

在TCP协议中，通常服务器端需要调用`listen()`函数来监听连接请求，客户端通过`connect()`发起连接。然而，在这段代码中，程序并没有真正涉及传统的服务器-客户端通信模式，而是直接连接到自己。这种情况下，不需要调用`listen()`，因为没有外部的客户端尝试连接到该套接字。以下是代码的主要逻辑解释：

### 1. **创建TCP套接字**

```c
int sock = socket(AF_INET, SOCK_STREAM, 0);
```

这行代码创建了一个TCP套接字，使用IPv4地址族（`AF_INET`）和流式传输模式（`SOCK_STREAM`），即TCP协议。

### 2. **设置本地地址结构**

```c
memset(&local, 0, sizeof(local));
local.sin_family = AF_INET;
local.sin_port = htons(LOCAL_TCP_PORT);
local.sin_addr.s_addr = htonl(LOCAL_IP_ADDR);
```

`local`变量用于存储本地IP地址和端口号。这里的地址是回环地址`127.0.0.1`（通过宏`LOCAL_IP_ADDR`定义），端口号是`34567`。

### 3. **允许地址复用**

```c
int flag = 1;
ret = setsockopt(sock, SOL_SOCKET, SO_REUSEADDR, &flag, sizeof(flag));
```

这行代码设置了`SO_REUSEADDR`选项，允许多个套接字在`bind()`时复用相同的本地地址。这在程序多次运行时避免地址已被占用的错误。

### 4. **绑定本地地址**

```c
ret = bind(sock, (const struct sockaddr *)&local, sizeof(local));
```

`bind()`函数将套接字绑定到本地地址（IP 和端口）。这是必要的，因为之后需要连接到这个套接字。

### 5. **自连接**

```c
ret = connect(sock, (const struct sockaddr *)&peer, sizeof(peer));
```

这里调用`connect()`函数，实际上是让该套接字连接到自己（本地地址）。因为`peer`和`local`都指向同一个地址，即`127.0.0.1:34567`，所以这是一次自连接。

### 6. **发送和接收数据**

```c
strcpy(buf, "Hello, myself~");
send(sock, buf, strlen(buf), 0);
recv(sock, buf, sizeof(buf), 0);
printf("Recv the msg: %s\n", buf);
```

在自连接建立后，程序通过`send()`函数发送数据，然后通过`recv()`函数接收数据。这种自连接场景中，发送和接收数据可以在同一个套接字上进行。

### 7. **为什么不需要`listen()`？**

在传统的TCP服务器中，`listen()`用于将套接字转换为一个监听套接字，等待外部客户端的连接请求。然而，这段代码中通过`connect()`直接向自己发起连接，因此没有外部的客户端存在，程序本身既是客户端又是服务器。没有必要调用`listen()`，因为并没有其他的客户端连接到这个套接字。

这种自连接的行为在某些特定场景下是有用的，比如调试或特殊的网络协议测试。

### 总结

- 代码实现了一个TCP自连接，即自己连接自己。
- 不需要`listen()`的原因是程序通过`connect()`直接发起自连接，传统的服务器-客户端模型在这里不适用。
- 程序成功发送并接收了数据，这种自连接可以用来测试网络通信的基本功能。
*/