#BASE IMAGE SETS UP GOLANG ENV AND REFLEX
FROM tnsmith/stubby4j:6.0.2

#SET WORKING DIRECTORY
WORKDIR /app

#COPY ALL STUB INTO WORKING DIRECTORY
COPY _tests/stubapi .

RUN cat schema/*.yaml >> cbs-ws.yaml

# CMD  java -jar /app/stubby4j-6.0.2.jar -da -ds -d cbs-ws.yaml -l 0.0.0.0 -s 8882
