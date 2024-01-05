FROM harbor.internal.cn:8443/publicharbor_apulis/cambricon-pytorch16-python375-ubuntu1804-sdk1-120:v2
ADD install-client-2.8.5 /root/
RUN /root/install-client-2.8.5 -g -q && \
    rm -rf /root/install-client-2.8.5
ENV LD LIBRARY PATH="/usr/lib/orion/neuware-2.8.5:$LD LIBRARY PATH
