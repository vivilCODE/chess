FROM node:18-alpine

WORKDIR /app

COPY ./public/ ./public/
COPY ./frontend/src/ ./src/
COPY package.json ./
COPY package-lock.json ./

RUN npm install

CMD ["npm", "start"]