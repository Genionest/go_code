TCP端口扫描

TCP握手

Opened Port
client -(1.syn)-> server
client <-(2.syn-ack)- server
client -(3.ack)-> server

Closed Port
client -(1.syn)-> server
client <-(4.rst)- server

Filtered Port
client -(1.syn)-> firewall server
        timeout
