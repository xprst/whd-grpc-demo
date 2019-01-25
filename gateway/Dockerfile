FROM envoyproxy/envoy:latest

RUN mkdir /var/log/envoy/
COPY ./script/gateway_start /etc/gateway_start
RUN chmod +x /etc/gateway_start
COPY ./proto.pb /etc/gateway_proto.pb
CMD /etc/gateway_start

