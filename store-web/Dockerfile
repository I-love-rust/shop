### STAGE 1: Build ###
FROM node:18.16.1-alpine as builder
WORKDIR /app
COPY ./store-web/package.json ./
COPY ./store-web/package-lock.json ./
RUN npm install
COPY ./store-web ./
RUN npm run build


### STAGE 2: Run ###
FROM nginx:1.25.1-alpine
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build /var/www/default
