#!/usr/bin/python3

import socket
s = socket.socket()
s.connect(("192.168.0.1", 22))
answer = s.recv(1024)
print(answer)
s.close()
