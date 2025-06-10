import { Calendar, Car, Users, TrendingUp } from 'lucide-react'

const stats = [
  {
    name: 'Viagens em Andamento',
    value: '12',
    icon: Calendar,
    change: '+4.75%',
    changeType: 'positive',
  },
  {
    name: 'Veículos Disponíveis',
    value: '8',
    icon: Car,
    change: '+2.02%',
    changeType: 'positive',
  },
  {
    name: 'Motoristas Ativos',
    value: '15',
    icon: Users,
    change: '-1.39%',
    changeType: 'negative',
  },
  {
    name: 'Faturamento Mensal',
    value: 'R$ 45.231,00',
    icon: TrendingUp,
    change: '+10.18%',
    changeType: 'positive',
  },
]

export function Dashboard() {
  return (
    <div>
      <h1 className="text-2xl font-semibold text-gray-900">Dashboard</h1>

      <div className="mt-8 grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
        {stats.map((item) => (
          <div
            key={item.name}
            className="relative overflow-hidden rounded-lg bg-white px-4 pb-12 pt-5 shadow sm:px-6 sm:pt-6"
          >
            <dt>
              <div className="absolute rounded-md bg-indigo-500 p-3">
                <item.icon className="h-6 w-6 text-white" aria-hidden="true" />
              </div>
              <p className="ml-16 truncate text-sm font-medium text-gray-500">{item.name}</p>
            </dt>
            <dd className="ml-16 flex items-baseline pb-6 sm:pb-7">
              <p className="text-2xl font-semibold text-gray-900">{item.value}</p>
              <p
                className={`ml-2 flex items-baseline text-sm font-semibold ${
                  item.changeType === 'positive' ? 'text-green-600' : 'text-red-600'
                }`}
              >
                {item.change}
              </p>
            </dd>
          </div>
        ))}
      </div>

      <div className="mt-8 grid grid-cols-1 gap-5 lg:grid-cols-2">
        {/* Aqui podemos adicionar gráficos de viagens por mês, disponibilidade de veículos, etc */}
        <div className="overflow-hidden rounded-lg bg-white shadow">
          <div className="p-6">
            <h3 className="text-lg font-medium leading-6 text-gray-900">
              Viagens por Mês
            </h3>
            <div className="mt-6 h-96">
              {/* TODO: Adicionar gráfico */}
              <div className="flex h-full items-center justify-center text-gray-500">
                Gráfico de viagens por mês
              </div>
            </div>
          </div>
        </div>

        <div className="overflow-hidden rounded-lg bg-white shadow">
          <div className="p-6">
            <h3 className="text-lg font-medium leading-6 text-gray-900">
              Disponibilidade de Veículos
            </h3>
            <div className="mt-6 h-96">
              {/* TODO: Adicionar gráfico */}
              <div className="flex h-full items-center justify-center text-gray-500">
                Gráfico de disponibilidade
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
} 