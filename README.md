# Lab of Tcpdump & Termshark in Go

![image](https://github.com/user-attachments/assets/5747ea3d-b17a-4f2f-8ed9-259fc897ea2f)

### How to replay

```sh
# obverse & dump
sudo tcpdump -i any tcp port 8080 -vv -w <pcap_file>

# read & analyze
termshark -r client-closes-connection.pcap

# optional
./fq-output.sh # internally runs fq
```

### Why [FIN, ACK] - [FIN, ACK] - [ACK] ?

It is also possible to terminate the connection by a 3-way handshake, when host A sends a FIN and host B replies with a FIN & ACK (combining two steps into one) and host A replies with an ACK.

 Tanenbaum, Andrew S. (2003-03-17). Computer Networks (Fourth ed.). Prentice Hall. ISBN 978-0-13-066102-9.

