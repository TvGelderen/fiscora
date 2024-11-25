FROM oven/bun AS builder
WORKDIR /app
COPY . .
RUN bun i
RUN bun run build

FROM oven/bun
COPY --from=builder /app/build .

EXPOSE 4173

ENV NODE_ENV=production
ENV PORT=4173

CMD ["bun", "run", "start"]
