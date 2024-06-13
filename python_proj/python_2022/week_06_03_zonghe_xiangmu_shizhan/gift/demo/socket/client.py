import socket

host = '127.0.0.1'
port = 29999
s = socket.socket()
s.connect((host, port))
print(s.recv(10240))
