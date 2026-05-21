function Header({ currentTime, user, onLogout }) {
  return (
    <div className="flex justify-between items-center mb-6 flex-wrap gap-2">
      <div>
        <h1 className="text-2xl font-bold text-slate-800 tracking-tight">Пульс</h1>
        <p className="text-xs text-gray-400">KPI • реальное время (задержка ~90 сек)</p>
      </div>
      <div className="flex items-center gap-3">
        <div className="bg-white px-3 py-1.5 rounded-full text-sm shadow-sm border text-gray-600">
          🗓️ 19–25 мая
        </div>
        <div className="bg-gray-800 text-white px-3 py-1.5 rounded-full text-xs font-mono">
          {currentTime}
        </div>
        
        {user && (
          <div className="flex items-center gap-2 bg-white px-3 py-1.5 rounded-full shadow-sm border">
            <span className="text-sm">👤 {user.login}</span>
            <span className="text-[10px] bg-gray-100 px-1.5 py-0.5 rounded-full">
              {user.role === 'admin' ? 'админ' : 'пользователь'}
            </span>
            <button 
              onClick={onLogout}
              className="text-xs text-red-500 hover:text-red-700 ml-1"
            >
              🚪
            </button>
          </div>
        )}
      </div>
    </div>
  )
}

export default Header