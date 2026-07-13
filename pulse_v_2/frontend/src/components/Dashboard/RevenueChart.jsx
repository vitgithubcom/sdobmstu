import { AreaChart, Area, XAxis, YAxis, Tooltip, ResponsiveContainer, CartesianGrid } from 'recharts'

function RevenueChart({ data }) {
  // Проверяем, что data есть и это массив
  const chartData = Array.isArray(data) && data.length > 0 ? data : []

  return (
    <div className="bg-white rounded-2xl p-5 shadow-sm border border-gray-100">
      <div className="flex justify-between items-center mb-4">
        <h3 className="font-semibold text-gray-800">📊 Выручка (факт vs план)</h3>
        <span className="text-xs text-gray-400 bg-gray-50 px-2 py-1 rounded-full">за неделю</span>
      </div>
      {chartData.length === 0 ? (
        <div className="h-64 flex items-center justify-center text-gray-400">
          <div className="text-center">
            <div className="text-4xl mb-2">📊</div>
            <p>Нет данных для отображения</p>
          </div>
        </div>
      ) : (
        <ResponsiveContainer width="100%" height={240}>
          <AreaChart data={chartData}>
            <defs>
              <linearGradient id="factGrad" x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor="#3b82f6" stopOpacity={0.3} />
                <stop offset="95%" stopColor="#3b82f6" stopOpacity={0} />
              </linearGradient>
            </defs>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="name" tick={{ fontSize: 12 }} />
            <YAxis tick={{ fontSize: 11 }} />
            <Tooltip />
            <Area 
              type="monotone" 
              dataKey="факт" 
              stroke="#3b82f6" 
              fill="url(#factGrad)" 
              strokeWidth={2} 
              name="Факт"
            />
            <Area 
              type="monotone" 
              dataKey="план" 
              stroke="#94a3b8" 
              strokeDasharray="5 5" 
              fill="none" 
              strokeWidth={2}
              name="План"
            />
          </AreaChart>
        </ResponsiveContainer>
      )}
    </div>
  )
}

export default RevenueChart