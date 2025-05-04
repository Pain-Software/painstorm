# Build
FROM node:20-alpine AS builder

WORKDIR /app

ARG PUBLIC_BACKEND_URL
ENV PUBLIC_BACKEND_URL=$PUBLIC_BACKEND_URL

COPY package*.json ./
RUN npm install --force

COPY . .

RUN echo "PUBLIC_BACKEND_URL=$PUBLIC_BACKEND_URL" > .env

RUN npm run build

RUN npm prune --production --force

# Runner
FROM node:20-alpine AS runner

WORKDIR /app

COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/node_modules ./node_modules

ENV NODE_ENV=production
ENV PORT=8081

EXPOSE 8081

CMD ["node", "build"]
