FROM ubuntu:20.04

ARG MIN=1
ARG MAX=5

ENV ARQUIVO_PADRAO=test_1.txt

WORKDIR /myapp
RUN for n in $(seq $MIN $MAX); do echo "Número $n" > test_$n.txt; done
CMD ["bash", "-c","cat $ARQUIVO_PADRAO"]
