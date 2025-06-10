import { TrendingUp, Car, Users, DollarSign } from 'lucide-react'

const stats = [
  {
    name: 'Viagens em Andamento',
    value: 12,
    icon: <TrendingUp className="text-blue-500 w-8 h-8" />,
    change: '+8%',
    changeType: 'positive',
  },
  {
    name: 'Veículos Disponíveis',
    value: 7,
    icon: <Car className="text-green-500 w-8 h-8" />,
    change: '-2%',
    changeType: 'negative',
  },
  {
    name: 'Motoristas Ativos',
    value: 15,
    icon: <Users className="text-purple-500 w-8 h-8" />,
    change: '+3%',
    changeType: 'positive',
  },
  {
    name: 'Faturamento Mensal',
    value: 'R$ 12.500',
    icon: <DollarSign className="text-yellow-500 w-8 h-8" />,
    change: '+12%',
    changeType: 'positive',
  },
]

export function Dashboard() {
  return (
    <div>
      <h2 className="text-2xl font-bold mb-6 text-gray-800">Dashboard</h2>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        {stats.map((stat) => (
          <div
            key={stat.name}
            className="card flex flex-col items-start gap-2 border border-gray-100 hover:shadow-lg transition-shadow duration-200"
          >
            <div className="flex items-center gap-3 mb-2">
              {stat.icon}
              <span className="text-gray-600 text-sm font-medium">{stat.name}</span>
            </div>
            <span className="text-3xl font-bold text-gray-900">{stat.value}</span>
            <span
              className={`text-xs font-semibold ${stat.changeType === 'positive' ? 'text-green-600' : 'text-red-600'}`}
            >
              {stat.change}
            </span>
          </div>
        ))}
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="card h-64 flex flex-col justify-center items-center border border-gray-100">
          <span className="text-gray-400">[Gráfico de Viagens por Mês]</span>
        </div>
        <div className="card h-64 flex flex-col justify-center items-center border border-gray-100">
          <span className="text-gray-400">[Gráfico de Disponibilidade de Veículos]</span>
        </div>
      </div>
    </div>
  )
} 