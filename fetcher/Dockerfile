FROM node:alpine

WORKDIR /app/fetcher

COPY package*.json ./

COPY yarn.lock ./

RUN yarn install

COPY . .

EXPOSE 8001

CMD [ "node", "main.js" ]