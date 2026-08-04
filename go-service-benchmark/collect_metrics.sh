#!/bin/bash

echo "=== Сбор метрик для диагностики ==="

# 1. Статистика Nginx
echo -e "\n--- Nginx Status ---"
curl -s http://localhost:84/nginx-status

# 2. Активные запросы в PostgreSQL
echo -e "\n--- Активные запросы PostgreSQL ---"
docker exec benchmark-postgres psql -U postgres -d testdb -c "
    SELECT pid, state, query, now() - query_start as duration 
    FROM pg_stat_activity 
    WHERE state = 'active' AND query NOT LIKE '%pg_stat_activity%'
    ORDER BY duration DESC
    LIMIT 10;
"

# 3. Медленные запросы
echo -e "\n--- Медленные запросы (pg_stat_statements) ---"
docker exec benchmark-postgres psql -U postgres -d testdb -c "
    SELECT query, calls, total_time, mean_time 
    FROM pg_stat_statements 
    ORDER BY mean_time DESC 
    LIMIT 10;
"

# 4. Go-метрики
echo -e "\n--- Go метрики пула соединений ---"
curl -s http://localhost:84/debug/metrics | jq .

# 5. Потребление ресурсов
echo -e "\n--- Потребление ресурсов контейнерами ---"
docker stats --no-stream --format "table {{.Container}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.NetIO}}"

# 6. Логи Go-сервиса (последние ошибки)
echo -e "\n--- Последние ошибки Go-сервиса ---"
docker logs --tail 20 benchmark-go-app 2>&1 | grep -i error || echo "Ошибок нет"

# 7. Логи PostgreSQL (последние записи)
echo -e "\n--- Последние логи PostgreSQL ---"
docker logs --tail 10 benchmark-postgres 2>&1 | grep -E "(LOG|ERROR|duration)" || echo "Логов нет"