import { useState } from 'react'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'

function SettingsPage() {
  const { user, logout } = useAuth()
  const currentTime = new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
  const [settings, setSettings] = useState({
    company_name: 'ООО "Пульс"',
    timezone: 'UTC+3',
    email_notifications: true,
    backup_enabled: true,
    backup_time: '02:00',
  })

  const handleSave = (e) => {
    e.preventDefault()
    alert('Настройки сохранены (мок)')
  }

  return (
    <div className="max-w-4xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header currentTime={currentTime} user={user} onLogout={logout} title="Настройки системы" />

      <form onSubmit={handleSave} className="space-y-6">
        <div className="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <h2 className="font-semibold text-gray-800 mb-4">Общие настройки</h2>
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Название компании</label>
              <input
                type="text"
                value={settings.company_name}
                onChange={(e) => setSettings({ ...settings, company_name: e.target.value })}
                className="w-full px-4 py-2 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 outline-none"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Часовой пояс</label>
              <select
                value={settings.timezone}
                onChange={(e) => setSettings({ ...settings, timezone: e.target.value })}
                className="w-full px-4 py-2 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 outline-none"
              >
                <option value="UTC+0">UTC+0</option>
                <option value="UTC+3">UTC+3 (Москва)</option>
                <option value="UTC+5">UTC+5 (Екатеринбург)</option>
                <option value="UTC+7">UTC+7 (Красноярск)</option>
                <option value="UTC+10">UTC+10 (Владивосток)</option>
              </select>
            </div>
          </div>
        </div>

        <div className="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <h2 className="font-semibold text-gray-800 mb-4">Уведомления</h2>
          <div className="space-y-3">
            <label className="flex items-center gap-3 cursor-pointer">
              <input
                type="checkbox"
                checked={settings.email_notifications}
                onChange={(e) => setSettings({ ...settings, email_notifications: e.target.checked })}
                className="w-4 h-4 accent-blue-600"
              />
              <span className="text-sm">Email-уведомления о критических событиях</span>
            </label>
          </div>
        </div>

        <div className="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <h2 className="font-semibold text-gray-800 mb-4">Бэкап</h2>
          <div className="space-y-3">
            <label className="flex items-center gap-3 cursor-pointer">
              <input
                type="checkbox"
                checked={settings.backup_enabled}
                onChange={(e) => setSettings({ ...settings, backup_enabled: e.target.checked })}
                className="w-4 h-4 accent-blue-600"
              />
              <span className="text-sm">Автоматический бэкап</span>
            </label>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Время бэкапа</label>
              <input
                type="time"
                value={settings.backup_time}
                onChange={(e) => setSettings({ ...settings, backup_time: e.target.value })}
                className="px-4 py-2 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 outline-none"
              />
            </div>
            <button type="button" className="text-blue-600 hover:text-blue-800 text-sm">
              ⟳ Сделать бэкап сейчас
            </button>
          </div>
        </div>

        <button type="submit" className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-xl transition">
          Сохранить настройки
        </button>
      </form>
    </div>
  )
}

export default SettingsPage