import { useState } from 'react'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'
import MobileNav from '../components/Layout/MobileNav'

function ReportsPage() {
  const { user, logout } = useAuth()
  const [activeTab, setActiveTab] = useState('reports')
  const [currentTime, setCurrentTime] = useState(
    new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
  )

  const reports = [
    { id: 1, name: 'Отчёт по выручке', type: 'PDF', date: '2024-01-15', size: '2.3 MB' },
    { id: 2, name: 'Анализ OEE', type: 'Excel', date: '2024-01-14', size: '1.8 MB' },
    { id: 3, name: 'Сводка по цехам', type: 'PDF', date: '2024-01-13', size: '4.1 MB' },
  ]

  return (
    <div className="max-w-7xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header currentTime={currentTime} user={user} onLogout={logout} title="Отчёты" />

      <div className="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <div className="p-4 border-b flex justify-between items-center flex-wrap gap-2">
          <div className="text-sm text-gray-500">Сохранённые отчёты</div>
          <button className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-xl text-sm transition">
            ➕ Создать отчёт
          </button>
        </div>

        <div className="divide-y divide-gray-100">
          {reports.map((report) => (
            <div key={report.id} className="p-4 flex justify-between items-center hover:bg-gray-50">
              <div>
                <div className="font-medium">{report.name}</div>
                <div className="text-sm text-gray-500">
                  {report.type} • {report.date} • {report.size}
                </div>
              </div>
              <div className="flex gap-2">
                <button className="text-blue-600 hover:text-blue-800 text-sm">📥 Скачать</button>
                <button className="text-gray-500 hover:text-red-600 text-sm">🗑️</button>
              </div>
            </div>
          ))}
        </div>
      </div>

      <MobileNav activeTab={activeTab} setActiveTab={setActiveTab} />
    </div>
  )
}

export default ReportsPage