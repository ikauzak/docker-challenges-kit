FROM ubuntu:20.04

ENV MAX_COUNT=10
ENV SLEEP_TIME=5
ENV COUNTER=0

# Necessário para ajustar data e hora no container
ENV TZ="America/Sao_Paulo"
RUN apt-get update -y && apt-get install tzdata

RUN echo 'while [  $COUNTER -lt $MAX_COUNT ]; do  ts=$(date +"%d/%m/%Y-%X") ; echo "${ts} ${HOSTNAME}" | tee -a arquivo_persistente.txt && sleep $SLEEP_TIME ; let COUNTER=COUNTER+1 ;  done' > /init.sh && chmod +x /init.sh

WORKDIR /outputs
CMD ["bash", "-c", "/init.sh"]
