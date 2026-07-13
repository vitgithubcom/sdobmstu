import { Link } from 'react-router-dom'

function MobileNav({ activeTab, setActiveTab }) {
  return (
    <div className="md:hidden fixed bottom-0 left-0 right-0 mobile-nav border-t border-gray-200 flex justify-around py-2 px-4">
      <Link 
        to="/dashboard"
        onClick={() => setActiveTab('dashboard')}
        className={`flex flex-col items-center p-2 rounded-xl ${activeTab === 'dashboard' ? 'text-blue-600 bg-blue-50' : 'text-gray-500'}`}
      >
        <span className="text-xl">📊</span>
        <span className="text-[11px]">Дашборд</span>
      </Link>
      
      <Link 
        to="/reports"
        onClick={() => setActiveTab('reports')}
        className={`flex flex-col items-center p-2 rounded-xl ${activeTab === 'reports' ? 'text-blue-600 bg-blue-50' : 'text-gray-500'}`}
      >
        <span className="text-xl">📄</span>
        <span className="text-[11px]">Отчёты</span>
      </Link>
      
      <Link 
        to="/profile"
        onClick={() => setActiveTab('profile')}
        className={`flex flex-col items-center p-2 rounded-xl ${activeTab === 'profile' ? 'text-blue-600 bg-blue-50' : 'text-gray-500'}`}
      >
        <span className="text-xl">👤</span>
        <span className="text-[11px]">Профиль</span>
      </Link>
    </div>
  )
}

export default MobileNav