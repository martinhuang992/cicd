FROM centos:7
COPY hello /root/server
COPY index.html /root/index.html
COPY index.jpg /root/index.jpg
RUN chmod +x /root/server
EXPOSE 8080
CMD /root/server
