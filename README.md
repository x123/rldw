rldw
============================
regex long, didn't write

rldw commands:
- `parse` - parse IPv4 and IPv6 addresses from stdin
- `generate` - generate random publicly routeable IPv4 and IPv6
  addresses

## Parse IPv4 and IPv6 addresses from stdin
`rldw parse` is used to quickly strip IPv4 and IPv6 addresses from stdin. This
is handy for rapidly pulling data out of log files for further processing,
without having to paste in complicated regular expressions.

:exclamation: `rldw parse` defaults with `--ipv4` to `true`, in order to parse
for ipv6 addresses you must use either `rldw parse -4=false -6` or `rldw
parse--ipv4=false --ipv6`

### Parse nft sets for IPv4 addresses and print stats
```shell
$ nft list set inet filter scanners | rldw parse -4 -s
45.33.80.242
45.76.146.20
45.148.10.81
46.174.8.146
59.126.157.116
61.177.172.104
62.197.136.119
62.210.13.20
66.228.33.237
71.6.232.8
# Statistics
# Lines parsed: 5
# IPv4 Addresses found/Lines parsed: 10/5
# Percentage of lines with IPv4 addresses: 200.00%
# Average number of IPv4 addresses found/Line: 2.00
```

### Parse journalctl logs for IPv4 addresses and print stats 
```shell
$ journalctl -n 5 --no-pager | rldw parse -4 -s
54.197.46.151
116.112.99.254
189.117.220.234
160.72.224.244
129.91.203.179
8.188.201.141
88.171.58.235
6.76.86.38
# Statistics
# Lines parsed: 5
# IPv4 Addresses found/Lines parsed: 8/5
# Percentage of lines with IPv4 addresses: 160.00%
# Average number of IPv4 addresses found/Line: 1.60
```

### Parse ip addr output for IPv6 addresses
| :zap: IPv6 address support is not perfect, use at your own risk |
|-----------------------------------------------------------------|
```shell
$ ip addr | ./rldw parse -4=false -6
::1
fe80::
```

### Parse logs for IPv6 addresses
| :zap: IPv6 address support is not perfect, use at your own risk |
|-----------------------------------------------------------------|
```shell
$ journalctl -n 100 --no-pager | rldw parse -6 -4=false
::
::1
ff01::
ff02::
ff02::
1::
0:0:0:1::
e900:0:0:1::
6dfb:79ba:36a3:9260:abad:f319:7bac:33b9
afa7:d32d:4a45:4017:67ff:8881:4041:94c1
```

## Generate random IPv4 and IPv6 addresses
`rldw generate` can be used to generate random publicly routeable IPv4 and IPv6
addresses. This can be handy for generating unit test datasets.

### IPv4: Generate 5 random IPv4 addresses
```shell
$ rldw generate -4 -c 5
62.150.113.112
176.59.94.136
105.238.3.48
98.55.211.53
144.35.170.232
```

### IPv6: Generate 5 random IPv6 addresses
```shell
$ rldw generate -4=false -6 -c 5
8b09:b7be:ec78:2195:6912:59a3:4eea:35a8
a344:670:7ecc:e744:c315:c83e:9b32:f73b
4a34:5330:e3a1:15c2:6924:5ce1:b468:5c36
227:ffef:c94b:c7f9:78b2:beef:c3b5:ecd3
5c04:5b27:add8:379e:82d2:9f18:6ae3:5a0d
```
