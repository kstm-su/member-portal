FROM node:20-alpine

WORKDIR /app

COPY package*.json pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile
EXPOSE 3000
COPY . .

CMD ["pnpm", "run", "dev"]
