FROM nginx:alpine
LABEL MAINTAINER="Sachin Parihar <sachinparihar98120@gmail.com>"
COPY *.html /usr/share/nginx/html/
WORKDIR /usr/share/nginx/html
EXPOSE 80   
CMD ["nginx", "-g", "daemon off;"]
