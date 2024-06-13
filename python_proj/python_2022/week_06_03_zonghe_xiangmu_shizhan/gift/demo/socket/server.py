# coding:utf-8
import socket

s = socket.socket()
host = '127.0.0.1'
port = 29999
s.bind((host, port))

s.listen(10)
while True:
    c, addr = s.accept()
    print('get connection from', addr)
    c.send(b'hello world')
    c.close()
