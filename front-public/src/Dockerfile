FROM node:alpine-12.7 AS front
LABEL Omekov Azamat <umekovazamat@gmail.com>
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM nginx:alpine
COPY --from=front /app/dist/front-public /usr/share/nginx/html
EXPOSE 80
