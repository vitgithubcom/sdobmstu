import { useState, useEffect } from 'react'
import { useParams, Link } from 'react-router-dom'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'
import { BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer, CartesianGrid } from 'recharts'
import LoadingSpinner from '../components/Common/LoadingSpinner'
import api from '../api/client'

function KPIDetailPage() {
  const { id } = useParams()
  const { user, logout } = useAuth()
  const [currentTime, setCurrentTime] = useState(
    new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
  )
  const [kpi, setKpi] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentTime(
        new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
      )
    }, 60000)
    return () => clearInterval(timer)
  }, [])

  useEffect(() => {
    fetchKPIDetail()
  }, [id])

  const fetchKPIDetail = async () => {
    setLoading(true)
    try {
      const response = await api.get(`/kpi/${id}`)
      const data = response.data
      
      // Преобразуем историю для графика
      if (data.history && data.history.length > 0) {
        data.history = data.history.map(item => ({
          period: item.period,
          факт: item.факт || 0,
          план: item.план || 0,
        }))
      } else {
        // Если истории нет — создаём моковые данные для демонстрации
        data.history = [
          { period: 'Пн', факт: 1820, план: 2100 },
          { period: 'Вт', факт: 2050, план: 2100 },
          { period: 'Ср', факт: 1980, план: 2100 },
          { period: 'Чт', факт: 2140, план: 2100 },
          { period: 'Пт', факт: 1780, план: 2100 },
          { period: 'Сб', факт: 1620, план: 1800 },
          { period: 'Вс', факт: 1450, план: 1600 },
        ]
      }
      
      setKpi(data)
    } catch (error) {
      console.error('Error fetching KPI detail:', error)
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return <LoadingSpinner fullScreen />
  }

  if (!kpi) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <p className="text-gray-500">Показатель не найден</p>
          <Link to="/dashboard" className="text-blue-600 hover:underline">← Вернуться</Link>
        </div>
      </div>
    )
  }

  return (
    <div className="max-w-5xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header currentTime={currentTime} user={user} onLogout={logout} title={`Детализация: ${kpi.name}`} />

      <Link to="/dashboard" className="text-blue-600 hover:underline text-sm mb-4 inline-block">
        ← Назад к дашборду
      </Link>

      <div className="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
        <div className="flex justify-between items-start mb-6">
          <div>
            <div className="text-sm text-gray-500">{kpi.code}</div>
            <h2 className="text-2xl font-bold">{kpi.name}</h2>
            <div className="text-gray-500 text-sm mt-1">Единица: {kpi.unit}</div>
            <div className="text-xs text-gray-400 mt-1">Источник: {kpi.source_system || 'Не указан'}</div>
          </div>
          <div className="text-right">
            <div className="text-3xl font-bold">{kpi.value}</div>
            <div className={`text-sm ${kpi.delta > 0 ? 'text-green-500' : 'text-red-500'}`}>
              {kpi.delta > 0 ? '▲' : '▼'} {Math.abs(kpi.delta).toFixed(1)}% к плану
            </div>
            <div className="text-xs text-gray-400 mt-1">
              Выполнение: {kpi.completion ? kpi.completion.toFixed(1) : '0'}%
            </div>
          </div>
        </div>

        <div className="h-64">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart data={kpi.history || []}>
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="period" />
              <YAxis />
              <Tooltip />
              <Bar dataKey="факт" fill="#3b82f6" name="Факт" />
              <Bar dataKey="план" fill="#94a3b8" name="План" />
            </BarChart>
          </ResponsiveContainer>
        </div>

        <div className="mt-6 grid grid-cols-2 md:grid-cols-4 gap-4">
          <div className="bg-gray-50 rounded-xl p-4 text-center">
            <div className="text-xs text-gray-500">Источник</div>
            <div className="font-medium">{kpi.source_system || '—'}</div>
          </div>
          <div className="bg-gray-50 rounded-xl p-4 text-center">
            <div className="text-xs text-gray-500">План</div>
            <div className="font-medium">{kpi.plan} {kpi.unit}</div>
          </div>
          <div className="bg-gray-50 rounded-xl p-4 text-center">
            <div className="text-xs text-gray-500">Факт</div>
            <div className="font-medium">{kpi.value} {kpi.unit}</div>
          </div>
          <div className="bg-gray-50 rounded-xl p-4 text-center">
            <div className="text-xs text-gray-500">Выполнение</div>
            <div className={`font-medium ${kpi.completion >= 100 ? 'text-green-600' : 'text-red-600'}`}>
              {kpi.completion ? kpi.completion.toFixed(1) : '0'}%
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default KPIDetailPage