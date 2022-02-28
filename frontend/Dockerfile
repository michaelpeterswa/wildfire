# -=-=-=-=-=-=- Compile Image -=-=-=-=-=-=-

FROM node AS builder

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci

COPY . .

RUN npm run build

# -=-=-=-=-=-=- Final Image -=-=-=-=-=-=-

FROM httpd:alpine
COPY ./dist/ /usr/local/apache2/htdocs/