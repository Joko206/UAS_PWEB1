// src/pages/Home.jsx
export default function Home() {
    return (
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center py-20">
          <h1 className="text-5xl font-bold text-white mb-6 animate-fade-in-up">
            Welcome to <span className="bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">AuthApp</span>
          </h1>
          <p className="text-xl text-gray-300 mb-8 max-w-2xl mx-auto">
            Secure authentication system with modern UI design and seamless user experience.
          </p>
          
          <div className="flex justify-center space-x-6">
            <a 
              href="/login" 
              className="bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-lg font-medium 
                transform transition-all duration-300 hover:scale-105 shadow-lg hover:shadow-blue-500/30"
            >
              Get Started
            </a>
          </div>
        </div>
      </div>
    )
  }