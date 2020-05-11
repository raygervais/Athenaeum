FROM node:alpine AS builder

# build environment
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./
COPY package-lock.json ./

RUN npm i --silent
RUN npm install react-scripts@3.4.1 -g --silent

COPY . ./

RUN npm run build

# production environment
FROM nginx:alpine

COPY --from=builder /app/build /usr/share/nginx/html

# handle React Router
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

