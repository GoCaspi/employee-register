FROM swaggerapi/swagger-ui

RUN mkdir /swag

COPY .collection.yml /swag

ENV SWAGGER_JSON=/swag/collection.yml

EXPOSE 8080