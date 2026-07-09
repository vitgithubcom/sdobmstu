import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'

function AdminRolesPage() {
  const { user, logout } = useAuth()
  const currentTime = new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })

  const roles = [
    { name: 'Администратор', role: 'admin', permissions: ['Всё'] },
    { name: 'Менеджер', role: 'manager', permissions: ['Дашборд', 'Отчёты', 'Просмотр интеграций'] },
    { name: 'Аналитик', role: 'analyst', permissions: ['Дашборд', 'Отчёты', 'Детализация KPI'] },
    { name: 'Просмотрщик', role: 'viewer', permissions: ['Дашборд'] },
  ]

  return (
    <div className="max-w-7xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header currentTime={currentTime} user={user} onLogout={logout} title="Управление ролями" />

      <div className="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <table className="w-full">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Роль</th>
              <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Права доступа</th>
              <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">Действия</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-100">
            {roles.map((role) => (
              <tr key={role.role} className="hover:bg-gray-50">
                <td className="px-4 py-3">
                  <div className="font-medium">{role.name}</div>
                  <div className="text-xs text-gray-500">@{role.role}</div>
                </td>
                <td className="px-4 py-3">
                  <div className="flex flex-wrap gap-1">
                    {role.permissions.map((perm, idx) => (
                      <span key={idx} className="text-xs bg-gray-100 px-2 py-1 rounded-full">
                        {perm}
                      </span>
                    ))}
                  </div>
                </td>
                <td className="px-4 py-3 text-right">
                  <button className="text-xs text-blue-600 hover:text-blue-800">✏️ Редактировать</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default AdminRolesPage