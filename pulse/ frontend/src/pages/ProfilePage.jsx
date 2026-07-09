import { useState } from 'react'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'
import MobileNav from '../components/Layout/MobileNav'
import api from '../api/client'

function ProfilePage() {
  const { user, logout } = useAuth()
  const [activeTab, setActiveTab] = useState('profile')
  const [currentTime, setCurrentTime] = useState(
    new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
  )
  const [formData, setFormData] = useState({
    full_name: user?.full_name || '',
    email: user?.email || '',
  })
  const [passwordData, setPasswordData] = useState({
    old_password: '',
    new_password: '',
    confirm_password: '',
  })
  const [message, setMessage] = useState({ type: '', text: '' })
  const [loading, setLoading] = useState(false)

  const handleUpdateProfile = async (e) => {
    e.preventDefault()
    setLoading(true)
    setMessage({ type: '', text: '' })

    try {
      await api.put('/users/profile', formData)
      setMessage({ type: 'success', text: 'Профиль успешно обновлён' })
    } catch (err) {
      setMessage({ type: 'error', text: err.response?.data?.error || 'Ошибка обновления' })
    } finally {
      setLoading(false)
    }
  }

  const handleChangePassword = async (e) => {
    e.preventDefault()
    if (passwordData.new_password !== passwordData.confirm_password) {
      setMessage({ type: 'error', text: 'Пароли не совпадают' })
      return
    }

    setLoading(true)
    setMessage({ type: '', text: '' })

    try {
      await api.put('/users/password', {
        old_password: passwordData.old_password,
        new_password: passwordData.new_password,
      })
      setMessage({ type: 'success', text: 'Пароль успешно изменён' })
      setPasswordData({ old_password: '', new_password: '', confirm_password: '' })
    } catch (err) {
      setMessage({ type: 'error', text: err.response?.data?.error || 'Ошибка смены пароля' })
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="max-w-4xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header 
        currentTime={currentTime} 
        user={user} 
        onLogout={logout} 
        title="Личный кабинет"
      />

      <div className="grid md:grid-cols-3 gap-6">
        {/* Боковое меню */}
        <div className="bg-white rounded-2xl p-4 shadow-sm border border-gray-100 h-fit">
          <div className="flex items-center gap-3 mb-4 pb-4 border-b">
            <div className="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center text-xl">
              {user?.full_name?.[0] || '👤'}
            </div>
            <div>
              <div className="font-medium">{user?.full_name}</div>
              <div className="text-xs text-gray-500">{user?.username}</div>
              <div className="text-[10px] bg-gray-100 px-2 py-0.5 rounded-full inline-block mt-1">
                {user?.role}
              </div>
            </div>
          </div>
          <nav className="space-y-1">
            <button className="w-full text-left px-3 py-2 rounded-xl bg-blue-50 text-blue-600 text-sm font-medium">
              📝 Основное
            </button>
            <button className="w-full text-left px-3 py-2 rounded-xl hover:bg-gray-50 text-gray-600 text-sm transition">
              🔒 Безопасность
            </button>
            <button className="w-full text-left px-3 py-2 rounded-xl hover:bg-gray-50 text-gray-600 text-sm transition">
              🔔 Уведомления
            </button>
          </nav>
        </div>

        {/* Формы */}
        <div className="md:col-span-2 space-y-6">
          {message.text && (
            <div className={`p-3 rounded-xl text-sm ${
              message.type === 'success' ? 'bg-green-50 text-green-700' : 'bg-red-50 text-red-700'
            }`}>
              {message.text}
            </div>
          )}

          {/* Профиль */}
          <div className="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
            <h2 className="font-semibold text-gray-800 mb-4">Основная информация</h2>
            <form onSubmit={handleUpdateProfile} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">ФИО</label>
                <input
                  type="text"
                  value={formData.full_name}
                  onChange={(e) => setFormData({ ...formData, full_name: e.target.value })}
                  className="w-full px-4 py-2 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
                <input
                  type="email"
                  value={formData.email}
                  onChange={(e) => setFormData({ ...formData, email: e.target.value })}
                  className="w-full px-4 py-2 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none"
                />
              </div>
              <button
                type="submit"
                disabled={loading}
                className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-xl transition disabled:opacity-50"
              >
                {loading ? 'Сохранение...' : 'Сохранить'}
              </button>
            </form>
          </div>

          {/* Смена пароля */}
          <div className="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
            <h2 className="font-semibold text-gray-800 mb-4">Смена пароля</h2>
            <form onSubmit={handleChangePassword} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Текущий пароль</label>
                <input
                  type="password"
                  value={passwordData.old_password}
                  onChange={(e) => setPasswordData({ ...passwordData, old_password: e.target.value })}
                  className="w-full px-4 py-2 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none"
                  required
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Новый пароль</label>
                <input
                  type="password"
                  value={passwordData.new_password}
                  onChange={(e) => setPasswordData({ ...passwordData, new_password: e.target.value })}
                  className="w-full px-4 py-2 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none"
                  required
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Подтверждение пароля</label>
                <input
                  type="password"
                  value={passwordData.confirm_password}
                  onChange={(e) => setPasswordData({ ...passwordData, confirm_password: e.target.value })}
                  className="w-full px-4 py-2 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none"
                  required
                />
              </div>
              <button
                type="submit"
                disabled={loading}
                className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-xl transition disabled:opacity-50"
              >
                {loading ? 'Смена...' : 'Сменить пароль'}
              </button>
            </form>
          </div>
        </div>
      </div>

      <MobileNav activeTab={activeTab} setActiveTab={setActiveTab} />
    </div>
  )
}

export default ProfilePage