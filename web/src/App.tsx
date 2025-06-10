import { BrowserRouter, Routes, Route, Navigate, Outlet } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { Layout } from './components/layout/Layout'
import { Dashboard } from './pages/dashboard/Dashboard'
import { Veiculos } from './pages/veiculos/Veiculos'
import { Motoristas } from './pages/motoristas/Motoristas'
import { Viagens } from './pages/viagens/Viagens'

const queryClient = new QueryClient()

function LayoutWrapper() {
  return (
    <Layout>
      <Outlet />
    </Layout>
  )
}

export function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route element={<LayoutWrapper />}>
            <Route index element={<Dashboard />} />
            <Route path="veiculos" element={<Veiculos />} />
            <Route path="motoristas" element={<Motoristas />} />
            <Route path="viagens" element={<Viagens />} />
            <Route path="*" element={<Navigate to="/" replace />} />
          </Route>
        </Routes>
      </BrowserRouter>
      <ReactQueryDevtools />
    </QueryClientProvider>
  )
}
