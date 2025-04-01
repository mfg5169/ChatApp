FROM golang:1.24.1-alpine3.20
#
# add bash to alpine Linux:
#
RUN apk update && apk upgrade
RUN apk add --no-cache bash
#
# turn off history file creation:
#
RUN echo "export HISTFILE=/dev/null" >> /etc/profile
#
# add a user (with no pwd) so we don't run as root:
#
RUN adduser -S user -G users -D
#

