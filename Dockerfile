FROM centos
COPY . .
EXPOSE 8080
CMD ["./rest-api"]
