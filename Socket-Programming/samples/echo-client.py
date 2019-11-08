#!/usr/bin/env python3

import socket

HOST = '127.0.0.1'  # The server's hostname or IP address
PORT = 1373        # The port used by the server

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    s.sendall('سلام دنیا'.encode())
    data = s.recv(1024)
    message = data.decode()
    print(f'received {message!r}')
