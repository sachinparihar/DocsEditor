FROM nginx:alpine
LABEL MAINTAINER="Sachin Parihar <sachinparihar98120@gmail.com>"
COPY *.html /usr/share/nginx/html/
COPY images /usr/share/nginx/html/images
WORKDIR /usr/share/nginx/html
EXPOSE 8000  
CMD ["nginx", "-g", "daemon off;"]
