# ===== СТРОИМ =====
FROM node:18-alpine AS builder

WORKDIR /app

# Копируем всё (включая исходники)
COPY . .

# Устанавливаем зависимости и собираем
RUN npm install && \
    npm run build

# ===== ОТДАЁМ =====
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]