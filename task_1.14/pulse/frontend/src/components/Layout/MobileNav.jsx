function MobileNav({ activeTab, setActiveTab }) {
  return (
    <div className="md:hidden fixed bottom-0 left-0 right-0 mobile-nav border-t border-gray-200 flex justify-around py-2 px-4">
      <button 
        onClick={() => setActiveTab('dashboard')} 
        className={`flex flex-col items-center p-2 rounded-xl ${activeTab === 'dashboard' ? 'text-blue-600 bg-blue-50' : 'text-gray-500'}`}
      >
        <span className="text-xl">📊</span>
        <span className="text-[11px]">Дашборд</span>
      </button>
      <button 
        onClick={() => setActiveTab('alerts')} 
        className={`flex flex-col items-center p-2 rounded-xl ${activeTab === 'alerts' ? 'text-amber-600 bg-amber-50' : 'text-gray-500'}`}
      >
        <span className="text-xl">⚠️</span>
        <span className="text-[11px]">Тревоги</span>
      </button>
      <button 
        onClick={() => setActiveTab('integrations')} 
        className={`flex flex-col items-center p-2 rounded-xl ${activeTab === 'integrations' ? 'text-green-600 bg-green-50' : 'text-gray-500'}`}
      >
        <span className="text-xl">🔄</span>
        <span className="text-[11px]">Источники</span>
      </button>
    </div>
  )
}

export default MobileNav