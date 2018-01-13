# GFW
Golang被墙只能选择翻墙，写得很简单，但是自己用来翻墙的话足够了

# 配置文件说明
本程序是使用Config.json文件做为配置文件，里面各项配置你可以自行定义，若是没有提前预备配置文件，那么会生成一个服务端的配置文件，服务端监听端口为59386，本地端口为12345，连接密码随机生成
<pre>
{
	"Local": "127.0.0.1:12345", 
	"Server": ":59386",
	"Current": "Server",
	"Password": "zLgOJxm372bcXGIDv/26gwCXEWeFxMhA+obUV47pTolYajuLyu7bqUKm0cNV3aAe4L5KF0zPnI9DMTlBXyB9HGCIsX9QEHIU9z/k+K/HtQQFKHXhwcXQK3jfsMb57aMJaxhLAtYIqvEtYQ25ZS7L/1MsmfMvXgsdE+J3FQdw5q7AaQ/ypJuArHb29d7TZJq2PtlG6LTjKiLskHmnyR8a+6gkWwyrJetzjZaKPW50jAHwPClWlCaf9GxRcYcyfHrSk5Jjka0b/F00SLI6ezhSR0lNWf69nlRaM+Wizp1+s6E2I22EgRZPuzeV56XNgjBomOrVbxLawgohRbzX2AY1RA=="
}

Local：配置监听端口，仅作为客户端时有用
Server：配置监听端口，仅作为服务端时有用
Current：默认为服务端Server，若是要使用客户端，则使用Local作为参数
Password：连接密码，默认会生成一个随机密码，你只需要确保客户端和服务端密码一致即可
</pre>

# 使用说明
简单说下使用方式，使用该项目生成的二进制文件后，你可以自行配置Config.json文件，程序会根据你配置的Config.json文件决定是作为服务端还是客户端

假设你是要将该程序作为服务端，那么你只需按照下面的配置Config.json文件即可实现一个翻墙服务端

<pre>
{
	"Local": ":12345",
	"Server": ":59386",
	"Current": "Server",
	"Password": "zLgOJxm372bcXGIDv/26gwCXEWeFxMhA+obUV47pTolYajuLyu7bqUKm0cNV3aAe4L5KF0zPnI9DMTlBXyB9HGCIsX9QEHIU9z/k+K/HtQQFKHXhwcXQK3jfsMb57aMJaxhLAtYIqvEtYQ25ZS7L/1MsmfMvXgsdE+J3FQdw5q7AaQ/ypJuArHb29d7TZJq2PtlG6LTjKiLskHmnyR8a+6gkWwyrJetzjZaKPW50jAHwPClWlCaf9GxRcYcyfHrSk5Jjka0b/F00SLI6ezhSR0lNWf69nlRaM+Wizp1+s6E2I22EgRZPuzeV56XNgjBomOrVbxLawgohRbzX2AY1RA=="
}
</pre>
按照上面配置好服务端配置文件后，只需要将你配置好的Config.json文件和生成的二进制文件一起上传到服务器并运行程序就可以了

如果要作为客户端使用，那么也只需要按下面的方式配置下Config.json文件即可。

假设你已经配置好服务端的服务器地址为8.8.8.8，设置的端口为59386，那么客户端的配置文件如下
<pre>
{
	"Local": "127.0.0.1:12345",
	"Server": "8.8.8.8:59386",
	"Current": "Local",
	"Password": "zLgOJxm372bcXGIDv/26gwCXEWeFxMhA+obUV47pTolYajuLyu7bqUKm0cNV3aAe4L5KF0zPnI9DMTlBXyB9HGCIsX9QEHIU9z/k+K/HtQQFKHXhwcXQK3jfsMb57aMJaxhLAtYIqvEtYQ25ZS7L/1MsmfMvXgsdE+J3FQdw5q7AaQ/ypJuArHb29d7TZJq2PtlG6LTjKiLskHmnyR8a+6gkWwyrJetzjZaKPW50jAHwPClWlCaf9GxRcYcyfHrSk5Jjka0b/F00SLI6ezhSR0lNWf69nlRaM+Wizp1+s6E2I22EgRZPuzeV56XNgjBomOrVbxLawgohRbzX2AY1RA=="
}
</pre>
配置好之后直接运行就可以了，而你本地的socks5代理端口为12345，该端口你可以自行修改使用。
