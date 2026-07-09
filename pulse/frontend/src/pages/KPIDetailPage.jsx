import { useState, useEffect } from 'react'
import { useParams, Link } from 'react-router-dom'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'
import { BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer } from 'recharts'
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
    fetchKPIDetail()
  }, [id])

  const fetchKPIDetail = async () => {
    setLoading(true)
    try {
      const response = await api.get(`/kpi/${id}`)
      setKpi(response.data)
    } catch (error) {
      console.error('Error fetching KPI detail:', error)
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>
    )
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
          </div>
          <div className="text-right">
            <div className="text-3xl font-bold">{kpi.value}</div>
            <div className={`text-sm ${kpi.delta > 0 ? 'text-green-500' : 'text-red-500'}`}>
              {kpi.delta > 0 ? '▲' : '▼'} {Math.abs(kpi.delta)}% к плану
            </div>
          </div>
        </div>

        <div className="h-64">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart data={kpi.history}>
              <XAxis dataKey="period" />
              <YAxis />
              <Tooltip />
              <Bar dataKey="факт" fill="#3b82f6" />
              <Bar dataKey="план" fill="#94a3b8" />
            </BarChart>
          </ResponsiveContainer>
        </div>

        <div className="mt-6 grid grid-cols-2 md:grid-cols-4 gap-4">
          <div className="bg-gray-50 rounded-xl p-4 text-center">
            <div className="text-xs text-gray-500">Источник</div>
            <div className="font-medium">{kpi.source_system}</div>
          </div>
          <div className="bg-gray-50 rounded-xl p-4 text-center">
            <div className="text-xs text-gray-500">План</div>
            <div className="font-medium">{kpi.plan} {kpi.unit}</div>
          </div>
          <div className="bg-gray-50 rounded-xl p-4 text-center">
            <div className="text-xs text-gray-500">Факт</div>
            <div className="font-medium">{kpi.fact} {kpi.unit}</div>
          </div>
          <div className="bg-gray-50 rounded-xl p-4 text-center">
            <div className="text-xs text-gray-500">Выполнение</div>
            <div className={`font-medium ${kpi.completion >= 100 ? 'text-green-600' : 'text-red-600'}`}>
              {kpi.completion}%
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default KPIDetailPage