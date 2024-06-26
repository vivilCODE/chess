FROM envoyproxy/envoy:v1.29-latest

COPY envoy.yaml /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml

EXPOSE 8080
CMD ["envoy", "-c", "/etc/envoy/envoy.yaml"]