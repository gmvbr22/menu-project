FROM node:current-alpine3.14
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
USER node
CMD [ "node", "app/main.js" ]