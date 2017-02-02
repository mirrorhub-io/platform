FROM ubuntu

COPY bin/mirrorhub /bin/

EXPOSE 9000 8080
CMD /bin/mirrorhub api
