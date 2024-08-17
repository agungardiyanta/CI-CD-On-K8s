#Check My Site
FROM node:20.0-alpine as build
WORKDIR /app

COPY package*.json ./
RUN npm install
COPY . .

CMD [ "node", "app.js" ]