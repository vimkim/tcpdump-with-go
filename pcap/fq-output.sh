#!/bin/bash

# Check if fq-output.log exists
if [ -e "fq-output.log" ]; then
  echo "Warning: fq-output.log already exists. Exiting to prevent overwriting."
  exit 1
fi

echo '.packets[3]' >>fq-output.log
fq '.packets[3]' ./closed-from-server.pcap >>fq-output.log
echo "" >>fq-output.log

echo '.packets[3].packet' >>fq-output.log
fq '.packets[3].packet' ./closed-from-server.pcap >>fq-output.log
echo "" >>fq-output.log

echo '.packets[3].packet.payload' >>fq-output.log
fq '.packets[3].packet.payload' ./closed-from-server.pcap >>fq-output.log
echo "" >>fq-output.log

echo '.packets[3].packet.payload.payload' >>fq-output.log
fq '.packets[3].packet.payload.payload' ./closed-from-server.pcap >>fq-output.log
echo "" >>fq-output.log
