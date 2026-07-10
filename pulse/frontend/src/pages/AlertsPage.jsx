import { useState, useEffect } from 'react'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'
import LoadingSpinner from '../components/Common/LoadingSpinner'
import api from '../api/client'

function AlertsPage() {
  const { user, logout } = useAuth()
  const [currentTime, setCurrentTime] = useState(
    new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
  )
  const [alerts, setAlerts] = useState([])
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
    fetchAlerts()
  }, [])

  const fetchAlerts = async () => {
    setLoading(true)
    try {
      const response = await api.get('/alerts')
      setAlerts(Array.isArray(response.data) ? response.data : [])
    } catch (error) {
      console.error('Error fetching alerts:', error)
      setAlerts([])
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return <LoadingSpinner fullScreen />
  }

  return (
    <div className="max-w-7xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header currentTime={currentTime} user={user} onLogout={logout} title="Все тревоги" />

      <div className="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <div className="p-4 border-b flex justify-between items-center">
          <div className="text-sm text-gray-500">
            Всего: {alerts.length} тревог
          </div>
          <button
            onClick={fetchAlerts}
            className="text-xs text-blue-600 hover:text-blue-800"
          >
            ⟳ Обновить
          </button>
        </div>

        {alerts.length === 0 ? (
          <div className="p-8 text-center text-gray-400">
            <div className="text-4xl mb-2">✅</div>
            <p>Нет активных тревог</p>
            <p className="text-sm mt-1">Все системы работают штатно</p>
          </div>
        ) : (
          <div className="divide-y divide-gray-100">
            {alerts.map((alert) => (
              <div key={alert.id} className="p-4 flex items-start gap-3 hover:bg-gray-50">
                <div className={`w-2 h-2 mt-2 rounded-full ${
                  alert.severity === 'critical' ? 'bg-red-500' :
                  alert.severity === 'warning' ? 'bg-amber-500' : 'bg-blue-400'
                }`} />
                <div className="flex-1">
                  <div className="flex justify-between items-start">
                    <div>
                      <span className="font-mono text-xs text-gray-400">{alert.system}</span>
                      <div className="text-gray-700">{alert.message}</div>
                    </div>
                    <span className={`text-xs px-2 py-1 rounded-full ${
                      alert.severity === 'critical' ? 'bg-red-100 text-red-700' :
                      alert.severity === 'warning' ? 'bg-amber-100 text-amber-700' : 'bg-blue-100 text-blue-700'
                    }`}>
                      {alert.severity === 'critical' ? 'Критично' :
                       alert.severity === 'warning' ? 'Внимание' : 'Инфо'}
                    </span>
                  </div>
                  <div className="text-xs text-gray-400 mt-1">
                    {alert.created_at ? new Date(alert.created_at).toLocaleString() : ''}
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  )
}

export default AlertsPage