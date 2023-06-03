#!/usr/bin/python3

import socket

TCP_IP = "192.168.0.1"
TCP_PORT = 6996
BUFFER = 100

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(("", 6996))
#s.bind((TCP_IP, TCP_PORT))
s.listen(1)

conn, addr = s.accept()
print("Connection Address:", addr)
while 1:
    data = conn.recv(BUFFER)
    if not data: break
    print("Received data", data)
    conn.send(data) #echo

conn.close
