FROM centos
COPY . .
EXPOSE 8080
CMD ["./server"]
