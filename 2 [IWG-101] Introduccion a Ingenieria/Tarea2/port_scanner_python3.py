import socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
target = input('URL a escanear: ')
def pscan(port):
    try:
        s.connect((target,port))
        return True
    except:
        return False
for x in range(1,100):
    if pscan(x):
        print('Puerto',x,'abierto')
    else:
    	print('Puerto',x,'cerrado')
