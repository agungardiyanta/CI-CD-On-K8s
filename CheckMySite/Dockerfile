#NodeJs Check Site
FROM node:20.0-alpine as build
WORKDIR /app

COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD [ "node", "app.js" ]