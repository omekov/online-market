FROM node:alpine-12.7 AS front
WORKDIR /app
COPY . .
RUN npm ci && npm run build

FROM nginx:alpine
COPY --from=front /app/dist/frontend-angular /usr/share/nginx/html
EXPOSE 80
