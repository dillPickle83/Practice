#!/bin/python3

import sys # allows us to enter cmd line args and other thing
import socket
from datetime import datetime

# python scanner.py <ip> | in how the script is run, scanner.py is arg 1 and the ip is the 2nd argument
if len(sys.argv) == 2:
    target = socket.gethostbyname(sys.argv[1]) # If a hostname is given then the gethostbyname function translates it to IPv4
else:
    print("Invalid number of arguments")
    print("Syntax: python3 scanner.py <ip>")
    sys.exit() # closes the program, else the program might be stuck

print("-" * 20)
print("Scanning target: " + target)
print("Time started: " + str(datetime.now()))
print("-" * 20)

try:
    for port in range(50, 85):
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM) # The AF_INET is the IPv4 and the SOCK_STREAM is the port
        socket.setdefaulttimeout(1) # Arg is a float
        result = s.connect_ex((target, port)) # Returns 0 is the connection happens and an error indicator if an error happens
        print(f"Checking port {port}")
        if result == 0:
            print(f"Port {port} is open")

except KeyboardInterrupt:
    print("Exiting program")
    sys.exit()

except socket.gaierror:
    print("Hostname couldn't be resolved")
    sys.exit()

except socket.error:
    print("Couldn't connect to server")
    sys.exit()
