import { Link } from 'react-router-dom'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'

function AdminPage() {
  const { user, logout } = useAuth()
  const currentTime = new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })

  const cards = [
    { title: '👥 Пользователи', desc: 'Управление пользователями системы', path: '/admin/users', icon: '👥' },
    { title: '🔑 Роли и права', desc: 'Настройка ролей и доступа', path: '/admin/roles', icon: '🔑' },
    { title: '🔄 Интеграции', desc: 'Управление источниками данных', path: '/admin/integrations', icon: '🔄' },
    { title: '📋 Аудит', desc: 'Логи действий пользователей', path: '/admin/audit', icon: '📋' },
    { title: '⚙️ Настройки', desc: 'Общие настройки системы', path: '/settings', icon: '⚙️' },
  ]

  return (
    <div className="max-w-7xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header currentTime={currentTime} user={user} onLogout={logout} title="Админ-панель" />

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
        {cards.map((card, idx) => (
          <Link key={idx} to={card.path} className="block">
            <div className="bg-white rounded-2xl p-6 shadow-sm border border-gray-100 hover:shadow-md transition">
              <div className="text-3xl mb-2">{card.icon}</div>
              <h3 className="font-semibold text-gray-800">{card.title}</h3>
              <p className="text-sm text-gray-500 mt-1">{card.desc}</p>
            </div>
          </Link>
        ))}
      </div>
    </div>
  )
}

export default AdminPage