import { useState, useEffect } from 'react'
import LoginForm from './components/Auth/LoginForm'
import Header from './components/Layout/Header'
import MobileNav from './components/Layout/MobileNav'
import KPICard from './components/Dashboard/KPICard'
import RevenueChart from './components/Dashboard/RevenueChart'
import AlertsList from './components/Dashboard/AlertsList'
import IntegrationsTable from './components/Integrations/IntegrationsTable'
import { mockKPI } from './data/mockData'

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [activeTab, setActiveTab] = useState('dashboard')
  const [currentTime, setCurrentTime] = useState(new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' }))
  const [user, setUser] = useState(null)

  useEffect(() => {
    const token = localStorage.getItem('pulse_token')
    const savedUser = localStorage.getItem('pulse_user')
    if (token && savedUser) {
      setIsAuthenticated(true)
      setUser(JSON.parse(savedUser))
    }
  }, [])

  useEffect(() => {
    if (!isAuthenticated) return
    const timer = setInterval(() => {
      setCurrentTime(new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' }))
    }, 60000)
    return () => clearInterval(timer)
  }, [isAuthenticated])

  const handleLogin = () => {
    const savedUser = localStorage.getItem('pulse_user')
    setUser(JSON.parse(savedUser))
    setIsAuthenticated(true)
  }

  const handleLogout = () => {
    localStorage.removeItem('pulse_token')
    localStorage.removeItem('pulse_user')
    setIsAuthenticated(false)
    setUser(null)
  }

  if (!isAuthenticated) {
    return <LoginForm onLogin={handleLogin} />
  }

  return (
    <div className="max-w-7xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header currentTime={currentTime} user={user} onLogout={handleLogout} />
      
      <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-3 mb-6">
        {mockKPI.map(kpi => <KPICard key={kpi.id} kpi={kpi} />)}
      </div>

      <div className="grid md:grid-cols-3 gap-6">
        <div className="md:col-span-2"><RevenueChart /></div>
        <div><AlertsList /></div>
      </div>

      <IntegrationsTable />

      <MobileNav activeTab={activeTab} setActiveTab={setActiveTab} />
      
      <div className="hidden md:block text-center text-[11px] text-gray-400 mt-8 border-t pt-4">
        Данные обновляются каждые 30 сек • ⚡ «Пульс» — корпоративный мониторинг KPI
      </div>
    </div>
  )
}

export default App