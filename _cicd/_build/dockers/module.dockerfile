FROM tnsmith/yottadb-golang:r1.30.0-1.15.6 as builder

#CMD directory where application is launched from
ARG BUILD_DATE
ARG VCS_REF

#SET WORKING DIRECTORY
WORKDIR /src

#COPY CODE INTO WORKSPACE
COPY . .

#BUILD APP
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -o /go/servapi ./cmd/servapi
# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -o /go/consumer ./cmd/consumer
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -o /go/job ./cmd/job
# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -o /go/event ./cmd/event

# REDUCE SIZE
RUN chmod +x -R /src/_cicd/_build/script
RUN chmod +x -R /src/_cicd/_build/yotta
RUN chmod +x -R /src/_cicd/_build/job
RUN chmod +x -R /go

# GET APP STREAMER
FROM tnsmith/yotta-streamer:v1.17-slim as streamer

#MINIMAL CONTAINER COPIES WHAT IS NEEDED
FROM tnsmith/yottadb:r1.30.0

ARG BUILD_DATE
ARG VCS_REF

WORKDIR /

# API
COPY --chown=ydbadm:ydbadm --from=builder /go/servapi /application/servapi
# COPY --chown=ydbadm:ydbadm --from=builder /go/consumer /application/consumer
COPY --chown=ydbadm:ydbadm --from=builder /go/job /application/job
# COPY --chown=ydbadm:ydbadm --from=builder /go/event /application/event
# STREAMER
# COPY --chown=ydbadm:ydbadm --from=streamer /application/app /application/streamer

# TRIGGER
COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/script/startService.sh /application/startService.sh
# COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/script/startTrigger.sh /application/startTrigger.sh
# COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/yotta/trigger/rtns/*.m /application/rtns/
# COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/yotta/trigger/rtns/indxRegen/*.m /application/rtns/
# COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/yotta/trigger/Z8804trigger.trg /application/trigger/Z8804trigger.trg
# COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/script/replaceIndex.sh /application/replaceIndex.sh

# JOB
COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/job/keepPodRunning.sh /application/keepPodRunning.sh
COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/job/jobExecute.sh /application/jobExecute.sh
COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/job/jobInfo.json /application/jobcfg/jobInfo.json

# DEPLOYMENT
COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/deployment/app /application/deployment
# COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/deployment/db /application/deployment

# STREAMER
# COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/yotta/ddl /application/ddl/
# COPY --chown=ydbadm:ydbadm --from=builder /src/_cicd/_build/yotta/global.json /application/global/

USER ydbadm

# SFTP
RUN mkdir -p /application/file/external/input
RUN mkdir -p /application/sftp
RUN mkdir -p /application/file/external/output

ENV APP_DB_DB_ENGINE YOTTA
ENV DOMAIN_NUMBER C7
ENV SERVICE_NUMBER 02
ENV DOMAIN_CONTEXT "/receipt"

ENV PATH "$PATH:/application"

# RUN chmod +x -R /application
RUN ls -lrt /application
# COPY --chown=ydbadm:ydbadm /application /application

ENTRYPOINT ["/application/startService.sh"]

# Following packages is added for help debugging
LABEL org.opencontainers.image.authors="jiraphon.sa <jiraphon.sa@tnis.com>" \
    org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.revision="${VCS_REF}" \
    org.opencontainers.image.vendor="T.N. Incorporation Ltd."