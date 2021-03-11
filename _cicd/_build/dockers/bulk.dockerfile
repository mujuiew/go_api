FROM tnsmith/bulk-api:v1.0.6

ARG CUSTOMER_NAME=LISA

LABEL CUSTOMER_NAME=${CUSTOMER_NAME}
RUN echo "used config ${CUSTOMER_NAME}"

WORKDIR /application

COPY _cicd/_build/script/initialDB/${CUSTOMER_NAME} /application/config/initial
COPY _cicd/_build/script/InitialDB.sh /application/initialDB.sh

RUN chmod -R 755 /application

ENTRYPOINT /application/initialDB.sh