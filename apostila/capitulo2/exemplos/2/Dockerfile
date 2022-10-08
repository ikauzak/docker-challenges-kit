FROM ubuntu:20.04
WORKDIR /myapp
RUN for n in $(seq 1 5); do echo "Número $n" > test_$n.txt; done
CMD ["test_1.txt"]
ENTRYPOINT ["cat"]
