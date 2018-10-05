FROM alpine

COPY log_cleaner /usr/cleaner/

WORKDIR /usr/cleaner

CMD ./log_cleaner