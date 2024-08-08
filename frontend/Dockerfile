FROM node:18-alpine AS build
WORKDIR /app
COPY package*.json .
RUN npm i --legacy-peer-deps
COPY . .
RUN npm run build
RUN npm prune --production --legacy-peer-deps

ENV NODE_ENV=production
ENV PORT=4173

EXPOSE 4173

CMD [ "node", "build" ]
