import React, { useState } from 'react';
import type { ReactNode } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { Car, Users, TrendingUp, LogOut, Moon, Sun } from 'lucide-react';

interface DashboardLayoutProps {
	children: ReactNode;
}

const navItems = [
	{ name: 'Dashboard', to: '/', icon: <TrendingUp className="w-5 h-5" /> },
	{ name: 'Veículos', to: '/veiculos', icon: <Car className="w-5 h-5" /> },
	{ name: 'Motoristas', to: '/motoristas', icon: <Users className="w-5 h-5" /> },
	{ name: 'Viagens', to: '/viagens', icon: <TrendingUp className="w-5 h-5" /> },
];

const user = {
	name: 'Usuário',
	avatar: 'https://ui-avatars.com/api/?name=Usuario&background=0D8ABC&color=fff',
};

export function DashboardLayout({ children }: DashboardLayoutProps) {
	const location = useLocation();
	const [dark, setDark] = useState(false);

	React.useEffect(() => {
		if (dark) {
			document.documentElement.classList.add('dark');
		} else {
			document.documentElement.classList.remove('dark');
		}
	}, [dark]);

	return (
		<div className="flex min-h-screen bg-gray-50 dark:bg-gray-900">
			{/* Sidebar */}
			<aside className="w-64 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 flex flex-col items-center py-6 px-2 min-h-screen">
				<div className="mb-8 flex flex-col items-center">
					<img src={user.avatar} alt="Avatar" className="w-16 h-16 rounded-full border-2 border-blue-500 mb-2" />
					<span className="font-semibold text-gray-700 dark:text-gray-200">{user.name}</span>
				</div>
				<nav className="flex-1 w-full">
					{navItems.map((item) => (
						<Link
							key={item.to}
							to={item.to}
							className={`flex items-center gap-3 px-4 py-3 rounded-lg mb-2 text-base font-medium transition-colors w-full
								${location.pathname === item.to ? 'bg-blue-100 text-blue-700 dark:bg-blue-900 dark:text-blue-200' : 'text-gray-600 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700'}`}
						>
							{item.icon}
							<span>{item.name}</span>
						</Link>
					))}
				</nav>
				<div className="mt-auto flex flex-col gap-2 w-full">
					<button
						onClick={() => setDark((d) => !d)}
						className="flex items-center gap-2 px-4 py-2 rounded-lg bg-gray-200 dark:bg-gray-700 text-gray-700 dark:text-gray-200 hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors w-full"
					>
						{dark ? <Sun className="w-5 h-5" /> : <Moon className="w-5 h-5" />}
						<span>{dark ? 'Modo Claro' : 'Modo Escuro'}</span>
					</button>
					<Link
						to="/logout"
						className="flex items-center gap-2 px-4 py-2 rounded-lg bg-red-500 text-white hover:bg-red-600 transition-colors w-full"
					>
						<LogOut className="w-5 h-5" />
						<span>Logout</span>
					</Link>
				</div>
			</aside>
			{/* Conteúdo principal */}
			<div className="flex-1 flex flex-col min-h-screen">
				<header className="bg-white dark:bg-gray-800 shadow-sm border-b border-gray-200 dark:border-gray-700 px-8 py-4 flex items-center">
					<h1 className="text-2xl font-bold text-blue-600 dark:text-blue-300">Agencia Viagens</h1>
				</header>
				<main className="flex-1 p-8 bg-gray-50 dark:bg-gray-900">
					{children}
				</main>
				<footer className="bg-gray-100 dark:bg-gray-800 p-4 text-center text-sm text-gray-600 dark:text-gray-400 border-t border-gray-200 dark:border-gray-700">
					&copy; 2025 Agencia Viagens. Todos os direitos reservados.
				</footer>
			</div>
		</div>
	);
} 