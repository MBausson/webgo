import socket as s
import requests as r

def request(content):
    client = s.socket(s.AF_INET, s.SOCK_STREAM)
    client.connect(("localhost", 8888))
    client.sendall(content.encode("utf-8 "))
    
    print(f"Sent {content}.")

def get(url):
    res = r.get(url=url)
    print(res)
    print(res.text)