FROM node:alpine-12.7 AS front
LABEL Omekov Azamat <umekovazamat@gmail.com>
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM nginx:alpine
COPY --from=front /app/dist/frontend-angular /usr/share/nginx/html
EXPOSE 80
