import { useState, useEffect } from 'react'
import { useAuth } from '../hooks/useAuth'
import Header from '../components/Layout/Header'
import MobileNav from '../components/Layout/MobileNav'
import KPICard from '../components/Dashboard/KPICard'
import RevenueChart from '../components/Dashboard/RevenueChart'
import AlertsList from '../components/Dashboard/AlertsList'
import IntegrationsTable from '../components/Integrations/IntegrationsTable'
import LoadingSpinner from '../components/Common/LoadingSpinner'
import api from '../api/client'

function DashboardPage() {
  const { user, logout } = useAuth()
  const [activeTab, setActiveTab] = useState('dashboard')
  const [currentTime, setCurrentTime] = useState(
    new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
  )
  const [kpiData, setKpiData] = useState([])
  const [chartData, setChartData] = useState([])
  const [alerts, setAlerts] = useState([])
  const [integrations, setIntegrations] = useState([])
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
    fetchData()
  }, [])

  const fetchData = async () => {
    setLoading(true)
    try {
      const [kpiRes, chartRes, alertsRes, integrationsRes] = await Promise.all([
        api.get('/kpi'),
        api.get('/chart'),
        api.get('/alerts'),
        api.get('/integrations'),
      ])

      // Преобразуем KPI
      const transformedKPI = kpiRes.data.map(item => ({
        id: item.id,
        title: item.name,
        value: item.value.toLocaleString(),
        unit: item.unit,
        plan: item.plan,
        delta: parseFloat(item.delta.toFixed(1)),
        trend: item.direction === 'up' ? (item.delta >= 0 ? 'up' : 'down') : (item.delta <= 0 ? 'up' : 'down'),
        status: item.status,
      }))

      // Преобразуем Chart Data — проверяем разные варианты ключей
      const transformedChart = chartRes.data.map(item => ({
        name: item.name || item.period || '—',
        факт: item.fact || item.fact_value || item.value || 0,
        план: item.plan || item.plan_value || 0,
      }))

      console.log('Chart data:', transformedChart) // ← для отладки в консоли браузера

      setKpiData(transformedKPI)
      setChartData(transformedChart)
      setAlerts(alertsRes.data)
      setIntegrations(integrationsRes.data)
    } catch (error) {
      console.error('Error fetching data:', error)
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return <LoadingSpinner fullScreen />
  }

  return (
    <div className="max-w-7xl mx-auto px-3 pb-20 pt-4 md:pt-6">
      <Header 
        currentTime={currentTime} 
        user={user} 
        onLogout={logout} 
        title="Дашборд"
      />
      
      <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-3 mb-6">
        {kpiData.map(kpi => <KPICard key={kpi.id} kpi={kpi} />)}
      </div>

      <div className="grid md:grid-cols-3 gap-6">
        <div className="md:col-span-2">
          <RevenueChart data={chartData} />
        </div>
        <div>
          <AlertsList alerts={alerts} />
        </div>
      </div>

      <IntegrationsTable integrations={integrations} />

      <MobileNav activeTab={activeTab} setActiveTab={setActiveTab} />
      
      <div className="hidden md:block text-center text-[11px] text-gray-400 mt-8 border-t pt-4">
        Данные обновляются каждые 30 сек • ⚡ «Пульс» — корпоративный мониторинг KPI
      </div>
    </div>
  )
}

export default DashboardPage