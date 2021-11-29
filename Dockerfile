 FROM httpd
 COPY ./index.html /usr/local/apache2/htdocs/
 WORKDIR /usr/local/apache2/htdocs/
 RUN chown -R daemon:daemon * && chmod -R 755 *
 EXPOSE 80
