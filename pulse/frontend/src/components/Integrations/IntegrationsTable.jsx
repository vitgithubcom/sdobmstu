function IntegrationsTable({ integrations }) {
  return (
    <div className="bg-white rounded-2xl p-5 shadow-sm border border-gray-100 mt-4">
      <h3 className="font-semibold text-gray-800 mb-3">🔄 Статус источников данных</h3>
      <div className="space-y-2">
        {integrations.map((integ, idx) => (
          <div key={idx} className="flex justify-between items-center text-sm py-2 border-b border-gray-50">
            <span className="font-medium">{integ.name}</span>
            <div className="flex gap-3 items-center">
              <span className="text-gray-400 text-xs">синхр. {new Date(integ.last_sync).toLocaleTimeString()}</span>
              <span className={`px-2 py-0.5 rounded-full text-[10px] font-mono ${
                integ.status === 'ok' ? 'bg-green-100 text-green-700' : 'bg-amber-100 text-amber-700'
              }`}>
                {integ.status === 'ok' ? '✅ онлайн' : '⚠️ задержка'}
              </span>
              <span className="text-xs text-gray-500">lag {integ.lag_seconds}с</span>
            </div>
          </div>
        ))}
      </div>
      <button className="text-xs text-gray-500 mt-3 hover:text-blue-600">
        ⟳ Принудительный опрос всех источников
      </button>
    </div>
  )
}

export default IntegrationsTable