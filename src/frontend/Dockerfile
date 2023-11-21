# Use Nginx alpine as the base image
FROM nginx:alpine

# Label the maintainer
LABEL MAINTAINER="Sachin Parihar <sachinparihar98120@gmail.com>"

# Copy HTML and CSS files to the Nginx html directory
COPY *.html /usr/share/nginx/html/
COPY images /usr/share/nginx/html/images

# Set the working directory
WORKDIR /usr/share/nginx/html

# Expose port 8000
EXPOSE 8000

# Copy a custom Nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]