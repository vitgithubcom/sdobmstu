import { useState, useEffect } from 'react'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'
import api from '../api/client'

function AdminIntegrationsPage() {
  const { user, logout } = useAuth()
  const [currentTime, setCurrentTime] = useState(
    new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
  )
  const [integrations, setIntegrations] = useState([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchIntegrations()
  }, [])

  const fetchIntegrations = async () => {
    setLoading(true)
    try {
      const response = await api.get('/integrations')
      setIntegrations(response.data)
    } catch (error) {
      console.error('Error fetching integrations:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleSync = async (id) => {
    try {
      await api.post(`/integrations/${id}/sync`)
      fetchIntegrations()
    } catch (error) {
      alert('Ошибка синхронизации')
    }
  }

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>
    )
  }

  return (
    <div className="max-w-7xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header currentTime={currentTime} user={user} onLogout={logout} title="Управление интеграциями" />

      <div className="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <div className="p-4 border-b flex justify-between items-center">
          <div className="text-sm text-gray-500">Всего: {integrations.length} источников</div>
          <button className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-xl text-sm transition">
            ⟳ Синхронизировать всё
          </button>
        </div>

        <div className="divide-y divide-gray-100">
          {integrations.map((integ) => (
            <div key={integ.id} className="p-4 flex justify-between items-center hover:bg-gray-50">
              <div>
                <div className="font-medium">{integ.name}</div>
                <div className="text-sm text-gray-500">
                  Последняя синхронизация: {new Date(integ.last_sync).toLocaleString()}
                  {integ.error_message && (
                    <span className="text-red-600 ml-2">⚠️ {integ.error_message}</span>
                  )}
                </div>
              </div>
              <div className="flex items-center gap-3">
                <span className={`px-2 py-1 rounded-full text-xs ${
                  integ.status === 'ok' ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'
                }`}>
                  {integ.status === 'ok' ? '✅ Онлайн' : '❌ Ошибка'}
                </span>
                <button
                  onClick={() => handleSync(integ.id)}
                  className="text-blue-600 hover:text-blue-800 text-sm"
                >
                  ⟳ Обновить
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default AdminIntegrationsPage