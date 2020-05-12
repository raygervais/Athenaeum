FROM node:alpine as base

# Move to working directory /app
WORKDIR /app

# Dependency Container
FROM base AS dependencies

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./

# Install Dependencies
RUN npm i --silent
RUN npm install react-scripts@3.4.1 -g --silent

# Builder Container
FROM base AS builder

COPY . ./

COPY --from=dependencies /app/node_modules /app/node_modules

RUN npm run build

# Final Production Container
FROM nginx:alpine

## Copy build to nginx
COPY --from=builder /app/build /usr/share/nginx/html

# Handle React Router
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

