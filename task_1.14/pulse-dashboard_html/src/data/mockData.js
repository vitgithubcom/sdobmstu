export const mockKPI = [
  { id: 1, title: "Выручка", value: "14 280", unit: "тыс. руб", plan: 15200, delta: -6.1, trend: "down", status: "critical" },
  { id: 2, title: "OEE (общая эффективность)", value: "73.5", unit: "%", plan: 78, delta: -4.5, trend: "down", status: "warning" },
  { id: 3, title: "Простои", value: "124", unit: "часов", plan: 80, delta: 55, trend: "down", status: "critical" },
  { id: 4, title: "Заказы в срок", value: "87.2", unit: "%", plan: 92, delta: -4.8, trend: "down", status: "warning" },
  { id: 5, title: "Склад (оборачив.)", value: "8.3", unit: "дней", plan: 7.2, delta: 15, trend: "down", status: "critical" }
]

export const mockChartData = [
  { name: "Пн", факт: 1820, план: 2100 },
  { name: "Вт", факт: 2050, план: 2100 },
  { name: "Ср", факт: 1980, план: 2100 },
  { name: "Чт", факт: 2140, план: 2100 },
  { name: "Пт", факт: 1780, план: 2100 },
  { name: "Сб", факт: 1620, план: 1800 },
  { name: "Вс", факт: 1450, план: 1600 }
]

export const mockAlerts = [
  { id: 1, system: "Mock-ERP", msg: "План продаж под угрозой (выполнение 68%)", severity: "critical", time: "12:34" },
  { id: 2, system: "Mock-MES", msg: "Станок #1042: превышение времени цикла", severity: "warning", time: "12:28" },
  { id: 3, system: "Mock-CRM", msg: "Не загружены сделки за последний час", severity: "critical", time: "12:15" },
  { id: 4, system: "Mock-Warehouse", msg: "Остатки по группе А ниже нормы", severity: "info", time: "11:58" }
]

export const mockIntegrations = [
  { name: "Mock-ERP", lastSync: "12:44:23", status: "ok", lag: "12 сек" },
  { name: "Mock-MES", lastSync: "12:44:01", status: "ok", lag: "34 сек" },
  { name: "Mock-CRM", lastSync: "12:30:05", status: "warning", lag: "14 мин" },
  { name: "Mock-Warehouse", lastSync: "12:42:10", status: "ok", lag: "2 мин" }
]