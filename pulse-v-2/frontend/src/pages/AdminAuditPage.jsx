import { useState, useEffect } from 'react'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'
import LoadingSpinner from '../components/Common/LoadingSpinner'
import api from '../api/client'

function AdminAuditPage() {
  const { user, logout } = useAuth()
  const [currentTime, setCurrentTime] = useState(
    new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
  )
  const [logs, setLogs] = useState([])
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
    fetchLogs()
  }, [])

  const fetchLogs = async () => {
    setLoading(true)
    try {
      const response = await api.get('/audit')
      // Преобразуем данные
      const logsData = Array.isArray(response.data) ? response.data.map(log => ({
        id: log.id,
        full_name: log.full_name || log.username || 'Система',
        username: log.username || '',
        action: log.action || 'Неизвестно',
        ip_address: log.ip_address || '-',
        created_at: log.created_at,
        details: log.details || ''
      })) : []
      setLogs(logsData)
    } catch (error) {
      console.error('Error fetching logs:', error)
      setLogs([])
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return <LoadingSpinner fullScreen />
  }

  return (
    <div className="max-w-7xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header currentTime={currentTime} user={user} onLogout={logout} title="Аудит действий" />

      <div className="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <div className="p-4 border-b flex justify-between items-center">
          <div className="text-sm text-gray-500">
            Всего записей: {logs.length}
          </div>
          <button
            onClick={fetchLogs}
            className="text-xs text-blue-600 hover:text-blue-800"
          >
            ⟳ Обновить
          </button>
        </div>

        {logs.length === 0 ? (
          <div className="p-8 text-center text-gray-400">
            <div className="text-4xl mb-2">📋</div>
            <p>Нет записей аудита</p>
            <p className="text-sm mt-1">Действия пользователей будут отображаться здесь</p>
          </div>
        ) : (
          <div className="overflow-x-auto">
            <table className="w-full">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Дата</th>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Пользователь</th>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Действие</th>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">IP</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-100">
                {logs.map((log) => (
                  <tr key={log.id} className="hover:bg-gray-50">
                    <td className="px-4 py-3 text-sm">
                      {log.created_at ? new Date(log.created_at).toLocaleString() : '—'}
                    </td>
                    <td className="px-4 py-3 text-sm">
                      <div>{log.full_name}</div>
                      <div className="text-xs text-gray-400">@{log.username}</div>
                    </td>
                    <td className="px-4 py-3 text-sm">
                      <span className="bg-gray-100 px-2 py-1 rounded-full text-xs">
                        {log.action}
                      </span>
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-500">{log.ip_address}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  )
}

export default AdminAuditPage